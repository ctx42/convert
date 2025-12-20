// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float64ToUint64 converts a given numeric value to uint64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToUint64(value float64) (uint64, error) {
	if value != math.Trunc(value) {
		wMsg := "uint64 requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < 0 {
		return 0, fmt.Errorf("float64 %w for uint64", ErrInvRange)
	}
	if value > Float64SafeIntMax {
		return 0, fmt.Errorf("float64 %w for uint64", ErrInvSafeRange)
	}
	return uint64(value), nil
}
