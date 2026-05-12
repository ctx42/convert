// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"bytes"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/goldy"
)

func Test_NewGenSrcToDst(t *testing.T) {
	t.Run("without options", func(t *testing.T) {
		// --- When ---
		have := NewGenSrcToDst("pkg")

		// --- Then ---
		assert.Equal(t, "pkg", have.pkg)
		assert.NotNil(t, have.code)
		assert.Cap(t, 20, have.imps)
		assert.Equal(t, Type{}, have.src)
		assert.Equal(t, Type{}, have.dst)
		assert.Equal(t, Options{}, have.ops)
	})

	t.Run("with options", func(t *testing.T) {
		// --- When ---
		have := NewGenSrcToDst("pkg", WithVerbose)

		// --- Then ---
		assert.Equal(t, "pkg", have.pkg)
		assert.NotNil(t, have.code)
		assert.Cap(t, 20, have.imps)
		assert.Equal(t, Type{}, have.src)
		assert.Equal(t, Type{}, have.dst)
		assert.Equal(t, Options{verbose: true}, have.ops)
	})
}

func Test_GenSrcToDst_GenerateCode_tabular(t *testing.T) {
	tt := []struct {
		testN string

		from Type
		to   Type
		pth  string
	}{
		{
			"int64 to float32",
			NumericType[int64](),
			NumericType[float32](),
			"testdata/int64_to_float32.gld",
		},
		{
			"int64 to float64",
			NumericType[int64](),
			NumericType[float64](),
			"testdata/int64_to_float64.gld",
		},
		{
			"int16 to int32",
			NumericType[int16](),
			NumericType[int32](),
			"testdata/int16_to_int32.gld",
		},
		{
			"int32 to int32",
			NumericType[int32](),
			NumericType[int32](),
			"testdata/int32_to_int32.gld",
		},
		{
			"int64 to int32",
			NumericType[int64](),
			NumericType[int32](),
			"testdata/int64_to_int32.gld",
		},
		{
			"uint16 to int32",
			NumericType[uint16](),
			NumericType[int32](),
			"testdata/uint16_to_int32.gld",
		},
		{
			"uint32 to int32",
			NumericType[uint32](),
			NumericType[int32](),
			"testdata/uint32_to_int32.gld",
		},
		{
			"uint64 to int32",
			NumericType[uint64](),
			NumericType[int32](),
			"testdata/uint64_to_int32.gld",
		},
		{
			"uint16 to uint32",
			NumericType[uint16](),
			NumericType[uint32](),
			"testdata/uint16_to_uint32.gld",
		},
		{
			"uint32 to uint32",
			NumericType[uint32](),
			NumericType[uint32](),
			"testdata/uint32_to_uint32.gld",
		},
		{
			"uint64 to uint32",
			NumericType[uint64](),
			NumericType[uint32](),
			"testdata/uint64_to_uint32.gld",
		},
		{
			"int16 to uint32",
			NumericType[int16](),
			NumericType[uint32](),
			"testdata/int16_to_uint32.gld",
		},
		{
			"int32 to uint32",
			NumericType[int32](),
			NumericType[uint32](),
			"testdata/int32_to_uint32.gld",
		},
		{
			"int64 to uint32",
			NumericType[int64](),
			NumericType[uint32](),
			"testdata/int64_to_uint32.gld",
		},
		{
			"float32 to float64",
			NumericType[float32](),
			NumericType[float64](),
			"testdata/float32_to_float64.gld",
		},
		{
			"float64 to float64",
			NumericType[float64](),
			NumericType[float64](),
			"testdata/float64_to_float64.gld",
		},
		{
			"float32 to float32",
			NumericType[float32](),
			NumericType[float32](),
			"testdata/float32_to_float32.gld",
		},
		{
			"float64 to float32",
			NumericType[float64](),
			NumericType[float32](),
			"testdata/float64_to_float32.gld",
		},
		{
			"float64 to uint64",
			NumericType[float64](),
			NumericType[uint64](),
			"testdata/float64_to_uint64.gld",
		},
		{
			"float64 to uint32",
			NumericType[float64](),
			NumericType[uint32](),
			"testdata/float64_to_uint32.gld",
		},
		{
			"float64 to uint16",
			NumericType[float64](),
			NumericType[uint16](),
			"testdata/float64_to_uint16.gld",
		},
		{
			"float32 to uint64",
			NumericType[float32](),
			NumericType[uint64](),
			"testdata/float32_to_uint64.gld",
		},
		{
			"float32 to uint32",
			NumericType[float32](),
			NumericType[uint32](),
			"testdata/float32_to_uint32.gld",
		},
		{
			"float32 to uint16",
			NumericType[float32](),
			NumericType[uint16](),
			"testdata/float32_to_uint16.gld",
		},
		{
			"float64 to int64",
			NumericType[float64](),
			NumericType[int64](),
			"testdata/float64_to_int64.gld",
		},
		{
			"float64 to int32",
			NumericType[float64](),
			NumericType[int32](),
			"testdata/float64_to_int32.gld",
		},
		{
			"float32 to int64",
			NumericType[float32](),
			NumericType[int64](),
			"testdata/float32_to_int64.gld",
		},
		{
			"float32 to int32",
			NumericType[float32](),
			NumericType[int32](),
			"testdata/float32_to_int32.gld",
		},
		{
			"float32 to int16",
			NumericType[float32](),
			NumericType[int16](),
			"testdata/float32_to_int16.gld",
		},
		{
			"uint64 to float32",
			NumericType[uint64](),
			NumericType[float32](),
			"testdata/uint64_to_float32.gld",
		},
		{
			"uint64 to float64",
			NumericType[uint64](),
			NumericType[float64](),
			"testdata/uint64_to_float64.gld",
		},
		{
			"uint32 to float64",
			NumericType[uint32](),
			NumericType[float64](),
			"testdata/uint32_to_float64.gld",
		},
		{
			"uint32 to float32",
			NumericType[uint32](),
			NumericType[float32](),
			"testdata/uint32_to_float32.gld",
		},
		{
			"int64 to float64",
			NumericType[int64](),
			NumericType[float64](),
			"testdata/int64_to_float64.gld",
		},
		{
			"int64 to float32",
			NumericType[int64](),
			NumericType[float32](),
			"testdata/int64_to_float32.gld",
		},
		{
			"int32 to float64",
			NumericType[int32](),
			NumericType[float64](),
			"testdata/int32_to_float64.gld",
		},
		{
			"int32 to float32",
			NumericType[int32](),
			NumericType[float32](),
			"testdata/int32_to_float32.gld",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			gen := NewGenSrcToDst("pkg")
			dst := &bytes.Buffer{}

			// --- When ---
			err := gen.GenerateCode(dst, tc.from, tc.to)

			// --- Then ---
			assert.NoError(t, err)
			gld := goldy.Open(t, tc.pth)
			assert.Equal(t, gld.String(), dst.String())
		})
	}
}

func Test_ConvTestGen_GenerateTest_tabular(t *testing.T) {
	tt := []struct {
		testN string

		from Type
		to   Type
		pth  string
	}{
		{
			"int64 to float32",
			NumericType[int64](),
			NumericType[float32](),
			"testdata/int64_to_float32_test.gld",
		},
		{
			"int64 to float64",
			NumericType[int64](),
			NumericType[float64](),
			"testdata/int64_to_float64_test.gld",
		},
		{
			"int16 to int32",
			NumericType[int16](),
			NumericType[int32](),
			"testdata/int16_to_int32_test.gld",
		},
		{
			"int32 to int32",
			NumericType[int32](),
			NumericType[int32](),
			"testdata/int32_to_int32_test.gld",
		},
		{
			"int64 to int32",
			NumericType[int64](),
			NumericType[int32](),
			"testdata/int64_to_int32_test.gld",
		},
		{
			"uint16 to int32",
			NumericType[uint16](),
			NumericType[int32](),
			"testdata/uint16_to_int32_test.gld",
		},
		{
			"uint32 to int32",
			NumericType[uint32](),
			NumericType[int32](),
			"testdata/uint32_to_int32_test.gld",
		},
		{
			"uint64 to int32",
			NumericType[uint64](),
			NumericType[int32](),
			"testdata/uint64_to_int32_test.gld",
		},
		{
			"uint16 to uint32",
			NumericType[uint16](),
			NumericType[uint32](),
			"testdata/uint16_to_uint32_test.gld",
		},
		{
			"uint32 to uint32",
			NumericType[uint32](),
			NumericType[uint32](),
			"testdata/uint32_to_uint32_test.gld",
		},
		{
			"uint64 to uint32",
			NumericType[uint64](),
			NumericType[uint32](),
			"testdata/uint64_to_uint32_test.gld",
		},
		{
			"int16 to uint32",
			NumericType[int16](),
			NumericType[uint32](),
			"testdata/int16_to_uint32_test.gld",
		},
		{
			"int32 to uint32",
			NumericType[int32](),
			NumericType[uint32](),
			"testdata/int32_to_uint32_test.gld",
		},
		{
			"int64 to uint32",
			NumericType[int64](),
			NumericType[uint32](),
			"testdata/int64_to_uint32_test.gld",
		},
		{
			"float32 to float64",
			NumericType[float32](),
			NumericType[float64](),
			"testdata/float32_to_float64_test.gld",
		},
		{
			"float64 to float64",
			NumericType[float64](),
			NumericType[float64](),
			"testdata/float64_to_float64_test.gld",
		},
		{
			"float32 to float32",
			NumericType[float32](),
			NumericType[float32](),
			"testdata/float32_to_float32_test.gld",
		},
		{
			"float64 to float32",
			NumericType[float64](),
			NumericType[float32](),
			"testdata/float64_to_float32_test.gld",
		},
		{
			"float64 to uint64",
			NumericType[float64](),
			NumericType[uint64](),
			"testdata/float64_to_uint64_test.gld",
		},
		{
			"float64 to uint32",
			NumericType[float64](),
			NumericType[uint32](),
			"testdata/float64_to_uint32_test.gld",
		},
		{
			"float64 to uint16",
			NumericType[float64](),
			NumericType[uint16](),
			"testdata/float64_to_uint16_test.gld",
		},
		{
			"float32 to uint64",
			NumericType[float32](),
			NumericType[uint64](),
			"testdata/float32_to_uint64_test.gld",
		},
		{
			"float32 to uint32",
			NumericType[float32](),
			NumericType[uint32](),
			"testdata/float32_to_uint32_test.gld",
		},
		{
			"float32 to uint16",
			NumericType[float32](),
			NumericType[uint16](),
			"testdata/float32_to_uint16_test.gld",
		},
		{
			"float64 to int64",
			NumericType[float64](),
			NumericType[int64](),
			"testdata/float64_to_int64_test.gld",
		},
		{
			"float64 to int32",
			NumericType[float64](),
			NumericType[int32](),
			"testdata/float64_to_int32_test.gld",
		},
		{
			"float32 to int64",
			NumericType[float32](),
			NumericType[int64](),
			"testdata/float32_to_int64_test.gld",
		},
		{
			"float32 to int32",
			NumericType[float32](),
			NumericType[int32](),
			"testdata/float32_to_int32_test.gld",
		},
		{
			"float32 to int16",
			NumericType[float32](),
			NumericType[int16](),
			"testdata/float32_to_int16_test.gld",
		},
		{
			"uint64 to float32",
			NumericType[uint64](),
			NumericType[float32](),
			"testdata/uint64_to_float32_test.gld",
		},
		{
			"uint64 to float64",
			NumericType[uint64](),
			NumericType[float64](),
			"testdata/uint64_to_float64_test.gld",
		},
		{
			"uint32 to float64",
			NumericType[uint32](),
			NumericType[float64](),
			"testdata/uint32_to_float64_test.gld",
		},
		{
			"uint32 to float32",
			NumericType[uint32](),
			NumericType[float32](),
			"testdata/uint32_to_float32_test.gld",
		},
		{
			"int64 to float64",
			NumericType[int64](),
			NumericType[float64](),
			"testdata/int64_to_float64_test.gld",
		},
		{
			"int64 to float32",
			NumericType[int64](),
			NumericType[float32](),
			"testdata/int64_to_float32_test.gld",
		},
		{
			"int32 to float64",
			NumericType[int32](),
			NumericType[float64](),
			"testdata/int32_to_float64_test.gld",
		},
		{
			"int32 to float32",
			NumericType[int32](),
			NumericType[float32](),
			"testdata/int32_to_float32_test.gld",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			gen := NewGenSrcToDst("pkg")
			dst := &bytes.Buffer{}

			// --- When ---
			err := gen.GenerateTest(dst, tc.from, tc.to)

			// --- Then ---
			assert.NoError(t, err)
			gld := goldy.Open(t, tc.pth)
			assert.Equal(t, gld.String(), dst.String())
		})
	}
}

func Test_GenSrcToDst_Generate(t *testing.T) {
	t.Run("error - not numeric conversion", func(t *testing.T) {
		// --- Given ---
		from := Type{value: NewValue("string")}
		to := Type{value: NewValue("time", "Time")}
		gen := NewGenSrcToDst("pkg")
		dst := &bytes.Buffer{}

		// --- When ---
		err := gen.GenerateCode(dst, from, to)

		// --- Then ---
		wMsg := "no conversion actions found for string to time.Time"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, []string{"time"}, gen.imps)
		assert.Empty(t, dst.String())
	})
}

func Test_GenSrcToDst_numericConvFuncBody(t *testing.T) {
	t.Run("error - unsupported conversion", func(t *testing.T) {
		// --- Given ---
		gen := NewGenSrcToDst("pkg")
		gen.src = Type{value: NewValue("string")}
		gen.dst = NumericType[time.Duration]()

		// --- When ---
		err := gen.convFuncBody(nil)

		// --- Then ---
		wMsg := "no conversion actions found for string to time.Duration"
		assert.ErrorEqual(t, wMsg, err)
	})

	t.Run("error - unsupported action", func(t *testing.T) {
		// --- Given ---
		gen := NewGenSrcToDst("pkg")
		acs := []Action{NewAction("unsupported", nil)}

		// --- When ---
		err := gen.convFuncBody(acs)

		// --- Then ---
		assert.ErrorEqual(t, "unsupported action: unsupported", err)
	})
}
