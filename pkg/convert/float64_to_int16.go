// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float64ToInt16 converts a given numeric value to int16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToInt16(value float64) (int16, error) {
	if value != math.Trunc(value) {
		wMsg := "int16 requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, fmt.Errorf("float64 %w for int16", ErrInvRange)
	}
	return int16(value), nil
}
