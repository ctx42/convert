// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"reflect"
	"time"
)

// typDuration is reflected [time.Duration].
var typDuration = reflect.TypeFor[time.Duration]()

// AnyToDuration converts the given value to [time.Duration] using the
// package-level registry.
func AnyToDuration(value any) (time.Duration, error) {
	val, err := AnyToInt64Using(registry, value)
	if err != nil {
		return 0, err
	}
	return time.Duration(val), nil
}

// AnyToDurationUsing converts the given value to [time.Duration] using the
// provided registry.
func AnyToDurationUsing(reg *Registry, value any) (time.Duration, error) {
	val, err := AnyToInt64Using(reg, value)
	if err != nil {
		return 0, err
	}
	return time.Duration(val), nil
}
