// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// IntToUint64 converts a given numeric value to uint64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func IntToUint64(value int) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("int %w for uint64", ErrInvRange)
	}
	return uint64(value), nil
}
