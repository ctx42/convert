// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"errors"
	"strconv"
	"strings"
)

// maxIntegerDigits is the number of digits in the longest integer any Go
// integer type holds (math.MaxUint64).
const maxIntegerDigits = 20

// decimalToInteger rewrites the decimal number literal s, optionally signed
// and with a fraction and an exponent (e.g. "-4.2e1"), as the digits of the
// whole number it denotes, without going through a floating-point value. It
// reports whether the number is negative; zero is never negative.
//
// Returns [ErrInvFormat] when s is not a decimal number literal, [ErrFraction]
// when the number is not whole, and [ErrInvRange] when the number has more
// digits than any integer type holds.
func decimalToInteger(s string) (bool, string, error) {
	var neg bool
	if s != "" && (s[0] == '+' || s[0] == '-') {
		neg = s[0] == '-'
		s = s[1:]
	}
	mnt, exs := s, "0"
	if i := strings.IndexAny(s, "eE"); i >= 0 {
		mnt, exs = s[:i], s[i+1:]
	}
	whole, frac, _ := strings.Cut(mnt, ".")
	digits := whole + frac
	if digits == "" || strings.TrimLeft(digits, "0123456789") != "" {
		return false, "", ErrInvFormat
	}
	exp, err := strconv.Atoi(exs)
	if err != nil && !errors.Is(err, strconv.ErrRange) {
		return false, "", ErrInvFormat
	}

	// The decimal point sits after point digits; leading zeros move it left,
	// trailing zeros do not move it.
	point := len(whole)
	trimmed := strings.TrimLeft(digits, "0")
	point -= len(digits) - len(trimmed)
	digits = strings.TrimRight(trimmed, "0")
	if digits == "" {
		return false, "0", nil
	}

	// Clamp the exponent so the arithmetic below cannot overflow. Past the
	// bound, a non-zero number has either more than maxIntegerDigits digits
	// or a non-zero digit after the point, so the result is unchanged.
	limit := 2*len(s) + maxIntegerDigits
	point += min(max(exp, -limit), limit)
	if point < len(digits) {
		return false, "", ErrFraction
	}
	if point > maxIntegerDigits {
		return false, "", ErrInvRange
	}
	return neg, digits + strings.Repeat("0", point-len(digits)), nil
}

// isDecimalNumber reports whether s consists only of characters used in a
// decimal number literal: digits, signs, a decimal point, and an exponent.
// It rejects the hexadecimal, underscore, infinity, and NaN forms that
// [strconv.ParseFloat] also accepts.
func isDecimalNumber(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
		case r == '+', r == '-', r == '.', r == 'e', r == 'E':
		default:
			return false
		}
	}
	return true
}
