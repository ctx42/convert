// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"errors"
	"math"
	"strconv"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToFloat32_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src  string
		want float32
	}{
		{"zero", "0", 0},
		{"integer", "42", 42},
		{"shortest representation", "0.1", 0.1},
		{"negative", "-0.1", -0.1},
		{"exponent", "1e-3", 0.001},
		{"exact value", "0.100000001490116119384765625", 0.1},
		{"max", "3.4028235e+38", math.MaxFloat32},
		{"smallest denormal", "1e-45", math.SmallestNonzeroFloat32},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToFloat32(tc.src)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_StringToFloat32_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src string
		err error
		msg string
	}{
		{
			"error - not representable integer",
			"16777217",
			ErrInvSafeRange,
			"value out of safe range: from string to float32",
		},
		{
			"error - not representable fraction",
			"0.1000000001",
			ErrInvSafeRange,
			"value out of safe range: from string to float32",
		},
		{
			"error - beyond float32 range",
			"1e39",
			ErrInvRange,
			"value out of range: from string to float32",
		},
		{
			"error - beyond float64 range",
			"1e400",
			ErrInvRange,
			"value out of range: from string to float32",
		},
		{
			"error - letters",
			"abc",
			ErrInvFormat,
			"invalid format: from string to float32",
		},
		{
			"error - malformed number",
			"4.2.1",
			ErrInvFormat,
			"invalid format: from string to float32",
		},
		{
			"error - infinity",
			"Inf",
			ErrInvFormat,
			"invalid format: from string to float32",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToFloat32(tc.src)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float32(0), have)
		})
	}
}

func Test_StringToFloat64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		src := "-4.2e1"

		// --- When ---
		have, err := StringToFloat64(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, -42.0, have)
	})

	t.Run("error - beyond float64 range", func(t *testing.T) {
		// --- Given ---
		src := "1e400"

		// --- When ---
		have, err := StringToFloat64(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to float64", err)
		assert.Equal(t, 0.0, have)
	})

	t.Run("error - malformed number", func(t *testing.T) {
		// --- Given ---
		src := "4.2.1"

		// --- When ---
		have, err := StringToFloat64(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.ErrorEqual(t, "invalid format: from string to float64", err)
		assert.Equal(t, 0.0, have)
	})

	t.Run("error - letters", func(t *testing.T) {
		// --- Given ---
		src := "abc"

		// --- When ---
		have, err := StringToFloat64(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.ErrorEqual(t, "invalid format: from string to float64", err)
		assert.Equal(t, 0.0, have)
	})
}

func Test_parseFloatError(t *testing.T) {
	t.Run("range", func(t *testing.T) {
		// --- Given ---
		err := &strconv.NumError{Num: "1e400", Err: strconv.ErrRange}

		// --- When ---
		have := parseFloatError(err, "float64")

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, have)
		assert.ErrorIs(t, strconv.ErrRange, have)
	})

	t.Run("syntax", func(t *testing.T) {
		// --- Given ---
		err := errors.New("syntax")

		// --- When ---
		have := parseFloatError(err, "float64")

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, have)
		assert.ErrorEqual(t, "invalid format: from string to float64", have)
	})
}
