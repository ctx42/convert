// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"reflect"
)

// typUintptr is reflected uint.
var typUintptr = reflect.TypeFor[uintptr]()

// AnyToUintptr converts the given value to uintptr using the package-level
// registry.
func AnyToUintptr(value any) (uintptr, error) {
	return AnyToUintptrUsing(registry, value)
}

// AnyToUintptrUsing converts the given value to uintptr using the provided
// registry.
func AnyToUintptrUsing(reg *Registry, value any) (uintptr, error) {
	from := reflect.TypeOf(value)
	wrp := reg.lookup(from, typUintptr)
	if wrp == nil {
		format := "%w: from %T to uintptr"
		return 0, fmt.Errorf(format, ErrUnkConv, value)
	}
	ret, err := wrp.cst(value)
	if err != nil {
		return 0, err
	}
	return ret.(uintptr), nil // nolint: forcetypeassert
}
