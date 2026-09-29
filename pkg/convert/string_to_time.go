// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
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
		cnvErr := NewError(ErrInvValue, "string", "time.Time")
		if src == "" {
			return time.Time{}, cnvErr
		}
		dst, err := time.Parse(layout, src)
		if err != nil {
			return time.Time{}, cnvErr.WithCause(err)
		}
		return dst, nil
	}
}
