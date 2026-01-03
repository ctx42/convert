// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// StringToDuration returns a parser function that converts a string to a
// [time.Duration]. If parsing fails, the returned function yields a zero
// [time.Duration] and an error describing the issue.
func StringToDuration(src string) (time.Duration, error) {
	if src == "" {
		return 0, NewError(ErrInvValue, "string", "time.Duration")
	}
	dst, err := time.ParseDuration(src)
	if err != nil {
		return 0, NewError(ErrInvValue, "string", "time.Duration")
	}
	return dst, nil
}
