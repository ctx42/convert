// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package convert

import (
	"fmt"
	"math"
)

// Float64ToUintptr converts a given numeric value to uintptr.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToUintptr(value float64) (uintptr, error) {
	if value != math.Trunc(value) {
		wMsg := "uintptr requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if value < 0 {
		return 0, fmt.Errorf("float64 %w for uintptr", ErrInvRange)
	}
	if value > math.MaxUint {
		return 0, fmt.Errorf("float64 %w for uintptr", ErrInvSafeRange)
	}
	return uintptr(value), nil
}
