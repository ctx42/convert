// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// UintptrToFloat32 converts a given numeric value to float32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToFloat32(value uintptr) (float32, error) {
	if value > Float32SafeIntMax {
		return 0, fmt.Errorf("uintptr %w for float32", ErrInvSafeRange)
	}
	return float32(value), nil
}
