// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Uint64ToInt64 converts a given numeric value to int64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint64ToInt64(value uint64) (int64, error) {
	if value > math.MaxInt64 {
		return 0, fmt.Errorf("uint64 %w for int64", ErrInvRange)
	}
	return int64(value), nil
}
