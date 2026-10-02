// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"errors"
	"strconv"
)

// signed is the set of signed integer types parsed from strings.
type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// unsigned is the set of unsigned integer types parsed from strings.
type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// StringToInt converts a base-10 string to int without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToInt(src string) (int, error) {
	return stringToSigned[int](src, strconv.IntSize, "int")
}

// StringToInt8 converts a base-10 string to int8 without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToInt8(src string) (int8, error) {
	return stringToSigned[int8](src, 8, "int8")
}

// StringToInt16 converts a base-10 string to int16 without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToInt16(src string) (int16, error) {
	return stringToSigned[int16](src, 16, "int16")
}

// StringToInt32 converts a base-10 string to int32 without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToInt32(src string) (int32, error) {
	return stringToSigned[int32](src, 32, "int32")
}

// StringToRune converts a base-10 string to rune without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToRune(src string) (rune, error) {
	return stringToSigned[rune](src, 32, "rune")
}

// StringToInt64 converts a base-10 string to int64 without precision loss.
//
// An integer literal (e.g. "-42") and a decimal floating-point literal (e.g.
// "4.2e1") are both parsed exactly over the whole int64 range; the latter
// must denote a whole number. Returns [ErrInvRange] for a value outside the
// destination range, [ErrFraction] for a number that is not whole, and
// [ErrInvFormat] for a string that is not a decimal number.
func StringToInt64(src string) (int64, error) {
	return stringToSigned[int64](src, 64, "int64")
}

// StringToUint converts a base-10 string to uint without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToUint(src string) (uint, error) {
	return stringToUnsigned[uint](src, strconv.IntSize, "uint")
}

// StringToUint8 converts a base-10 string to uint8 without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToUint8(src string) (uint8, error) {
	return stringToUnsigned[uint8](src, 8, "uint8")
}

// StringToByte converts a base-10 string to byte without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToByte(src string) (byte, error) {
	return stringToUnsigned[byte](src, 8, "byte")
}

// StringToUint16 converts a base-10 string to uint16 without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToUint16(src string) (uint16, error) {
	return stringToUnsigned[uint16](src, 16, "uint16")
}

// StringToUint32 converts a base-10 string to uint32 without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToUint32(src string) (uint32, error) {
	return stringToUnsigned[uint32](src, 32, "uint32")
}

// StringToUint64 converts a base-10 string to uint64 without precision loss.
// See [StringToInt64] for the accepted formats.
func StringToUint64(src string) (uint64, error) {
	return stringToUnsigned[uint64](src, 64, "uint64")
}

// stringToSigned parses src as a signed integer of the given bit size. When
// src is not an integer literal, it is rewritten as one by decimalToInteger.
func stringToSigned[T signed](src string, bits int, dst string) (T, error) {
	n, err := strconv.ParseInt(src, 10, bits)
	if errors.Is(err, strconv.ErrSyntax) {
		var neg bool
		var digits string
		if neg, digits, err = decimalToInteger(src); err != nil {
			return 0, NewError(err, "string", dst)
		}
		if neg {
			digits = "-" + digits
		}
		n, err = strconv.ParseInt(digits, 10, bits)
	}
	if err != nil {
		return 0, NewError(ErrInvRange, "string", dst).WithCause(err)
	}
	return T(n), nil
}

// stringToUnsigned parses src as an unsigned integer of the given bit size.
// When src is not an integer literal, it is rewritten as one by
// decimalToInteger.
func stringToUnsigned[T unsigned](src string, bits int, dst string) (T, error) {
	n, err := strconv.ParseUint(src, 10, bits)
	if errors.Is(err, strconv.ErrSyntax) {
		var neg bool
		var digits string
		if neg, digits, err = decimalToInteger(src); err != nil {
			return 0, NewError(err, "string", dst)
		}
		if neg {
			return 0, NewError(ErrInvRange, "string", dst)
		}
		n, err = strconv.ParseUint(digits, 10, bits)
	}
	if err != nil {
		return 0, NewError(ErrInvRange, "string", dst).WithCause(err)
	}
	return T(n), nil
}
