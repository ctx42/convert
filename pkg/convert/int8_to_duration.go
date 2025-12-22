// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// Int8ToDuration converts a given numeric value to [time.Duration].
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int8ToDuration(value int8) (time.Duration, error) {
	return time.Duration(value), nil
}
