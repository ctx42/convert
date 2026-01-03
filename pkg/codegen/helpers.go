// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"fmt"
	"reflect"
)

// IsSigned checks if the provided [Number] is a signed type.
func IsSigned[T Number]() bool {
	typ := reflect.TypeFor[T]()
	switch typ.Kind() {
	case reflect.Uint:
		return false

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return false

	case reflect.Uintptr:
		return false

	case reflect.Int:
		return true

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true

	case reflect.Float32, reflect.Float64:
		return true

	default:
		return false
	}
}

// IsFloat checks if the provided [Number] is a floating-point type.
func IsFloat[T Number]() bool {
	typ := reflect.TypeFor[T]()
	switch typ.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// MinValue returns a [Value] representing the minimum value for the given type.
func MinValue(typ Type) *Value {
	if typ.IsInteger() {
		return MinInteger(typ.size, typ.IsSigned())
	}
	return MinSafeFloat(typ.Size())
}

// MaxValue returns a [Value] representing the maximum value for the given type.
func MaxValue(typ Type) *Value {
	if typ.IsInteger() {
		return MaxInteger(typ.size, typ.IsSigned())
	}
	return MaxSafeFloat(typ.Size())
}

// MinInteger returns a [Value] representing a minimum integer value of the
// given size.
func MinInteger(size int, signed bool) *Value {
	if signed {
		return NewValue("math", fmt.Sprintf("MinInt%d", size))
	}
	return NewValue("0")
}

// MaxInteger returns a [Value] representing a maximum integer value of the
// given size.
func MaxInteger(size int, signed bool) *Value {
	if signed {
		return NewValue("math", fmt.Sprintf("MaxInt%d", size))
	}
	return NewValue("math", fmt.Sprintf("MaxUint%d", size))
}

// MinSafeFloat returns a [Value] representing a minimum safe integer value for
// the given floating-point size.
//
// See:
//   - [convert.Float32SafeIntMin]
//   - [convert.Float64SafeIntMin]
func MinSafeFloat(size int) *Value {
	return NewValue(fmt.Sprintf("Float%dSafeIntMin", size))
}

// MaxSafeFloat returns a [Value] representing a maximum safe integer value for
// the given floating-point size.
//
// See:
//   - [convert.Float32SafeIntmax]
//   - [convert.Float64SafeIntmax]
func MaxSafeFloat(size int) *Value {
	return NewValue(fmt.Sprintf("Float%dSafeIntMax", size))
}
