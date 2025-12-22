// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typFloat64 is reflected float64.
var typFloat64 = reflect.TypeFor[float64]()

// AnyToFloat64 converts the given value to float64 using the package-level
// registry.
func AnyToFloat64(value any) (float64, error) {
	return AnyToFloat64Using(registry, value)
}

// AnyToFloat64Using converts the given value to float64 using the provided
// registry.
func AnyToFloat64Using(reg *Registry, value any) (float64, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typFloat64)
	if wrp == nil {
		format := "%w: from %T to float64"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(float64), nil // nolint: forcetypeassert
}
