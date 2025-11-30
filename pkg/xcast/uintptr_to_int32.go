// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// UintptrToInt32 converts a given numeric value to int32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToInt32(value uintptr) (int32, error) {
	if value > uintptr(math.MaxInt32) {
		return 0, fmt.Errorf("uintptr %w for int32", ErrInvRange)
	}
	return int32(value), nil
}
