// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_decimalToInteger_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src     string
		wNeg    bool
		wDigits string
	}{
		{"integer", "42", false, "42"},
		{"plus sign", "+42", false, "42"},
		{"minus sign", "-42", true, "42"},
		{"zero", "0", false, "0"},
		{"negative zero", "-0.0e5", false, "0"},
		{"leading zeros", "007", false, "7"},
		{"whole fraction", "42.000", false, "42"},
		{"trailing point", "42.", false, "42"},
		{"leading point", ".5e1", false, "5"},
		{"exponent", "4.2e1", false, "42"},
		{"upper case exponent", "42E0", false, "42"},
		{"signed exponent", "-4200e-2", true, "42"},
		{"exponent adds zeros", "1e19", false, "10000000000000000000"},
		{
			"fraction beyond float64",
			"9007199254740993.0",
			false,
			"9007199254740993",
		},
		{"leading zeros fraction", "0.0042e4", false, "42"},
		{
			"max digits",
			"1.8446744073709551615e19",
			false,
			"18446744073709551615",
		},
		{"huge exponent zero", "0e99999999999999999999", false, "0"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			hNeg, hDigits, err := decimalToInteger(tc.src)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.wNeg, hNeg)
			assert.Equal(t, tc.wDigits, hDigits)
		})
	}
}

func Test_decimalToInteger_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src string
		err error
	}{
		{"error - empty", "", ErrInvFormat},
		{"error - sign only", "-", ErrInvFormat},
		{"error - point only", ".", ErrInvFormat},
		{"error - letters", "abc", ErrInvFormat},
		{"error - double sign", "+-1", ErrInvFormat},
		{"error - two points", "4.2.1", ErrInvFormat},
		{"error - space", "4 2", ErrInvFormat},
		{"error - empty exponent", "4e", ErrInvFormat},
		{"error - exponent without mantissa", "e5", ErrInvFormat},
		{"error - fractional exponent", "4e1.5", ErrInvFormat},
		{"error - hexadecimal", "0x10", ErrInvFormat},
		{"error - underscore", "1_0", ErrInvFormat},
		{"error - infinity", "Inf", ErrInvFormat},
		{"error - fraction", "4.2", ErrFraction},
		{"error - fraction beyond float64", "1.0000000000000001", ErrFraction},
		{"error - tiny", "1e-400", ErrFraction},
		{
			"error - huge negative exponent",
			"1e-99999999999999999999",
			ErrFraction,
		},
		{"error - too many digits", "1e20", ErrInvRange},
		{"error - huge exponent", "1e99999999999999999999", ErrInvRange},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			hNeg, hDigits, err := decimalToInteger(tc.src)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.False(t, hNeg)
			assert.Equal(t, "", hDigits)
		})
	}
}

func Test_isDecimalNumber_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src  string
		want bool
	}{
		{"integer", "42", true},
		{"signed exponent", "-4.2e+1", true},
		{"empty", "", false},
		{"letters", "abc", false},
		{"space", "4 2", false},
		{"hexadecimal", "0x10", false},
		{"underscore", "1_0", false},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := isDecimalNumber(tc.src)

			// --- Then ---
			assert.Equal(t, tc.want, have)
		})
	}
}
