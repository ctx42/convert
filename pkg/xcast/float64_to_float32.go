// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Float64ToFloat32 converts a given numeric value to float32. But this is one
// of those cases when it's easy to lose precision; therefore, only the casts
// of whole numbers are allowed in range from Float32SafeIntMin to
// Float32SafeIntMax.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToFloat32(value float64) (float32, error) {
	if value != math.Trunc(value) {
		wMsg := "non-integer float64 to float32: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if value < Float32SafeIntMin {
		return 0, fmt.Errorf("float64 %w for float32", ErrInvSafeRange)
	}
	if value > Float32SafeIntMax {
		return 0, fmt.Errorf("float64 %w for float32", ErrInvSafeRange)
	}
	return float32(value), nil
}
