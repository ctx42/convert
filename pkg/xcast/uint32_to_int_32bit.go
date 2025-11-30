// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package xcast

import (
	"fmt"
	"math"
)

// Uint32ToInt converts a given numeric value to int.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint32ToInt(value uint32) (int, error) {
	if value > math.MaxInt {
		return 0, fmt.Errorf("uint32 %w for int", ErrInvRange)
	}
	return int(value), nil
}
