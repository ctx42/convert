// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Uint64ToUint16 converts a given numeric value to uint16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint64ToUint16(value uint64) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("uint64 %w for uint16", ErrInvRange)
	}
	return uint16(value), nil
}
