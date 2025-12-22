// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// DurationToUint64 converts a given [time.Duration] value to uint64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func DurationToUint64(value time.Duration) (uint64, error) {
	return Int64ToUint64(int64(value))
}
