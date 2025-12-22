// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// DurationToInt16 converts a given [time.Duration] value to int16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func DurationToInt16(value time.Duration) (int16, error) {
	return Int64ToInt16(int64(value))
}
