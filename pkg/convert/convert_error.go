// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
)

// Error represents conversion error.
type Error struct {
	Err error  // Underlying error.
	Fmt string // Format string.
	Src any    // Source type name.
	Dst any    // Destination type name.
}

// NewError constructs new [Error] instance.
func NewError(err error, src, dst any) Error {
	return Error{
		Err: err,
		Fmt: "%v: from %v to %v",
		Src: src,
		Dst: dst,
	}
}

// Format customize error format string.
//
// The [Error.Error] method uses [fmt.Sprintf] to construct the error message
// and always passes format arguments in order: [Error.Err], [Error.Src],
// [Error.Dst].
func (e Error) Format(format string) Error {
	e.Fmt = format
	return e
}

func (e Error) Error() string { return fmt.Sprintf(e.Fmt, e.Err, e.Src, e.Dst) }

func (e Error) Unwrap() error { return e.Err }

// ChangeErrDstName changes the destination type name if the error an instance
// of [Error]. Returns nil for nil error. Returns the original error if it's
// not an instance of [Error].
func ChangeErrDstName(err error, dst string) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(Error); ok { // nolint: errorlint
		e.Dst = dst
		return e
	}
	return err
}
