// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typInt64 is reflected int64.
var typInt64 = reflect.TypeFor[int64]()

// AnyToInt64 converts the given value to int64 using the package-level registry.
func AnyToInt64(value any) (int64, error) {
	return AnyToInt64Using(registry, value)
}

// AnyToInt64Using converts the given value to int64 using the provided registry.
func AnyToInt64Using(reg *Registry, value any) (int64, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typInt64)
	if wrp == nil {
		format := "%w: from %T to int64"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(int64), nil // nolint: forcetypeassert
}
