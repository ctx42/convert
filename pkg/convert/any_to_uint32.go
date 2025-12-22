// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typUint32 is reflected uint32.
var typUint32 = reflect.TypeFor[uint32]()

// AnyToUint32 converts the given value to uint32 using the package-level
// registry.
func AnyToUint32(value any) (uint32, error) {
	return AnyToUint32Using(registry, value)
}

// AnyToUint32Using converts the given value to uint32 using the provided
// registry.
func AnyToUint32Using(reg *Registry, value any) (uint32, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typUint32)
	if wrp == nil {
		format := "%w: from %T to uint32"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uint32), nil // nolint: forcetypeassert
}
