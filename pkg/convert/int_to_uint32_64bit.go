// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"fmt"
	"math"
)

// IntToUint32 converts a given numeric value to uint32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func IntToUint32(value int) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, fmt.Errorf("int %w for uint32", ErrInvRange)
	}
	return uint32(value), nil
}
