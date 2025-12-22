// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typUint is reflected uint.
var typUint = reflect.TypeFor[uint]()

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
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uint), nil // nolint: forcetypeassert
}
