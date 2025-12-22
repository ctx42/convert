// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// UintptrToDuration converts a given numeric value to [time.Duration].
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func UintptrToDuration(value uintptr) (time.Duration, error) {
	val, err := UintptrToInt64(value)
	if err != nil {
		return 0, err
	}
	return time.Duration(val), nil
}
