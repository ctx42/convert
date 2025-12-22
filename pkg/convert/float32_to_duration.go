// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// Float32ToDuration converts a given numeric value to [time.Duration].
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Float32ToDuration(value float32) (time.Duration, error) {
	val, err := Float32ToInt64(value)
	if err != nil {
		return 0, err
	}
	return time.Duration(val), nil
}
