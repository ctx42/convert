// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"fmt"
	"math"
)

// Float64ToInt converts a given numeric value to int.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float64ToInt(value float64) (int, error) {
	if value != math.Trunc(value) {
		wMsg := "int requires non-fractional float64: %w"
		return 0, fmt.Errorf(wMsg, ErrInvValue)

	}
	if value < Float64SafeIntMin {
		return 0, fmt.Errorf("float64 %w for int", ErrInvSafeRange)
	}
	if value > Float64SafeIntMax {
		return 0, fmt.Errorf("float64 %w for int", ErrInvSafeRange)
	}
	return int(value), nil
}
