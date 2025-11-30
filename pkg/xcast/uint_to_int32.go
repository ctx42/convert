// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// UintToInt32 converts a given numeric value to int32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintToInt32(value uint) (int32, error) {
	if value > math.MaxInt32 {
		return 0, fmt.Errorf("uint %w for int32", ErrInvRange)
	}
	return int32(value), nil
}
