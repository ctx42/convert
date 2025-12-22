// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typInt32 is reflected int32.
var typInt32 = reflect.TypeFor[int32]()

// AnyToInt32 converts the given value to int32 using the package-level
// registry.
func AnyToInt32(value any) (int32, error) {
	return AnyToInt32Using(registry, value)
}

// AnyToInt32Using converts the given value to int32 using the provided
// registry.
func AnyToInt32Using(reg *Registry, value any) (int32, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typInt32)
	if wrp == nil {
		format := "%w: from %T to int32"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(int32), nil // nolint: forcetypeassert
}
