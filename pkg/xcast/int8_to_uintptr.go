// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
)

// Int8ToUintptr converts a given numeric value to uintptr.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int8ToUintptr(value int8) (uintptr, error) {
	if value < 0 {
		return 0, fmt.Errorf("int8 %w for uintptr", ErrInvRange)
	}
	return uintptr(value), nil
}
