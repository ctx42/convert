// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"errors"
	"strconv"
)

// StringToFloat32 converts a decimal string to float32 without precision
// loss.
//
// The string is accepted when it is the shortest decimal representation of
// a float32 value, as written by [strconv.FormatFloat] with bit size 32
// (e.g. "0.1"), or when it denotes a float32 value exactly. Returns
// [ErrInvSafeRange] for a string that float32 cannot represent, such as
// "16777217", [ErrInvRange] for a value beyond the float32 range, and
// [ErrInvFormat] for a string that is not a decimal number.
func StringToFloat32(src string) (float32, error) {
	if !isDecimalNumber(src) {
		return 0, NewError(ErrInvFormat, "string", "float32")
	}
	f64, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return 0, parseFloatError(err, "float32")
	}
	f32, err := strconv.ParseFloat(src, 32)
	if err != nil {
		return 0, parseFloatError(err, "float32")
	}
	if f32 == f64 {
		return float32(f32), nil
	}
	short := strconv.FormatFloat(f32, 'g', -1, 32)
	if back, _ := strconv.ParseFloat(short, 64); back != f64 {
		return 0, NewError(ErrInvSafeRange, "string", "float32")
	}
	return float32(f32), nil
}

// StringToFloat64 converts a decimal string to float64. Returns
// [ErrInvRange] for a value beyond the float64 range and [ErrInvFormat] for
// a string that is not a decimal number.
func StringToFloat64(src string) (float64, error) {
	if !isDecimalNumber(src) {
		return 0, NewError(ErrInvFormat, "string", "float64")
	}
	f64, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return 0, parseFloatError(err, "float64")
	}
	return f64, nil
}

// parseFloatError returns the conversion error for an error returned by
// [strconv.ParseFloat] when parsing to the dst type.
func parseFloatError(err error, dst string) error {
	if errors.Is(err, strconv.ErrRange) {
		return NewError(ErrInvRange, "string", dst).WithCause(err)
	}
	return NewError(ErrInvFormat, "string", dst).WithCause(err)
}
