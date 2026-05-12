// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"bytes"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/goldy"
)

func Test_NewGenAnyToDst(t *testing.T) {
	t.Run("without options", func(t *testing.T) {
		// --- When ---
		have := NewGenAnyToDst("pkg")

		// --- Then ---
		assert.Equal(t, "pkg", have.pkg)
		assert.NotNil(t, have.code)
		assert.Cap(t, 20, have.imps)
		assert.Equal(t, Type{}, have.dst)
		assert.Equal(t, Options{}, have.ops)
	})

	t.Run("with options", func(t *testing.T) {
		// --- When ---
		have := NewGenAnyToDst("pkg", WithVerbose)

		// --- Then ---
		assert.Equal(t, "pkg", have.pkg)
		assert.NotNil(t, have.code)
		assert.Cap(t, 20, have.imps)
		assert.Equal(t, Type{}, have.dst)
		assert.Equal(t, Options{verbose: true}, have.ops)
	})
}

func Test_GenAnyToDst_GenerateCode_tabular(t *testing.T) {
	tt := []struct {
		testN string

		dst Type
		pth string
	}{
		{
			"any to float32",
			NumericType[float32](),
			"testdata/any_to_float32.gld",
		},
		{
			"any to byte",
			NumericType[uint8]().Alias("byte"),
			"testdata/any_to_byte.gld",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			gen := NewGenAnyToDst("pkg")
			dst := &bytes.Buffer{}

			// --- When ---
			err := gen.GenerateCode(dst, tc.dst)

			// --- Then ---
			assert.NoError(t, err)
			gld := goldy.Open(t, tc.pth)
			assert.Equal(t, gld.String(), dst.String())
		})
	}
}

func Test_GenAnyToDst_GenerateCode(t *testing.T) {
	t.Run("error - unsupported conversion", func(t *testing.T) {
		// --- Given ---
		dst := Type{value: NewValue("string")}
		gen := NewGenAnyToDst("pkg")
		buf := &bytes.Buffer{}

		// --- When ---
		err := gen.GenerateCode(buf, dst)

		// --- Then ---
		wMsg := "unsupported conversion between any and string"
		assert.ErrorEqual(t, wMsg, err)
		assert.Empty(t, buf.String())
	})
}

func Test_GenAnyToDst_GenerateTest_tabular(t *testing.T) {
	tt := []struct {
		testN string

		dst Type
		pth string
	}{
		{
			"int64 to float32",
			NumericType[float32](),
			"testdata/any_to_float32_test.gld",
		},
		{
			"int to uint",
			NumericType[uint](),
			"testdata/any_to_uint_test.gld",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			gen := NewGenAnyToDst("pkg")
			buf := &bytes.Buffer{}

			// --- When ---
			err := gen.GenerateTest(buf, tc.dst)

			// --- Then ---
			assert.NoError(t, err)
			gld := goldy.Open(t, tc.pth)
			assert.Equal(t, gld.String(), buf.String())
		})
	}
}
