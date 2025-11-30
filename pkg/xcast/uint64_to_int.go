// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Uint64ToInt converts a given numeric value to int.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint64ToInt(value uint64) (int, error) {
	if value > math.MaxInt {
		return 0, fmt.Errorf("uint64 %w for int", ErrInvRange)
	}
	return int(value), nil
}
