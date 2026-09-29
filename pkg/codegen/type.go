// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"fmt"
	"reflect"
)

// Type describes a type.
type Type struct {
	value   *Value // Represents the type.
	size    int    // Number of bits needed to store a value of the type.
	bits    int    // See [Type.Bits] method documentation.
	numeric bool   // Does the type represent a number?
	signed  bool   // Is the type signed number? False for non-numeric types.
	float   bool   // Does the type represent a floating-point number?
	alias   *Value // When set, the type is an alias or has a base type of.

	// Is the type size platform-dependent (int, uint, uintptr)?
	platform bool
}

// NumericType constructs a [Type] instance from the provided [Number] type.
func NumericType[T Number]() Type {
	typ := reflect.TypeFor[T]()
	size := int(typ.Size()) * 8
	signed := IsSigned[T]()
	float := IsFloat[T]()
	kind := typ.Kind()
	platform := kind == reflect.Int ||
		kind == reflect.Uint ||
		kind == reflect.Uintptr

	bits := size
	if float {
		// The float64 can exactly represent
		// all integers in the range [-2^53+1,2^53-1]
		bits = 53
		if size == 32 {
			// The float32 can exactly represent
			// all integers in the range [-2^24+1,2^24-1].
			bits = 24
		}
	} else if signed {
		bits--
	}

	return Type{
		value:    NewValue(typ.PkgPath(), typ.Name()),
		size:     size,
		bits:     bits,
		numeric:  true,
		signed:   signed,
		float:    float,
		platform: platform,
	}
}

// Alias creates a type alias for the current type.
//
// Requires at least one argument: the alias name.
//
// For non-builtin types, two arguments may be provided: the first is the
// package path, the second is the alias (type) name.
//
// Panics when no arguments or more than two arguments are provided.
//
// Examples:
//
//	NumericType[uint8]().Alias("byte")
//	NumericType[int64]().Alias("time", "Duration")
func (typ Type) Alias(arg string, args ...string) Type {
	if arg == "" || len(args) > 1 {
		panic("invalid number of arguments for Type.Alias")
	}
	typ.alias = typ.value
	typ.value = NewValue(arg, args...)
	return typ
}

// IsAlias returns true if the type is an alias.
func (typ Type) IsAlias() bool { return typ.alias != nil }

// Name returns the type name without a package name.
func (typ Type) Name() string { return typ.value.Name(false) }

// Title returns titled type name.
func (typ Type) Title() string { return typ.value.Name(true) }

// Doc returns the type name for documentation purposes.
func (typ Type) Doc() string {
	format := "%s"
	if len(typ.Imports()) > 0 {
		format = "[%s]"
	}
	return fmt.Sprintf(format, typ.Code())
}

// Code returns the type name with a package name.
func (typ Type) Code() string { return typ.value.Code() }

// Imports return the package import paths required for the type. Returns nil
// when the type doesn't require any import paths.
func (typ Type) Imports() []string { return typ.value.Imports() }

// IsNumeric returns true if the type is a numeric type.
func (typ Type) IsNumeric() bool { return typ.numeric }

// IsSigned returns true if the type is a numeric signed type.
func (typ Type) IsSigned() bool { return typ.signed }

// IsUnsigned returns true if the type is a numeric unsigned type.
func (typ Type) IsUnsigned() bool { return !typ.signed }

// IsFloat returns true if the type is a numeric float-point type.
func (typ Type) IsFloat() bool { return typ.numeric && typ.float }

// IsInteger returns true if the type is a numeric integer type.
func (typ Type) IsInteger() bool { return typ.numeric && !typ.float }

// IsPlatform returns true if the type size depends on the platform: 32 or 64
// bits, like int, uint, and uintptr.
func (typ Type) IsPlatform() bool { return typ.platform }

// Size returns the size of the type in bits.
func (typ Type) Size() int { return typ.size }

// Bits returns the number of bits used for values on each side of the zero
// value. For floating-point numbers this is set to the number of bits occupied
// by the maximum safe integer value that can be represented by the given
// floating-point number.
func (typ Type) Bits() int { return typ.bits }

// onPlatform returns the type as it is on a platform with the given word size
// in bits (32 or 64). Returns the type unchanged when its size is not
// platform-dependent.
func (typ Type) onPlatform(size int) Type {
	if !typ.platform {
		return typ
	}
	typ.size = size
	typ.bits = size
	if typ.signed {
		typ.bits--
	}
	return typ
}

// ConvActions returns the list of actions / checks needed for safe conversion
// from the current type to the target type on both 32-bit and 64-bit
// platforms. Method returns nil when conversion is not supported.
func (typ Type) ConvActions(target Type) []Action {
	if !typ.IsPlatform() && !target.IsPlatform() {
		return typ.convActions(target)
	}
	return mergeActions(
		target,
		typ.platformConvActions(target, 32),
		typ.platformConvActions(target, 64),
	)
}

// platformConvActions returns the list of actions / checks needed for safe
// conversion from the current type to the target type on a platform with the
// given word size in bits (32 or 64).
func (typ Type) platformConvActions(target Type, size int) []Action {
	return typ.onPlatform(size).convActions(target.onPlatform(size))
}

// convActions returns the list of actions / checks needed for safe conversion
// from the current type to the target type, assuming the type sizes are
// fixed. Method returns nil when conversion is not supported.
func (typ Type) convActions(target Type) []Action {
	// Conversions between non-numeric types are not supported.
	if !typ.IsNumeric() || !target.IsNumeric() {
		return nil
	}

	if typ.Code() == target.Code() {
		return []Action{NewAction(CastNotNeeded, nil)}
	}

	var actions []Action

	// Handle cases where at least one type is a float.
	if typ.IsFloat() || target.IsFloat() {
		return typ.floatConvActions(target)
	}

	if typ.IsSigned() && target.IsUnsigned() {
		actions = append(actions, NewAction(CheckIsNonNegative, nil))
	}

	if typ.Bits() > target.Bits() {
		if typ.IsSigned() && target.IsSigned() {
			tgtMin := MinInteger(target.Size(), target.IsSigned())
			actions = append(actions, NewAction(CheckUnderflows, tgtMin))
		}
		tgtMax := MaxInteger(target.Size(), target.IsSigned())
		actions = append(actions, NewAction(CheckOverflows, tgtMax))
	}
	return append(actions, NewAction(CastDirectly, nil))
}

// floatConvActions returns the list of actions / checks needed for safe
// conversion from the current type to the target type where one of the types
// is a floating-point number. Method returns nil when conversion is not
// supported.
//
// nolint: cyclop
func (typ Type) floatConvActions(target Type) []Action {
	// Conversions between non-numeric types are not supported.
	if !typ.IsNumeric() || !target.IsNumeric() {
		return nil
	}

	// One of the types must be a floating-point number.
	if !typ.IsFloat() && !target.IsFloat() {
		return nil
	}

	var actions []Action

	if typ.IsFloat() && target.IsFloat() {
		// Every value of a narrower floating-point type is exactly
		// representable by a wider one, including NaN and infinities.
		if typ.Bits() <= target.Bits() {
			return []Action{NewAction(CastDirectly, nil)}
		}

		name := fmt.Sprintf("MaxFloat%d", target.Size())
		return []Action{
			NewAction(CheckFloatRange, NewValue("math", name)),
			NewAction(CheckFloatExact, nil),
			NewAction(CastDirectly, nil),
		}
	}

	if target.IsInteger() {
		// Actions common to all float-to-integer conversions.
		actions = append(
			actions,
			NewAction(CastToFloat64, nil),
			NewAction(CheckIsNumber, nil),
			NewAction(CheckIsFinite, nil),
			NewAction(CheckIsWhole, nil),
		)

		if target.IsUnsigned() {
			actions = append(actions, NewAction(CheckIsNonNegative, nil))
		}

		if typ.Bits() < target.Bits() {
			if target.IsSigned() {
				minTgt := MinSafeFloat(typ.Size())
				actions = append(
					actions,
					NewAction(CheckFloatSafeToIntMin, minTgt),
				)
			}
			typMax := MaxSafeFloat(typ.Size())
			actions = append(
				actions,
				NewAction(CheckFloatSafeToIntMax, typMax),
			)
			return append(actions, NewAction(CastDirectly, nil))
		}

		if target.IsSigned() {
			tgtMin := MinInteger(target.Size(), target.IsSigned())
			actions = append(actions, NewAction(CheckUnderflows, tgtMin))
		}

		tgtMax := MaxInteger(target.Size(), target.IsSigned())
		return append(
			actions,
			NewAction(CheckOverflows, tgtMax),
			NewAction(CastDirectly, nil),
		)
	}

	// The integer-to-float conversion below.

	if typ.Bits() > target.Bits() {
		if typ.IsSigned() && target.IsSigned() {
			minTgt := MinSafeFloat(target.Size())
			actions = append(actions, NewAction(CheckIntSafeToFloatMin, minTgt))
		}
		typMax := MaxSafeFloat(target.Size())
		actions = append(actions, NewAction(CheckIntSafeToFloatMax, typMax))
	}

	return append(actions, NewAction(CastDirectly, nil))
}
