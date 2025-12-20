// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// UintptrToInt8 converts a given numeric value to int8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToInt8(value uintptr) (int8, error) {
	if value > math.MaxInt8 {
		return 0, fmt.Errorf("uintptr %w for int8", ErrInvRange)
	}
	return int8(value), nil
}
