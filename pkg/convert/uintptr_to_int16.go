// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// UintptrToInt16 converts a given numeric value to int16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToInt16(value uintptr) (int16, error) {
	if value > math.MaxInt16 {
		return 0, fmt.Errorf("uintptr %w for int16", ErrInvRange)
	}
	return int16(value), nil
}
