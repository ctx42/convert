// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Uint32ToUint8 converts a given numeric value to uint8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint32ToUint8(value uint32) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("uint32 %w for uint8", ErrInvRange)
	}
	return uint8(value), nil
}
