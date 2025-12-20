// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float32ToInt64 converts a given numeric value to int64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToInt64(value float32) (int64, error) {
	f64 := float64(value)
	if f64 != math.Trunc(f64) {
		wMsg := "int64 requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if value < Float32SafeIntMin {
		return 0, fmt.Errorf("float32 %w for int64", ErrInvSafeRange)
	}
	if value > Float32SafeIntMax {
		return 0, fmt.Errorf("float32 %w for int64", ErrInvSafeRange)
	}
	return int64(f64), nil
}
