// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package xcast

import (
	"fmt"
	"math"
)

// Float64ToUint converts a given numeric value to uint.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToUint(value float64) (uint, error) {
	if value != math.Trunc(value) {
		wMsg := "uint requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)
	}
	if value < 0 {
		return 0, fmt.Errorf("float64 %w for uint", ErrInvRange)
	}
	if value > Float64SafeIntMax {
		return 0, fmt.Errorf("float64 %w for uint", ErrInvSafeRange)
	}
	return uint(value), nil
}
