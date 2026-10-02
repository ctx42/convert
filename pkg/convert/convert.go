// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
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
// of a value from the Src type to the Dst type. On success, it returns the
// converted value and a nil error. On failure (e.g., truncation, underflow,
// overflow, or semantic loss), it returns the zero value of Dst along with a
// non-nil error describing the issue.
type SrcToDst[Src, Dst any] func(Src) (Dst, error)

// AnyToAny is a non-generic version of [SrcToDst]. The behavior is exactly the
// same in terms of conversion and error handling.
type AnyToAny func(any) (any, error)

// Sentinel errors.
var (
	// ErrInvRange is returned when a value isn't within a valid range.
	ErrInvRange = errors.New("value out of range")

	// ErrInvSafeRange is returned when a value is within the valid range, but
	// converting it may lead to precision loss.
	ErrInvSafeRange = errors.New("value out of safe range")

	// ErrUnsType is returned when conversion for a type is not defined.
	ErrUnsType = errors.New("unsupported type")

	// ErrInvType is returned when a type is not valid in a given conversion
	// context.
	ErrInvType = errors.New("invalid type")

	// ErrInvValue is returned when a value is not valid in a given conversion
	// context.
	ErrInvValue = errors.New("invalid value")

	// ErrFraction is returned when a floating-point value has a fractional
	// part, but the conversion requires a whole number.
	ErrFraction = errors.New("must be a whole number")

	// ErrInvFormat is returned when a value's format is not valid in a given
	// conversion context.
	ErrInvFormat = errors.New("invalid format")

	// ErrUnkConv is returned when a conversion is undefined for given types.
	ErrUnkConv = errors.New("conversion undefined")

	// ErrUns represents an explicitly unsupported conversion.
	ErrUns = errors.New("unsupported conversion")

	// ErrNilRegistry is returned when a conversion is attempted with a nil
	// [Registry], for example one set with [WithRegistry].
	ErrNilRegistry = errors.New("nil registry")
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
// Conversions using a nil registry fail with [ErrNilRegistry].
func WithRegistry(reg *Registry) Option {
	return func(ops *Options) { ops.reg = reg }
}

// SupportedTypes returns a slice of [reflect.Type] instances containing all
// types for which the package provides built-in converters. Any pair of types
// from this list can be converted in either direction.
func SupportedTypes() []reflect.Type {
	return []reflect.Type{
		reflect.TypeFor[int](),
		reflect.TypeFor[int8](),
		reflect.TypeFor[int16](),
		reflect.TypeFor[int32](),
		reflect.TypeFor[int64](),
		reflect.TypeFor[uint](),
		reflect.TypeFor[uint8](),
		reflect.TypeFor[uint16](),
		reflect.TypeFor[uint32](),
		reflect.TypeFor[uint64](),
		reflect.TypeFor[float32](),
		reflect.TypeFor[float64](),
		reflect.TypeFor[uintptr](),
		reflect.TypeFor[time.Duration](),
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

// MaxUintptr is the maximum value of the uintptr type on the current platform.
// It is an untyped constant, like [math.MaxUint], so it can be compared with
// any numeric type able to represent it.
const MaxUintptr = 1<<(32<<(^uintptr(0)>>63)) - 1

func init() {
	// Converters implemented by hand.
	Register(BoolToBool)
	Register(StringToDuration)
	Register(StringToString)
	Register(StringToTime(time.RFC3339Nano))
	Register(StringToInt)
	Register(StringToInt8)
	Register(StringToInt16)
	Register(StringToInt32)
	Register(StringToInt64)
	Register(StringToUint)
	Register(StringToUint8)
	Register(StringToUint16)
	Register(StringToUint32)
	Register(StringToUint64)

	// Generated numeric converters.
	registerNumeric()
}
