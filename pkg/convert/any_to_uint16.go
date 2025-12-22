// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typUint16 is reflected uint16.
var typUint16 = reflect.TypeFor[uint16]()

// AnyToUint16 converts the given value to uint16 using the package-level
// registry.
func AnyToUint16(value any) (uint16, error) {
	return AnyToUint16Using(registry, value)
}

// AnyToUint16Using converts the given value to uint16 using the provided
// registry.
func AnyToUint16Using(reg *Registry, value any) (uint16, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typUint16)
	if wrp == nil {
		format := "%w: from %T to uint16"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uint16), nil // nolint: forcetypeassert
}
