// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float64ToUint32 converts a given numeric value to uint32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToUint32(value float64) (uint32, error) {
	if value != math.Trunc(value) {
		wMsg := "uint32 requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < 0 || value > math.MaxUint32 {
		return 0, fmt.Errorf("float64 %w for uint32", ErrInvRange)
	}
	return uint32(value), nil
}
