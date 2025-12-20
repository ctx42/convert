// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Uint16ToInt16 converts a given numeric value to int16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint16ToInt16(value uint16) (int16, error) {
	if value > math.MaxInt16 {
		return 0, fmt.Errorf("uint16 %w for int16", ErrInvRange)
	}
	return int16(value), nil
}
