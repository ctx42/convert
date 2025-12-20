// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float32ToInt16 converts a given numeric value to int16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToInt16(value float32) (int16, error) {
	f64 := float64(value)
	if f64 != math.Trunc(f64) {
		wMsg := "int16 requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if f64 < math.MinInt16 || f64 > math.MaxInt16 {
		return 0, fmt.Errorf("float32 %w for int16", ErrInvRange)
	}
	return int16(f64), nil
}
