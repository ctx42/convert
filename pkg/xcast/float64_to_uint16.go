// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float64ToUint16 converts a given numeric value to uint16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToUint16(value float64) (uint16, error) {
	if value != math.Trunc(value) {
		wMsg := "uint16 requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < 0 || value > math.MaxUint16 {
		return 0, fmt.Errorf("float64 %w for uint16", ErrInvRange)
	}
	return uint16(value), nil
}
