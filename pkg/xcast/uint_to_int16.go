// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// UintToInt16 converts a given numeric value to int16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintToInt16(value uint) (int16, error) {
	if value > math.MaxInt16 {
		return 0, fmt.Errorf("uint %w for int16", ErrInvRange)
	}
	return int16(value), nil
}
