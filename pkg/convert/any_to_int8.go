// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typInt8 is reflected int8.
var typInt8 = reflect.TypeFor[int8]()

// AnyToInt8 converts the given value to int8 using the package-level registry.
func AnyToInt8(value any) (int8, error) {
	return AnyToInt8Using(registry, value)
}

// AnyToInt8Using converts the given value to int8 using the provided registry.
func AnyToInt8Using(reg *Registry, value any) (int8, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typInt8)
	if wrp == nil {
		format := "%w: from %T to int8"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(int8), nil // nolint: forcetypeassert
}
