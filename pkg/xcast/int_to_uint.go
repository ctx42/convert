// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// IntToUint converts a given numeric value to uint.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func IntToUint(value int) (uint, error) {
	if value < 0 {
		return 0, fmt.Errorf("int %w for uint", ErrInvRange)
	}
	return uint(value), nil
}
