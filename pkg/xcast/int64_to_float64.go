// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// Int64ToFloat64 converts a given numeric value to float64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int64ToFloat64(value int64) (float64, error) {
	if value < Float64SafeIntMin {
		return 0, fmt.Errorf("float64 %w for float64", ErrInvSafeRange)
	}
	if value > Float64SafeIntMax {
		return 0, fmt.Errorf("float64 %w for float64", ErrInvSafeRange)
	}
	return float64(value), nil
}
