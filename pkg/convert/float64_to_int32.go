// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float64ToInt32 converts a given numeric value to int32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToInt32(value float64) (int32, error) {
	if value != math.Trunc(value) {
		wMsg := "int32 requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < math.MinInt32 {
		return 0, fmt.Errorf("float64 %w for int32", ErrInvRange)
	}
	if value > math.MaxInt32 {
		return 0, fmt.Errorf("float64 %w for int32", ErrInvRange)
	}
	return int32(value), nil
}
