// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv

import (
	"fmt"
	"reflect"

	"github.com/ctx42/convert/pkg/xcast"
)

// Cached reflection types.
var (
	typInt     = reflect.TypeFor[int]()
	typInt64   = reflect.TypeFor[int64]()
	typUint    = reflect.TypeFor[uint]()
	typUint64  = reflect.TypeFor[uint64]()
	typFloat64 = reflect.TypeFor[float64]()
)

// AnyToInt converts the given value to int using the package-level registry.
func AnyToInt(value any) (int, error) {
	return AnyToIntUsing(registry, value)
}

// AnyToIntUsing converts the given value to int using the provided registry.
func AnyToIntUsing(reg *Registry, value any) (int, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typInt)
	if wrp == nil {
		format := "%w: from %T to int"
		return 0, fmt.Errorf(format, xcast.ErrUndConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(int), nil
}

// AnyToInt64 converts the given value to int64 using the package-level registry.
func AnyToInt64(value any) (int64, error) {
	return AnyToInt64Using(registry, value)
}

// AnyToInt64Using converts the given value to int64 using the provided registry.
func AnyToInt64Using(reg *Registry, value any) (int64, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typInt64)
	if wrp == nil {
		format := "%w: from %T to int64"
		return 0, fmt.Errorf(format, xcast.ErrUndConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(int64), nil
}

// AnyToUint converts the given value to uint using the package-level registry.
func AnyToUint(value any) (uint, error) {
	return AnyToUintUsing(registry, value)
}

// AnyToUintUsing converts the given value to uint using the provided registry.
func AnyToUintUsing(reg *Registry, value any) (uint, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typUint)
	if wrp == nil {
		format := "%w: from %T to uint"
		return 0, fmt.Errorf(format, xcast.ErrUndConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uint), nil
}

// AnyToUint64 converts the given value to uint64 using the package-level
// registry.
func AnyToUint64(value any) (uint64, error) {
	return AnyToUint64Using(registry, value)
}

// AnyToUint64Using converts the given value to uint64 using the provided
// registry.
func AnyToUint64Using(reg *Registry, value any) (uint64, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typUint64)
	if wrp == nil {
		format := "%w: from %T to uint64"
		return 0, fmt.Errorf(format, xcast.ErrUndConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uint64), nil
}

// AnyToFloat64 converts the given value to float64 using the package-level
// registry.
func AnyToFloat64(value any) (float64, error) {
	return AnyToFloat64Using(registry, value)
}

// AnyToFloat64Using converts the given value to float64 using the provided
// registry.
func AnyToFloat64Using(reg *Registry, value any) (float64, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typFloat64)
	if wrp == nil {
		format := "%w: from %T to float64"
		return 0, fmt.Errorf(format, xcast.ErrUndConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(float64), nil
}
