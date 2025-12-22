// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typUint64 is reflected uint64.
var typUint64 = reflect.TypeFor[uint64]()

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
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uint64), nil // nolint: forcetypeassert
}
