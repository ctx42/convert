// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// DurationToInt64 converts a given [time.Duration] value to int64.
func DurationToInt64(value time.Duration) (int64, error) {
	return int64(value), nil
}
