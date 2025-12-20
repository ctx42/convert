// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
)

// Int32ToUint64 converts a given numeric value to uint64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int32ToUint64(value int32) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("int32 %w for uint64", ErrInvRange)
	}
	return uint64(value), nil
}
