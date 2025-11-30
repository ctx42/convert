// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// Int16ToUint16 converts a given numeric value to uint16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int16ToUint16(value int16) (uint16, error) {
	if value < 0 {
		return 0, fmt.Errorf("int16 %w for uint16", ErrInvRange)
	}
	return uint16(value), nil
}
