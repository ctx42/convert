// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typUint8 is reflected uint8.
var typUint8 = reflect.TypeFor[uint8]()

// AnyToUint8 converts the given value to uint8 using the package-level
// registry.
func AnyToUint8(value any) (uint8, error) {
	return AnyToUint8Using(registry, value)
}

// AnyToUint8Using converts the given value to uint8 using the provided
// registry.
func AnyToUint8Using(reg *Registry, value any) (uint8, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typUint8)
	if wrp == nil {
		format := "%w: from %T to uint8"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uint8), nil // nolint: forcetypeassert
}
