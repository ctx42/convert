// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package xcast provides utilities for lossless type conversions.
package xcast

import (
	"errors"
	"reflect"
)

// SupportedTypes returns a slice of [reflect.Type] instances containing all
// types for which the package provides built-in converters. Any pair of types
// from this list can be converted in either direction.
func SupportedTypes() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf(0), // int
		reflect.TypeOf(int8(0)),
		reflect.TypeOf(int16(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(int64(0)),
		reflect.TypeOf(uint(0)),
		reflect.TypeOf(uint8(0)),
		reflect.TypeOf(uint16(0)),
		reflect.TypeOf(uint32(0)),
		reflect.TypeOf(uint64(0)),
		reflect.TypeOf(float32(0)),
		reflect.TypeOf(float64(0)),
		reflect.TypeOf(uintptr(0)),

		// TODO(rz):
		// reflect.TypeOf("string"),
		// reflect.TypeOf(time.Time{}),
		// reflect.TypeOf(time.Duration(0)),
	}
}

// Sentinel errors.
var (
	// ErrInvRange used when the input value isn't within a valid range of the
	// destination type.
	ErrInvRange = errors.New("value out of range")

	// ErrInvSafeRange used when input is within the valid range for conversion,
	// but conversion may lose precision.
	ErrInvSafeRange = errors.New("value out of safe range")

	// ErrUnkType used when conversion for a type is not defined.
	ErrUnkType = errors.New("unknown type")

	// ErrInvType used when a type is not valid in a given conversion context.
	ErrInvType = errors.New("invalid type")

	// ErrInvValue used when a value is not valid in a given conversion context.
	ErrInvValue = errors.New("invalid value")

	// ErrInvFormat used when a value's format is not valid in a given
	// conversion context.
	ErrInvFormat = errors.New("invalid format")

	// ErrUndConv used when conversion is undefined for given types.
	ErrUndConv = errors.New("conversion undefined")

	// ErrUnsupported represents explicitly not supported conversion.
	ErrUnsupported = errors.New("cast not supported")
)

// Safe integer boundaries for exact round-trip float32 to int32 conversions.
// The float32 can exactly represent all integers in the range [-2^24+1,2^24-1].
// Outside this range, some integers cannot be represented precisely, and
// conversion to / from integers will lose information or round incorrectly.
const (
	// Float32SafeIntMin represents the smallest int32 exactly representable by
	// the float32 type.
	Float32SafeIntMin = -(1 << 24) + 1

	// Float32SafeIntMax represents the biggest int32 exactly representable by
	// the float32 type.
	Float32SafeIntMax = 1<<24 - 1
)

// Safe integer boundaries for round-trip float64 to int64 conversions. The
// float64 can exactly represent all integers in the range [-2^53+1,2^53-1].
// Outside this range, some integers cannot be represented precisely, and
// conversion to / from integers will lose information or round incorrectly.
const (
	// Float64SafeIntMin represents the smallest int exactly representable by
	// the float64 type.
	Float64SafeIntMin = -(1 << 53) + 1

	// Float64SafeIntMax represents the biggest int exactly representable by
	// the float64 type.
	Float64SafeIntMax = 1<<53 - 1
)
