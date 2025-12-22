// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typInt is reflected int.
var typInt = reflect.TypeFor[int]()

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
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(int), nil // nolint: forcetypeassert
}
