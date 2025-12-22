// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// DurationToUint32 converts a given [time.Duration] value to uint32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func DurationToUint32(value time.Duration) (uint32, error) {
	return Int64ToUint32(int64(value))
}
