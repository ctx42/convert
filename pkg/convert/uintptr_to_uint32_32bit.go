// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package convert

import (
	"fmt"
	"math"
)

// UintptrToUint32 converts a given numeric value to uint32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToUint32(value uintptr) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("uintptr %w for uint32", ErrInvRange)
	}
	return uint32(value), nil
}
