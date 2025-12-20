// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"fmt"
	"math"
)

// UintptrToInt64 converts a given numeric value to int64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToInt64(value uintptr) (int64, error) {
	if value > math.MaxInt64 {
		return 0, fmt.Errorf("uintptr %w for int64", ErrInvRange)
	}
	return int64(value), nil
}
