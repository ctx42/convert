// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// DurationToInt32 converts a given [time.Duration] value to int32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func DurationToInt32(value time.Duration) (int32, error) {
	return Int64ToInt32(int64(value))
}
