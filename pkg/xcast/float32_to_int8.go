// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float32ToInt8 converts a given numeric value to int8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToInt8(value float32) (int8, error) {
	f64 := float64(value)
	if f64 != math.Trunc(f64) {
		wMsg := "int8 requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if f64 < math.MinInt8 || f64 > math.MaxInt8 {
		return 0, fmt.Errorf("float32 %w for int8", ErrInvRange)
	}
	return int8(f64), nil
}
