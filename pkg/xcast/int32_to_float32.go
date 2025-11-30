// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// Int32ToFloat32 converts a given numeric value to float32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int32ToFloat32(value int32) (float32, error) {
	if value < Float32SafeIntMin {
		return 0, fmt.Errorf("int32 %w for float32", ErrInvSafeRange)
	}
	if value > Float32SafeIntMax {
		return 0, fmt.Errorf("int32 %w for float32", ErrInvSafeRange)
	}
	return float32(value), nil
}
