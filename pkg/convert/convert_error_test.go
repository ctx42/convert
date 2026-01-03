// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"errors"
	"testing"

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

func Test_Error_Error(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- Given ---
		err := NewError(errors.New("test error"), "byte", "int")

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
	e := NewError(err, "byte", "int")

	// --- When ---
	have := e.Unwrap()

	// --- Then ---
	assert.Same(t, err, have)
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
