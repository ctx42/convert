// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// DurationToUint8 converts a given [time.Duration] value to uint8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func DurationToUint8(value time.Duration) (uint8, error) {
	return Int64ToUint8(int64(value))
}
