// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"time"
)

// StringToTime returns a parser function that converts a string to a
// [time.Time] according to the specified layout. The layout must follow
// [time.Parse] rules. If parsing fails, the returned function yields a zero
// [time.Time] and an error describing the issue.
func StringToTime(format string) func(value string) (time.Time, error) {
	return func(value string) (time.Time, error) {
		if value == "" {
			format := "cannot convert (parse) an empty string to time.Time: %w"
			return time.Time{}, fmt.Errorf(format, ErrInvValue)
		}
		t, err := time.Parse(format, value)
		if err != nil {
			return time.Time{}, fmt.Errorf("%w: %w", ErrInvValue, err)
		}
		return t, nil
	}
}
