// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typInt16 is reflected int16.
var typInt16 = reflect.TypeFor[int16]()

// AnyToInt16 converts the given value to int16 using the package-level
// registry.
func AnyToInt16(value any) (int16, error) {
	return AnyToInt16Using(registry, value)
}

// AnyToInt16Using converts the given value to int16 using the provided
// registry.
func AnyToInt16Using(reg *Registry, value any) (int16, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typInt16)
	if wrp == nil {
		format := "%w: from %T to int16"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(int16), nil // nolint: forcetypeassert
}
