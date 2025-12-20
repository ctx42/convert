// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float32ToInt32 converts a given numeric value to int32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToInt32(value float32) (int32, error) {
	f64 := float64(value)
	if f64 != math.Trunc(f64) {
		wMsg := "int32 requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if value < Float32SafeIntMin {
		return 0, fmt.Errorf("float32 %w for int32", ErrInvSafeRange)
	}
	if value > Float32SafeIntMax {
		return 0, fmt.Errorf("float32 %w for int32", ErrInvSafeRange)
	}
	return int32(f64), nil
}
