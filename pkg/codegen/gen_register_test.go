// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"bytes"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/goldy"
)

func Test_NewGenRegister(t *testing.T) {
	t.Run("without options", func(t *testing.T) {
		// --- When ---
		have := NewGenRegister("pkg")

		// --- Then ---
		assert.Equal(t, "pkg", have.pkg)
		assert.NotNil(t, have.code)
		assert.Cap(t, 20, have.imps)
		assert.Equal(t, Options{}, have.ops)
	})

	t.Run("with options", func(t *testing.T) {
		// --- Given ---
		opt := WithCopyright("copyright")

		// --- When ---
		have := NewGenRegister("pkg", opt)

		// --- Then ---
		assert.Equal(t, "pkg", have.pkg)
		assert.Equal(t, Options{copyright: "copyright"}, have.ops)
	})
}

func Test_GenRegister_GenerateCode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		gen := NewGenRegister("pkg")
		dst := &bytes.Buffer{}
		types := []Type{
			NumericType[int8](),
			NumericType[int64]().Alias("time", "Duration"),
		}

		// --- When ---
		err := gen.GenerateCode(dst, types)

		// --- Then ---
		assert.NoError(t, err)
		gld := goldy.Open(t, "testdata/register.gld")
		assert.Equal(t, gld.String(), dst.String())
	})

	t.Run("error - no types", func(t *testing.T) {
		// --- Given ---
		gen := NewGenRegister("pkg")
		dst := &bytes.Buffer{}

		// --- When ---
		err := gen.GenerateCode(dst, nil)

		// --- Then ---
		assert.ErrorEqual(t, "no types to register converters for", err)
		assert.Empty(t, dst.String())
	})
}
