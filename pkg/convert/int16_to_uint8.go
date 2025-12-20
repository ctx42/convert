// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Int16ToUint8 converts a given numeric value to uint8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int16ToUint8(value int16) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, fmt.Errorf("int16 %w for uint8", ErrInvRange)
	}
	return uint8(value), nil
}
