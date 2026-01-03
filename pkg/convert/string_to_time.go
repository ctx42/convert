// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// StringToTime returns a parser function that converts a string to a
// [time.Time] according to the specified layout. The layout must follow
// [time.Parse] rules. If parsing fails, the returned function yields a zero
// [time.Time] and an error describing the issue.
func StringToTime(layout string) func(value string) (time.Time, error) {
	return func(src string) (time.Time, error) {
		if src == "" {
			return time.Time{}, NewError(ErrInvValue, "string", "time.Time")
		}
		dst, err := time.Parse(layout, src)
		if err != nil {
			return time.Time{}, NewError(ErrInvValue, "string", "time.Time")
		}
		return dst, nil
	}
}
