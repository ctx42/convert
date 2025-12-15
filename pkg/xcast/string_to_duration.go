// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"fmt"
	"time"
)

// StringToDuration returns a parser function that converts a string to a
// [time.Duration]. If parsing fails, the returned function yields a zero
// [time.Duration] and an error describing the issue.
func StringToDuration(value string) (time.Duration, error) {
	if value == "" {
		format := "cannot convert (parse) an empty string to time.Duration: %w"
		return 0, fmt.Errorf(format, ErrInvValue)
	}
	t, err := time.ParseDuration(value)
	if err != nil {
		format := "cannot convert (parse) %q as time.Duration: %w"
		return 0, fmt.Errorf(format, value, ErrInvValue)
	}
	return t, nil
}
