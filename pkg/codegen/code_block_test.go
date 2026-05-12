// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"bytes"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_MustCodeBlock(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tpl := "a{{ .letter }}c"

		// --- When ---
		have := MustCodeBlock(tpl, "test")

		// --- Then ---
		assert.NotNil(t, have)
	})

	t.Run("panic - on error", func(t *testing.T) {
		// --- Given ---
		tpl := "a{{ .letter c"

		// --- When ---
		msg := assert.PanicMsg(t, func() { _ = MustCodeBlock("test", tpl) })

		// --- Then ---
		assert.Contain(t, "template: test:1: ", *msg)
	})
}

func Test_NewCodeBlock(t *testing.T) {
	t.Run("template without renderImports", func(t *testing.T) {
		// --- Given ---
		tpl := "a{{ .letter }}c"

		// --- When ---
		have, err := NewCodeBlock("test", tpl)

		// --- Then ---
		assert.NoError(t, err)
		assert.Nil(t, have.imps)
		assert.Equal(t, "test", have.tpl.Name())
	})

	t.Run("template with renderImports", func(t *testing.T) {
		// --- Given ---
		tpl := "import math\nimport time\na{{ .letter }}c"

		// --- When ---
		have, err := NewCodeBlock("test", tpl)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []string{"math", "time"}, have.imps)
		assert.Equal(t, "test", have.tpl.Name())
	})

	t.Run("error - invalid template", func(t *testing.T) {
		// --- Given ---
		tpl := "a{{ .letter c"

		// --- When ---
		have, err := NewCodeBlock("test", tpl)

		// --- Then ---
		assert.ErrorContain(t, "template: test:1: ", err)
		assert.Nil(t, have)
	})
}

func Test_CodeBlock_Imports(t *testing.T) {
	t.Run("without renderImports", func(t *testing.T) {
		// --- Given ---
		cb := &CodeBlock{}

		// --- When ---
		have := cb.Imports()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("with renderImports", func(t *testing.T) {
		// --- Given ---
		cb := &CodeBlock{imps: []string{"math", "time"}}

		// --- When ---
		have := cb.Imports()

		// --- Then ---
		assert.Equal(t, []string{"math", "time"}, have)
	})
}

func Test_CodeBlock_Write(t *testing.T) {
	t.Run("without renderImports", func(t *testing.T) {
		// --- Given ---
		dst := &bytes.Buffer{}
		tpl := "\r\n \ta{{ .letter }}c"
		data := map[string]any{"letter": "b"}
		cb := MustCodeBlock("test", tpl)

		// --- When ---
		err := cb.Render(dst, 2, data)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "\t\tabc\n", dst.String())
	})

	t.Run("with renderImports", func(t *testing.T) {
		// --- Given ---
		dst := &bytes.Buffer{}
		tpl := "import math\nimport time\n\n\n\r\n \ta{{ .letter }}c"
		data := map[string]any{"letter": "b"}
		cb := MustCodeBlock("test", tpl)

		// --- When ---
		err := cb.Render(dst, 0, data)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "abc\n", dst.String())
	})

	t.Run("error - missing template data", func(t *testing.T) {
		// --- Given ---
		dst := &bytes.Buffer{}
		tpl := "a{{ .letter }}c"
		cb := MustCodeBlock("test", tpl)

		// --- When ---
		err := cb.Render(dst, 0, nil)

		// --- Then ---
		wMsg := "template: test:1:4: executing \"test\" at <.letter>:" +
			" map has no entry for key \"letter\""
		assert.ErrorEqual(t, wMsg, err)
		assert.Empty(t, dst.String())
	})
}
