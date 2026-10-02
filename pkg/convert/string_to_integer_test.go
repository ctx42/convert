// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"strconv"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToInt(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := strconv.Itoa(math.MaxInt)

		// --- When ---
		have, err := StringToInt(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, math.MaxInt, have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "9223372036854775808"

		// --- When ---
		have, err := StringToInt(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to int", err)
		assert.Equal(t, int(0), have)
	})
}

func Test_StringToInt8(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "127"

		// --- When ---
		have, err := StringToInt8(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int8(math.MaxInt8), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "128"

		// --- When ---
		have, err := StringToInt8(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to int8", err)
		assert.Equal(t, int8(0), have)
	})
}

func Test_StringToInt16(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "32767"

		// --- When ---
		have, err := StringToInt16(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int16(math.MaxInt16), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "32768"

		// --- When ---
		have, err := StringToInt16(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to int16", err)
		assert.Equal(t, int16(0), have)
	})
}

func Test_StringToInt32(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "2147483647"

		// --- When ---
		have, err := StringToInt32(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int32(math.MaxInt32), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "2147483648"

		// --- When ---
		have, err := StringToInt32(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to int32", err)
		assert.Equal(t, int32(0), have)
	})
}

func Test_StringToRune(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "2147483647"

		// --- When ---
		have, err := StringToRune(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, rune(math.MaxInt32), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "2147483648"

		// --- When ---
		have, err := StringToRune(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to rune", err)
		assert.Equal(t, rune(0), have)
	})
}

func Test_StringToInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src  string
		want int64
	}{
		{"zero", "0", 0},
		{"positive", "42", 42},
		{"negative", "-42", -42},
		{"plus sign", "+42", 42},
		{"min", "-9223372036854775808", math.MinInt64},
		{"max", "9223372036854775807", math.MaxInt64},
		{"beyond float64 precision", "9007199254740993", 9007199254740993},
		{"whole number with fraction", "42.0", 42},
		{"exponent", "4.2e1", 42},
		{"upper case exponent", "42E0", 42},
		{"exponent beyond float64 precision", "1e17", 100000000000000000},
		{"exponent max", "9.223372036854775807e18", math.MaxInt64},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToInt64(tc.src)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_StringToInt64_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src string
		err error
		msg string
	}{
		{
			"error - above max",
			"9223372036854775808",
			ErrInvRange,
			"value out of range: from string to int64",
		},
		{
			"error - below min",
			"-9223372036854775809",
			ErrInvRange,
			"value out of range: from string to int64",
		},
		{
			"error - fraction",
			"4.2",
			ErrFraction,
			"must be a whole number: from string to int64",
		},
		{
			"error - fraction beyond float64 precision",
			"1.0000000000000001",
			ErrFraction,
			"must be a whole number: from string to int64",
		},
		{
			"error - fraction near float64 precision limit",
			"9007199254740991.4",
			ErrFraction,
			"must be a whole number: from string to int64",
		},
		{
			"error - fraction below float64 range",
			"1e-400",
			ErrFraction,
			"must be a whole number: from string to int64",
		},
		{
			"error - exponent above max",
			"9.223372036854775808e18",
			ErrInvRange,
			"value out of range: from string to int64",
		},
		{
			"error - exponent below min",
			"-9.223372036854775809e18",
			ErrInvRange,
			"value out of range: from string to int64",
		},
		{
			"error - float beyond float64 range",
			"1e400",
			ErrInvRange,
			"value out of range: from string to int64",
		},
		{
			"error - empty",
			"",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
		{
			"error - letters",
			"abc",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
		{
			"error - malformed number",
			"4.2.1",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
		{
			"error - surrounding space",
			" 42",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
		{
			"error - hexadecimal",
			"0x1p4",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
		{
			"error - underscore",
			"1_000",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
		{
			"error - infinity",
			"Inf",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
		{
			"error - NaN",
			"NaN",
			ErrInvFormat,
			"invalid format: from string to int64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToInt64(tc.src)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})
	}
}

func Test_StringToUint(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := strconv.FormatUint(math.MaxUint, 10)

		// --- When ---
		have, err := StringToUint(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint(math.MaxUint), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "18446744073709551616"

		// --- When ---
		have, err := StringToUint(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to uint", err)
		assert.Equal(t, uint(0), have)
	})
}

func Test_StringToUint8(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "255"

		// --- When ---
		have, err := StringToUint8(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint8(math.MaxUint8), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "256"

		// --- When ---
		have, err := StringToUint8(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to uint8", err)
		assert.Equal(t, uint8(0), have)
	})
}

func Test_StringToByte(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "255"

		// --- When ---
		have, err := StringToByte(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, byte(math.MaxUint8), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "256"

		// --- When ---
		have, err := StringToByte(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to byte", err)
		assert.Equal(t, byte(0), have)
	})
}

func Test_StringToUint16(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "65535"

		// --- When ---
		have, err := StringToUint16(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint16(math.MaxUint16), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "65536"

		// --- When ---
		have, err := StringToUint16(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to uint16", err)
		assert.Equal(t, uint16(0), have)
	})
}

func Test_StringToUint32(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		// --- Given ---
		src := "4294967295"

		// --- When ---
		have, err := StringToUint32(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint32(math.MaxUint32), have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "4294967296"

		// --- When ---
		have, err := StringToUint32(src)

		// --- Then ---
		assert.ErrorIs(t, ErrInvRange, err)
		assert.ErrorEqual(t, "value out of range: from string to uint32", err)
		assert.Equal(t, uint32(0), have)
	})
}

func Test_StringToUint64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src  string
		want uint64
	}{
		{"zero", "0", 0},
		{"positive", "42", 42},
		{"plus sign", "+42", 42},
		{"max", "18446744073709551615", math.MaxUint64},
		{"beyond float64 precision", "9007199254740993", 9007199254740993},
		{"exponent", "4.2e1", 42},
		{"plus sign max", "+18446744073709551615", math.MaxUint64},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToUint64(tc.src)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_StringToUint64_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src string
		err error
		msg string
	}{
		{
			"error - above max",
			"18446744073709551616",
			ErrInvRange,
			"value out of range: from string to uint64",
		},
		{
			"error - plus sign above max",
			"+18446744073709551616",
			ErrInvRange,
			"value out of range: from string to uint64",
		},
		{
			"error - negative",
			"-1",
			ErrInvRange,
			"value out of range: from string to uint64",
		},
		{
			"error - fraction",
			"4.2",
			ErrFraction,
			"must be a whole number: from string to uint64",
		},
		{
			"error - letters",
			"abc",
			ErrInvFormat,
			"invalid format: from string to uint64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToUint64(tc.src)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint64(0), have)
		})
	}
}

func Test_StringToInt64_registered(t *testing.T) {
	// --- Given ---
	src := "9007199254740993"

	// --- When ---
	have, err := AnyToInt64(src)

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, int64(9007199254740993), have)
}
