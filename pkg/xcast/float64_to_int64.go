// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float64ToInt64 converts a given numeric value to int64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToInt64(value float64) (int64, error) {
	if value != math.Trunc(value) {
		wMsg := "int64 requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < Float64SafeIntMin {
		return 0, fmt.Errorf("float64 %w for int64", ErrInvSafeRange)
	}
	if value > Float64SafeIntMax {
		return 0, fmt.Errorf("float64 %w for int64", ErrInvSafeRange)
	}
	return int64(value), nil
}
