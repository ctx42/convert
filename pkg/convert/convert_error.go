// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"errors"
	"fmt"
)

// Error represents conversion error.
type Error struct {
	Err   error  // Underlying error.
	Cause error  // Optional error that caused the conversion to fail.
	Fmt   string // Format string.
	Src   any    // Source type name.
	Dst   any    // Destination type name.
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

// WithCause sets the error that caused the conversion to fail.
//
// The cause does not change the message returned by [Error.Error], but both
// [Error.Err] and the cause are matched by [errors.Is] and [errors.As]; see
// [Error.Is] and [Error.As].
func (e Error) WithCause(cause error) Error {
	e.Cause = cause
	return e
}

func (e Error) Error() string { return fmt.Sprintf(e.Fmt, e.Err, e.Src, e.Dst) }

func (e Error) Unwrap() error { return e.Err }

// Is reports whether the [Error.Cause] matches the target. It lets
// [errors.Is] match the cause, while [Error.Unwrap] exposes [Error.Err].
func (e Error) Is(target error) bool {
	return e.Cause != nil && errors.Is(e.Cause, target)
}

// As finds the first error in the [Error.Cause] tree that matches the target.
// It lets [errors.As] match the cause, while [Error.Unwrap] exposes
// [Error.Err].
func (e Error) As(target any) bool {
	return e.Cause != nil && errors.As(e.Cause, target)
}

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
