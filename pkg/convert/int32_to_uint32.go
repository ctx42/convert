// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
)

// Int32ToUint32 converts a given numeric value to uint32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int32ToUint32(value int32) (uint32, error) {
	if value < 0 {
		return 0, fmt.Errorf("int32 %w for uint32", ErrInvRange)
	}
	return uint32(value), nil
}
