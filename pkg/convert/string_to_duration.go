// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"time"
)

// StringToDuration converts a string to a [time.Duration]. If parsing fails,
// it returns a zero [time.Duration] and an error describing the issue.
func StringToDuration(src string) (time.Duration, error) {
	cnvErr := NewError(ErrInvValue, "string", "time.Duration")
	if src == "" {
		return 0, cnvErr
	}
	dst, err := time.ParseDuration(src)
	if err != nil {
		return 0, cnvErr.WithCause(err)
	}
	return dst, nil
}
