// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float32ToUintptr converts a given numeric value to uintptr.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToUintptr(value float32) (uintptr, error) {
	f64 := float64(value)
	if f64 != math.Trunc(f64) {
		wMsg := "uintptr requires non-fractional float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if f64 < 0 {
		return 0, fmt.Errorf("float32 %w for uintptr", ErrInvRange)
	}
	if value > Float32SafeIntMax {
		return 0, fmt.Errorf("float32 %w for uintptr", ErrInvSafeRange)
	}
	return uintptr(f64), nil
}
