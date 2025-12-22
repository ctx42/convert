// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typFloat32 is reflected float32.
var typFloat32 = reflect.TypeFor[float32]()

// AnyToFloat32 converts the given value to float32 using the package-level
// registry.
func AnyToFloat32(value any) (float32, error) {
	return AnyToFloat32Using(registry, value)
}

// AnyToFloat32Using converts the given value to float32 using the provided
// registry.
func AnyToFloat32Using(reg *Registry, value any) (float32, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typFloat32)
	if wrp == nil {
		format := "%w: from %T to float32"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(float32), nil // nolint: forcetypeassert
}
