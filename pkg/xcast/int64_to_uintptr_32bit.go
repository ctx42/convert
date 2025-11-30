// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package xcast

import (
	"fmt"
	"math"
)

// Int64ToUintptr converts a given numeric value to uintptr.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int64ToUintptr(value int64) (uintptr, error) {
	if value < 0 {
		return 0, fmt.Errorf("int64 %w for uintptr", ErrInvRange)
	}
	if value > math.MaxUint {
		return 0, fmt.Errorf("int64 %w for uintptr", ErrInvRange)
	}
	return uintptr(value), nil
}
