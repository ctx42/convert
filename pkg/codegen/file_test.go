// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"bytes"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/goldy"
)

func Test_file_writeCode(t *testing.T) {
	// --- Given ---
	fil := newFile("pkg")

	// --- When ---
	fil.writeCode("%s", "abc")

	// --- Then ---
	assert.Equal(t, "pkg", fil.pkg)
	assert.Equal(t, "abc", fil.code.String())
	assert.Len(t, 0, fil.imps)
}

func Test_file_addImport(t *testing.T) {
	// --- Given ---
	fil := newFile("pkg")

	// --- When ---
	fil.addImport("", "abc", "", "abc")

	// --- Then ---
	assert.Equal(t, []string{"", "abc", "", "abc"}, fil.imps)
}

func Test_file_renderImports(t *testing.T) {
	t.Run("without imports", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")

		// --- When ---
		have := fil.renderImports()

		// --- Then ---
		assert.Empty(t, fil.code.String())
		assert.Empty(t, have)
	})

	t.Run("with imports", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		fil.addImport("xyz", "abc")

		// --- When ---
		have := fil.renderImports()

		// --- Then ---
		assert.Empty(t, fil.code.String())

		want := "import (\n\t\"abc\"\n\t\"xyz\"\n)\n\n"
		assert.Equal(t, want, have)
	})

	t.Run("empty import strings are removed", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		fil.addImport("", "abc", "", "xyz")

		// --- When ---
		have := fil.renderImports()

		// --- Then ---
		assert.Empty(t, fil.code.String())

		want := "import (\n\t\"abc\"\n\t\"xyz\"\n)\n\n"
		assert.Equal(t, want, have)
	})

	t.Run("only empty import strings", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		fil.addImport("", "")

		// --- When ---
		have := fil.renderImports()

		// --- Then ---
		assert.Empty(t, have)
	})

	t.Run("duplicates are rendered once", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		fil.addImport("xyz", "abc", "xyz")

		// --- When ---
		have := fil.renderImports()

		// --- Then ---
		want := "import (\n\t\"abc\"\n\t\"xyz\"\n)\n\n"
		assert.Equal(t, want, have)
		assert.Equal(t, []string{"xyz", "abc", "xyz"}, fil.imps)
	})

	t.Run("imports are sorted alphabetically ", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		fil.addImport("xyz", "abc")

		// --- When ---
		have := fil.renderImports()

		// --- Then ---
		assert.Empty(t, fil.code.String())

		want := "import (\n\t\"abc\"\n\t\"xyz\"\n)\n\n"
		assert.Equal(t, want, have)
	})
}

func Test_file_tabularTest(t *testing.T) {
	t.Run("without header", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		table := MustCodeBlock("table", "tt := []int{\n")
		loop := MustCodeBlock("loop", "import abc\nuse({{.v}}, tt)\n")
		data := map[string]any{"v": 1}

		// --- When ---
		err := fil.tabularTest("Test_A", nil, table, loop, "\t\t1,\n", data)

		// --- Then ---
		assert.NoError(t, err)
		want := "" +
			"func Test_A(t *testing.T) {\n" +
			"\ttt := []int{\n" +
			"\t\t1,\n" +
			"\t}\n" +
			"\n" +
			"\tuse(1, tt)\n" +
			"}\n"
		assert.Equal(t, want, fil.code.String())
		assert.Equal(t, []string{"abc"}, fil.imps)
	})

	t.Run("with header after other code", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		fil.writeCode("code\n")
		header := MustCodeBlock("header", "import xyz\nskip()\n")
		table := MustCodeBlock("table", "tt := []int{\n")
		loop := MustCodeBlock("loop", "use(tt)\n")

		// --- When ---
		err := fil.tabularTest("Test_A", header, table, loop, "", nil)

		// --- Then ---
		assert.NoError(t, err)
		want := "" +
			"code\n" +
			"\n" +
			"func Test_A(t *testing.T) {\n" +
			"\tskip()\n" +
			"\n" +
			"\ttt := []int{\n" +
			"\t}\n" +
			"\n" +
			"\tuse(tt)\n" +
			"}\n"
		assert.Equal(t, want, fil.code.String())
		assert.Equal(t, []string{"xyz"}, fil.imps)
	})

	t.Run("error - header", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		header := MustCodeBlock("header", "{{.missing}}\n")
		table := MustCodeBlock("table", "tt\n")
		loop := MustCodeBlock("loop", "loop\n")
		data := map[string]any{}

		// --- When ---
		err := fil.tabularTest("Test_A", header, table, loop, "", data)

		// --- Then ---
		assert.ErrorContain(t, `map has no entry for key "missing"`, err)
	})

	t.Run("error - table", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		table := MustCodeBlock("table", "{{.missing}}\n")
		loop := MustCodeBlock("loop", "loop\n")
		data := map[string]any{}

		// --- When ---
		err := fil.tabularTest("Test_A", nil, table, loop, "", data)

		// --- Then ---
		assert.ErrorContain(t, `map has no entry for key "missing"`, err)
	})

	t.Run("error - loop", func(t *testing.T) {
		// --- Given ---
		fil := newFile("pkg")
		table := MustCodeBlock("table", "tt\n")
		loop := MustCodeBlock("loop", "{{.missing}}\n")
		data := map[string]any{}

		// --- When ---
		err := fil.tabularTest("Test_A", nil, table, loop, "", data)

		// --- Then ---
		assert.ErrorContain(t, `map has no entry for key "missing"`, err)
	})
}

func Test_file_reset(t *testing.T) {
	// --- Given ---
	fil := newFile("pkg")
	fil.writeCode("abc")
	fil.addImport("xyz")

	// --- When ---
	fil.reset()

	// --- Then ---
	assert.Equal(t, "pkg", fil.pkg)
	assert.Empty(t, fil.imps)
	assert.Empty(t, fil.code.String())
}

func Test_file_WriteTo(t *testing.T) {
	t.Run("all file parts written", func(t *testing.T) {
		// --- Given ---
		fil := newFile(
			"pkg",
			WithCopyright("// copyright.\n"),
			WithGeneratedBy("// generated by.\n"),
		)
		fil.addImport("time", "", "math", "")
		fil.writeCode("%s", "func my() {}")
		buf := &bytes.Buffer{}

		// --- When ---
		have, err := fil.WriteTo(buf)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int64(84), have)

		gld := goldy.Open(t, "testdata/file_write_to.gld")
		assert.Equal(t, gld.String(), buf.String())
	})
}
