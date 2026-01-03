// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_IsSigned_tabular(t *testing.T) {
	tt := []struct {
		testN string

		want bool
		have bool
	}{
		{"int", true, IsSigned[int]()},
		{"int8", true, IsSigned[int8]()},
		{"int16", true, IsSigned[int16]()},
		{"int32", true, IsSigned[int32]()},
		{"int64", true, IsSigned[int64]()},

		{"uint", false, IsSigned[uint]()},
		{"uint8", false, IsSigned[uint8]()},
		{"uint16", false, IsSigned[uint16]()},
		{"uint32", false, IsSigned[uint32]()},
		{"uint64", false, IsSigned[uint64]()},
		{"uintptr", false, IsSigned[uintptr]()},

		{"float32", true, IsSigned[float32]()},
		{"float64", true, IsSigned[float64]()},

		{"Int64", true, IsSigned[Int64]()},
		{"Uint64", false, IsSigned[Uint64]()},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.have)
		})
	}
}

func Test_IsFloat_tabular(t *testing.T) {
	tt := []struct {
		testN string

		want bool
		have bool
	}{
		{"int", false, IsFloat[int]()},
		{"float32", true, IsFloat[float32]()},
		{"float64", true, IsFloat[float64]()},
		{"Int64", false, IsFloat[Int64]()},
		{"Float64", true, IsFloat[Float64]()},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.have)
		})
	}
}

func Test_MinValue_tabular(t *testing.T) {
	tt := []struct {
		testN string

		typ  Type
		want string
	}{
		{"int64", NumericType[int64](), "math.MinInt64"},
		{"uint64", NumericType[uint64](), "0"},
		{"float32", NumericType[float32](), "Float32SafeIntMin"},
		{"float64", NumericType[float64](), "Float64SafeIntMin"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := MinValue(tc.typ)

			// --- Then ---
			assert.Equal(t, tc.want, have.Code())
		})
	}
}

func Test_MaxValue_tabular(t *testing.T) {
	tt := []struct {
		testN string

		typ  Type
		want string
	}{
		{"int64", NumericType[int64](), "math.MaxInt64"},
		{"uint64", NumericType[uint64](), "math.MaxUint64"},
		{"float32", NumericType[float32](), "Float32SafeIntMax"},
		{"float64", NumericType[float64](), "Float64SafeIntMax"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := MaxValue(tc.typ)

			// --- Then ---
			assert.Equal(t, tc.want, have.Code())
		})
	}
}

func Test_MinInteger_tabular(t *testing.T) {
	tt := []struct {
		testN string

		size   int
		signed bool
		want   string
	}{
		{"int8", 8, true, "math.MinInt8"},
		{"int16", 16, true, "math.MinInt16"},
		{"int32", 32, true, "math.MinInt32"},
		{"int64", 64, true, "math.MinInt64"},

		{"uint8", 8, false, "0"},
		{"uint16", 16, false, "0"},
		{"uint32", 32, false, "0"},
		{"uint64", 64, false, "0"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := MinInteger(tc.size, tc.signed)

			// --- Then ---
			assert.Equal(t, tc.want, have.Code())
		})
	}
}

func Test_MaxInteger_tabular(t *testing.T) {
	tt := []struct {
		testN string

		size   int
		signed bool
		want   string
	}{
		{"int8", 8, true, "math.MaxInt8"},
		{"int16", 16, true, "math.MaxInt16"},
		{"int32", 32, true, "math.MaxInt32"},
		{"int64", 64, true, "math.MaxInt64"},

		{"uint8", 8, false, "math.MaxUint8"},
		{"uint16", 16, false, "math.MaxUint16"},
		{"uint32", 32, false, "math.MaxUint32"},
		{"uint64", 64, false, "math.MaxUint64"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := MaxInteger(tc.size, tc.signed)

			// --- Then ---
			assert.Equal(t, tc.want, have.Code())
		})
	}
}

func Test_MinSafeFloat_tabular(t *testing.T) {
	tt := []struct {
		testN string

		size int
		want string
	}{
		{"float32", 32, "Float32SafeIntMin"},
		{"float64", 64, "Float64SafeIntMin"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := MinSafeFloat(tc.size)

			// --- Then ---
			assert.Equal(t, tc.want, have.Code())
		})
	}
}

func Test_MaxSafeFloat_tabular(t *testing.T) {
	tt := []struct {
		testN string

		size int
		want string
	}{
		{"float32", 32, "Float32SafeIntMax"},
		{"float64", 64, "Float64SafeIntMax"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := MaxSafeFloat(tc.size)

			// --- Then ---
			assert.Equal(t, tc.want, have.Code())
		})
	}
}
