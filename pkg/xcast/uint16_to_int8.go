// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"math"
)

// Uint16ToInt8 converts a given numeric value to int8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint16ToInt8(value uint16) (int8, error) {
	if value > math.MaxInt8 {
		return 0, fmt.Errorf("uint16 %w for int8", ErrInvRange)
	}
	return int8(value), nil
}
