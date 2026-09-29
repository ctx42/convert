// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewError(t *testing.T) {
	// --- Given ---
	err := errors.New("test error")

	// --- When ---
	have := NewError(err, "byte", "int")

	// --- Then ---
	assert.Equal(t, "byte", have.Src)
	assert.Equal(t, "int", have.Dst)
	assert.Same(t, err, have.Err)
	assert.Nil(t, have.Cause)
}

func Test_Error_Format(t *testing.T) {
	// --- Given ---
	err := NewError(errors.New("test error"), "byte", "int")

	// --- When ---
	have := err.Format("err: %v, src: %v, dst: %v")

	// --- Then ---
	assert.Equal(t, "err: %v, src: %v, dst: %v", have.Fmt)
	assert.NotEqual(t, err, have)
	assert.Equal(t, "err: test error, src: byte, dst: int", have.Error())
}

func Test_Error_WithCause(t *testing.T) {
	// --- Given ---
	err := NewError(errors.New("test error"), "byte", "int")
	cause := errors.New("cause")

	// --- When ---
	have := err.WithCause(cause)

	// --- Then ---
	assert.Same(t, cause, have.Cause)
	assert.Nil(t, err.Cause)
}

func Test_Error_Error(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- Given ---
		err := NewError(errors.New("test error"), "byte", "int")

		// --- When ---
		have := err.Error()

		// --- Then ---
		assert.Equal(t, "test error: from byte to int", have)
	})

	t.Run("with cause", func(t *testing.T) {
		// --- Given ---
		err := NewError(errors.New("test error"), "byte", "int").
			WithCause(errors.New("cause"))

		// --- When ---
		have := err.Error()

		// --- Then ---
		assert.Equal(t, "test error: from byte to int", have)
	})

	t.Run("custom format", func(t *testing.T) {
		// --- Given ---
		err := NewError(errors.New("test error"), "byte", "int").
			Format("err: %v, src: %v, dst: %v")

		// --- When ---
		have := err.Error()

		// --- Then ---
		assert.Equal(t, "err: test error, src: byte, dst: int", have)
	})
}

func Test_Error_Unwrap(t *testing.T) {
	// --- Given ---
	err := errors.New("test error")
	e := NewError(err, "byte", "int").WithCause(errors.New("cause"))

	// --- When ---
	have := e.Unwrap()

	// --- Then ---
	assert.Same(t, err, have)
}

func Test_Error_Is(t *testing.T) {
	t.Run("matches cause", func(t *testing.T) {
		// --- Given ---
		cause := errors.New("cause")
		e := NewError(ErrInvValue, "byte", "int").WithCause(cause)

		// --- When ---
		have := e.Is(cause)

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("matches wrapped cause", func(t *testing.T) {
		// --- Given ---
		cause := errors.New("cause")
		wrapped := fmt.Errorf("wrap: %w", cause)
		e := NewError(ErrInvValue, "byte", "int").WithCause(wrapped)

		// --- When ---
		have := e.Is(cause)

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("no match", func(t *testing.T) {
		// --- Given ---
		e := NewError(ErrInvValue, "byte", "int").
			WithCause(errors.New("cause"))
		target := errors.New("other")

		// --- When ---
		have := e.Is(target)

		// --- Then ---
		assert.False(t, have)
	})

	t.Run("no cause", func(t *testing.T) {
		// --- Given ---
		e := NewError(ErrInvValue, "byte", "int")

		// --- When ---
		have := e.Is(ErrInvValue)

		// --- Then ---
		assert.False(t, have)
	})

	t.Run("errors Is matches sentinel and cause", func(t *testing.T) {
		// --- Given ---
		cause := errors.New("cause")
		err := error(NewError(ErrInvValue, "byte", "int").WithCause(cause))

		// --- When ---
		hSentinel := errors.Is(err, ErrInvValue)
		hCause := errors.Is(err, cause)

		// --- Then ---
		assert.True(t, hSentinel)
		assert.True(t, hCause)
	})
}

func Test_Error_As(t *testing.T) {
	t.Run("matches cause", func(t *testing.T) {
		// --- Given ---
		cause := &time.ParseError{Message: "cause"}
		e := NewError(ErrInvValue, "byte", "int").WithCause(cause)
		var target *time.ParseError

		// --- When ---
		have := e.As(&target)

		// --- Then ---
		assert.True(t, have)
		assert.Same(t, cause, target)
	})

	t.Run("no match", func(t *testing.T) {
		// --- Given ---
		e := NewError(ErrInvValue, "byte", "int").
			WithCause(errors.New("cause"))
		var target *time.ParseError

		// --- When ---
		have := e.As(&target)

		// --- Then ---
		assert.False(t, have)
		assert.Nil(t, target)
	})

	t.Run("no cause", func(t *testing.T) {
		// --- Given ---
		e := NewError(ErrInvValue, "byte", "int")
		var target *time.ParseError

		// --- When ---
		have := e.As(&target)

		// --- Then ---
		assert.False(t, have)
		assert.Nil(t, target)
	})
}

func Test_ChangeErrDstName(t *testing.T) {
	t.Run("change success", func(t *testing.T) {
		// --- Given ---
		err := NewError(errors.New("test error"), "byte", "int")

		// --- When ---
		have := ChangeErrDstName(err, "other")

		// --- Then ---
		assert.ErrorEqual(t, "test error: from byte to other", have)
	})

	t.Run("keeps cause", func(t *testing.T) {
		// --- Given ---
		cause := errors.New("cause")
		err := NewError(ErrInvValue, "byte", "int").WithCause(cause)

		// --- When ---
		have := ChangeErrDstName(err, "other")

		// --- Then ---
		assert.ErrorEqual(t, "invalid value: from byte to other", have)
		assert.ErrorIs(t, cause, have)
	})

	t.Run("nil error", func(t *testing.T) {
		// --- When ---
		have := ChangeErrDstName(nil, "other")

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("change success", func(t *testing.T) {
		// --- Given ---
		err := errors.New("test error")

		// --- When ---
		have := ChangeErrDstName(err, "other")

		// --- Then ---
		assert.Same(t, err, have)
	})
}
