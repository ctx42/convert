// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float64ToUint8 converts a given numeric value to uint8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToUint8(value float64) (uint8, error) {
	if value != math.Trunc(value) {
		wMsg := "uint8 requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < 0 || value > math.MaxUint8 {
		return 0, fmt.Errorf("float64 %w for uint8", ErrInvRange)
	}
	return uint8(value), nil
}
