// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Float32ToUint16 converts a given numeric value to uint16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToUint16(value float32) (uint16, error) {
	f64 := float64(value)
	if f64 != math.Trunc(f64) {
		wMsg := "uint16 requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint16 {
		return 0, fmt.Errorf("float32 %w for uint16", ErrInvRange)
	}
	return uint16(f64), nil
}
