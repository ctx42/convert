// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float64ToInt8 converts a given numeric value to int8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToInt8(value float64) (int8, error) {
	if value != math.Trunc(value) {
		wMsg := "int8 requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, fmt.Errorf("float64 %w for int8", ErrInvRange)
	}
	return int8(value), nil
}
