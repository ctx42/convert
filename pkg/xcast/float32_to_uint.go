// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float32ToUint converts a given numeric value to uint.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToUint(value float32) (uint, error) {
	f64 := float64(value)
	if f64 != math.Trunc(f64) {
		wMsg := "uint requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if f64 < 0 {
		return 0, fmt.Errorf("float32 %w for uint", ErrInvRange)
	}
	if value > Float32SafeIntMax {
		return 0, fmt.Errorf("float32 %w for uint", ErrInvSafeRange)
	}
	return uint(f64), nil
}
