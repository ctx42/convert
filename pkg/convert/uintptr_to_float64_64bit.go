// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"fmt"
)

// UintptrToFloat64 converts a given numeric value to float64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToFloat64(value uintptr) (float64, error) {
	if value > Float64SafeIntMax {
		return 0, fmt.Errorf("uintptr %w for float64", ErrInvSafeRange)
	}
	return float64(value), nil
}
