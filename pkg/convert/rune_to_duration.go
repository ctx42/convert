// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// RuneToDuration converts a given numeric value to [time.Duration].
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func RuneToDuration(value rune) (time.Duration, error) {
	return Int32ToDuration(value)
}
