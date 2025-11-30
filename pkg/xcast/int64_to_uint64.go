// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// Int64ToUint64 converts a given numeric value to uint64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int64ToUint64(value int64) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("int64 %w for uint64", ErrInvRange)
	}
	return uint64(value), nil
}
