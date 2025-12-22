// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"time"
)

// UintToDuration converts a given numeric value to [time.Duration].
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintToDuration(value uint) (time.Duration, error) {
	val, err := UintToInt64(value)
	if err != nil {
		return 0, err
	}
	return time.Duration(val), nil
}
