// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package convert provides utilities for lossless type conversions.
package convert

import (
	"errors"
	"reflect"
	"time"
)

// registry represents package level [Registry].
var registry = &Registry{}

// Register adds the provided converter to the package-level [Registry].
// If a converter for the same source-destination type pair already exists,
// it is replaced, and the previous converter is returned; otherwise nil is
// returned.
func Register[Src, Dst any](cnv SrcToDst[Src, Dst]) SrcToDst[Src, Dst] {
	return RegisterConverter(registry, cnv)
}

// Lookup returns the converter for the given source-destination type pair from
// the package-level [Registry]. Returns nil if no converter was registered for
// the given source-destination type pair.
func Lookup[Src, Dst any]() SrcToDst[Src, Dst] {
	return LookupConverter[Src, Dst](registry)
}

// SrcToDst represents a converter function that attempts lossless conversion
// of a value from the type "From" to the "To" type. On success, it returns the
// converted value and a nil error. On failure (e.g., truncation, underflow,
// overflow, or semantic loss), it returns the zero value of "To" along with a
// non-nil error describing the issue.
type SrcToDst[Src, Dst any] func(Src) (Dst, error)

// AnyToAny is a non-generic version of [SrcToDst]. The behavior is exactly the
// same in terms of conversion and error handling.
type AnyToAny func(any) (any, error)

// Sentinel errors.
var (
	// ErrInvRange used when a value isn't within a valid range.
	ErrInvRange = errors.New("value out of range")

	// ErrInvSafeRange used when a value is within the valid range, but it may
	// lead to precision loss.
	ErrInvSafeRange = errors.New("value out of safe range")

	// ErrUnsType used when conversion for a type is not defined.
	ErrUnsType = errors.New("unsupported type")

	// ErrInvType used when a type is not valid in a given conversion context.
	ErrInvType = errors.New("invalid type")

	// ErrInvValue used when a value is not valid in a given conversion context.
	ErrInvValue = errors.New("invalid value")

	// ErrFraction used when a value must not be a floating-point number with a
	// fraction in a given conversion context.
	ErrFraction = errors.New("must be a whole number")

	// ErrInvFormat used when a value's format is not valid in a given
	// conversion context.
	ErrInvFormat = errors.New("invalid format")

	// ErrUnkConv used when conversion is undefined for given types.
	ErrUnkConv = errors.New("conversion undefined")

	// ErrUns represents explicitly not supported conversion.
	ErrUns = errors.New("unsupported conversion")
)

// ToAnyAny returns [AnyToAny] based on [SrcToDst].
func ToAnyAny[Src, Dst any](conv SrcToDst[Src, Dst]) AnyToAny {
	return func(value any) (any, error) {
		var ok bool
		var src Src
		if src, ok = value.(Src); !ok {
			var dst Dst
			return dst, NewError(ErrInvType, src, value).
				Format("%v: expected %T got %T")
		}
		return conv(src)
	}
}

// Option is a signature for a conversion option function.
type Option func(*Options)

// Options represent conversion function options.
type Options struct {
	reg *Registry
}

// NewOptions constructs an instance of [Options].
//
// By default, the package-level registry is used.
func NewOptions(ops ...Option) Options {
	opts := Options{reg: registry}
	for _, op := range ops {
		op(&opts)
	}
	return opts
}

// WithRegistry returns an [Option] for setting the registry for conversions.
func WithRegistry(reg *Registry) Option {
	return func(ops *Options) { ops.reg = reg }
}

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
	}
}

// Safe integer boundaries for the exact round-trip of float32 to integer
// conversions. The float32 can exactly represent all integers in the range
// [-2^24+1,2^24-1]. Outside this range, some integers cannot be represented
// precisely, and conversion to / from integers will lose information or round
// incorrectly.
const (
	// Float32SafeBits number of bits in a safe range [-2^24+1,2^24-1].
	Float32SafeBits = 24

	// Float32SafeIntMin represents the smallest integer exactly representable
	// by the float32 type.
	Float32SafeIntMin = -(1 << Float32SafeBits) + 1

	// Float32SafeIntMax represents the biggest integer exactly representable
	// by the float32 type.
	Float32SafeIntMax = 1<<Float32SafeBits - 1
)

// Safe integer boundaries for the exact round-trip of float64 to integer
// conversions. The float64 can exactly represent all integers in the range
// [-2^53+1,2^53-1]. Outside this range, some integers cannot be represented
// precisely, and conversion to / from integers will lose information or round
// incorrectly.
const (
	// Float64SafeBits number of bits in a safe range [-2^53+1,2^53-1].
	Float64SafeBits = 53

	// Float64SafeIntMin represents the smallest int exactly representable by
	// the float64 type.
	Float64SafeIntMin = -(1 << Float64SafeBits) + 1

	// Float64SafeIntMax represents the biggest int exactly representable by
	// the float64 type.
	Float64SafeIntMax = 1<<Float64SafeBits - 1
)

func init() {
	// Converters implemented by hand.
	Register(BoolToBool)
	Register(StringToDuration)
	Register(StringToString)
	Register(StringToTime(time.RFC3339Nano))

	// Generated numeric converters.
	Register(DurationToByte)
	Register(DurationToDuration)
	Register(DurationToFloat32)
	Register(DurationToFloat64)
	Register(DurationToInt)
	Register(DurationToInt16)
	Register(DurationToInt32)
	Register(DurationToInt64)
	Register(DurationToInt8)
	Register(DurationToUint)
	Register(DurationToUint16)
	Register(DurationToUint32)
	Register(DurationToUint64)
	Register(DurationToUint8)
	Register(DurationToUintptr)
	Register(Float32ToDuration)
	Register(Float32ToFloat32)
	Register(Float32ToFloat64)
	Register(Float32ToInt)
	Register(Float32ToInt16)
	Register(Float32ToInt32)
	Register(Float32ToInt32)
	Register(Float32ToInt64)
	Register(Float32ToInt8)
	Register(Float32ToUint)
	Register(Float32ToUint16)
	Register(Float32ToUint32)
	Register(Float32ToUint64)
	Register(Float32ToUint8)
	Register(Float32ToUint8)
	Register(Float32ToUintptr)
	Register(Float64ToDuration)
	Register(Float64ToFloat32)
	Register(Float64ToFloat64)
	Register(Float64ToInt)
	Register(Float64ToInt16)
	Register(Float64ToInt32)
	Register(Float64ToInt32)
	Register(Float64ToInt64)
	Register(Float64ToInt8)
	Register(Float64ToUint)
	Register(Float64ToUint16)
	Register(Float64ToUint32)
	Register(Float64ToUint64)
	Register(Float64ToUint8)
	Register(Float64ToUint8)
	Register(Float64ToUintptr)
	Register(Int16ToDuration)
	Register(Int16ToFloat32)
	Register(Int16ToFloat64)
	Register(Int16ToInt)
	Register(Int16ToInt16)
	Register(Int16ToInt32)
	Register(Int16ToInt32)
	Register(Int16ToInt64)
	Register(Int16ToInt8)
	Register(Int16ToUint)
	Register(Int16ToUint16)
	Register(Int16ToUint32)
	Register(Int16ToUint64)
	Register(Int16ToUint8)
	Register(Int16ToUint8)
	Register(Int16ToUintptr)
	Register(Int32ToDuration)
	Register(Int32ToFloat32)
	Register(Int32ToFloat32)
	Register(Int32ToFloat64)
	Register(Int32ToFloat64)
	Register(Int32ToInt)
	Register(Int32ToInt)
	Register(Int32ToInt16)
	Register(Int32ToInt16)
	Register(Int32ToInt32)
	Register(Int32ToInt32)
	Register(Int32ToInt32)
	Register(Int32ToInt32)
	Register(Int32ToInt64)
	Register(Int32ToInt64)
	Register(Int32ToInt8)
	Register(Int32ToInt8)
	Register(Int32ToUint)
	Register(Int32ToUint)
	Register(Int32ToUint16)
	Register(Int32ToUint16)
	Register(Int32ToUint32)
	Register(Int32ToUint32)
	Register(Int32ToUint64)
	Register(Int32ToUint64)
	Register(Int32ToUint8)
	Register(Int32ToUint8)
	Register(Int32ToUint8)
	Register(Int32ToUint8)
	Register(Int32ToUintptr)
	Register(Int32ToUintptr)
	Register(Int64ToDuration)
	Register(Int64ToFloat32)
	Register(Int64ToFloat64)
	Register(Int64ToInt)
	Register(Int64ToInt16)
	Register(Int64ToInt32)
	Register(Int64ToInt32)
	Register(Int64ToInt64)
	Register(Int64ToInt8)
	Register(Int64ToUint)
	Register(Int64ToUint16)
	Register(Int64ToUint32)
	Register(Int64ToUint64)
	Register(Int64ToUint8)
	Register(Int64ToUint8)
	Register(Int64ToUintptr)
	Register(Int8ToDuration)
	Register(Int8ToFloat32)
	Register(Int8ToFloat64)
	Register(Int8ToInt)
	Register(Int8ToInt16)
	Register(Int8ToInt32)
	Register(Int8ToInt32)
	Register(Int8ToInt64)
	Register(Int8ToInt8)
	Register(Int8ToUint)
	Register(Int8ToUint16)
	Register(Int8ToUint32)
	Register(Int8ToUint64)
	Register(Int8ToUint8)
	Register(Int8ToUint8)
	Register(Int8ToUintptr)
	Register(IntToDuration)
	Register(IntToFloat32)
	Register(IntToFloat64)
	Register(IntToInt)
	Register(IntToInt16)
	Register(IntToInt32)
	Register(IntToInt32)
	Register(IntToInt64)
	Register(IntToInt8)
	Register(IntToUint)
	Register(IntToUint16)
	Register(IntToUint32)
	Register(IntToUint64)
	Register(IntToUint8)
	Register(IntToUint8)
	Register(IntToUintptr)
	Register(Uint16ToDuration)
	Register(Uint16ToFloat32)
	Register(Uint16ToFloat64)
	Register(Uint16ToInt)
	Register(Uint16ToInt16)
	Register(Uint16ToInt32)
	Register(Uint16ToInt32)
	Register(Uint16ToInt64)
	Register(Uint16ToInt8)
	Register(Uint16ToUint)
	Register(Uint16ToUint16)
	Register(Uint16ToUint32)
	Register(Uint16ToUint64)
	Register(Uint16ToUint8)
	Register(Uint16ToUint8)
	Register(Uint16ToUintptr)
	Register(Uint32ToDuration)
	Register(Uint32ToFloat32)
	Register(Uint32ToFloat64)
	Register(Uint32ToInt)
	Register(Uint32ToInt16)
	Register(Uint32ToInt32)
	Register(Uint32ToInt32)
	Register(Uint32ToInt64)
	Register(Uint32ToInt8)
	Register(Uint32ToUint)
	Register(Uint32ToUint16)
	Register(Uint32ToUint32)
	Register(Uint32ToUint64)
	Register(Uint32ToUint8)
	Register(Uint32ToUint8)
	Register(Uint32ToUintptr)
	Register(Uint64ToDuration)
	Register(Uint64ToFloat32)
	Register(Uint64ToFloat64)
	Register(Uint64ToInt)
	Register(Uint64ToInt16)
	Register(Uint64ToInt32)
	Register(Uint64ToInt32)
	Register(Uint64ToInt64)
	Register(Uint64ToInt8)
	Register(Uint64ToUint)
	Register(Uint64ToUint16)
	Register(Uint64ToUint32)
	Register(Uint64ToUint64)
	Register(Uint64ToUint8)
	Register(Uint64ToUint8)
	Register(Uint64ToUintptr)
	Register(Uint8ToDuration)
	Register(Uint8ToFloat32)
	Register(Uint8ToFloat32)
	Register(Uint8ToFloat64)
	Register(Uint8ToFloat64)
	Register(Uint8ToInt)
	Register(Uint8ToInt)
	Register(Uint8ToInt16)
	Register(Uint8ToInt16)
	Register(Uint8ToInt32)
	Register(Uint8ToInt32)
	Register(Uint8ToInt32)
	Register(Uint8ToInt32)
	Register(Uint8ToInt64)
	Register(Uint8ToInt64)
	Register(Uint8ToInt8)
	Register(Uint8ToInt8)
	Register(Uint8ToUint)
	Register(Uint8ToUint)
	Register(Uint8ToUint16)
	Register(Uint8ToUint16)
	Register(Uint8ToUint32)
	Register(Uint8ToUint32)
	Register(Uint8ToUint64)
	Register(Uint8ToUint64)
	Register(Uint8ToUint8)
	Register(Uint8ToUint8)
	Register(Uint8ToUint8)
	Register(Uint8ToUint8)
	Register(Uint8ToUintptr)
	Register(Uint8ToUintptr)
	Register(UintToDuration)
	Register(UintToFloat32)
	Register(UintToFloat64)
	Register(UintToInt)
	Register(UintToInt16)
	Register(UintToInt32)
	Register(UintToInt32)
	Register(UintToInt64)
	Register(UintToInt8)
	Register(UintToUint)
	Register(UintToUint16)
	Register(UintToUint32)
	Register(UintToUint64)
	Register(UintToUint8)
	Register(UintToUint8)
	Register(UintToUintptr)
	Register(UintptrToDuration)
	Register(UintptrToFloat32)
	Register(UintptrToFloat64)
	Register(UintptrToInt)
	Register(UintptrToInt16)
	Register(UintptrToInt32)
	Register(UintptrToInt32)
	Register(UintptrToInt64)
	Register(UintptrToInt8)
	Register(UintptrToUint)
	Register(UintptrToUint16)
	Register(UintptrToUint32)
	Register(UintptrToUint64)
	Register(UintptrToUint8)
	Register(UintptrToUint8)
	Register(UintptrToUintptr)
}
