// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// IntToUint16 converts a given numeric value to uint16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func IntToUint16(value int) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, fmt.Errorf("int %w for uint16", ErrInvRange)
	}
	return uint16(value), nil
}
