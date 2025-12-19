// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv

import (
	"reflect"
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/convert/pkg/xcast"
)

func Test_Register(t *testing.T) {
	t.Run("register new", func(t *testing.T) {
		// --- Given ---
		type A struct{}
		type B struct{}
		cnv := func(from A) (to B, err error) { return to, nil }

		// --- When ---
		have := Register(cnv)

		// --- Then ---
		assert.Nil(t, have)
		from, to := reflect.TypeFor[A](), reflect.TypeFor[B]()
		assert.Same(t, cnv, registry.m[from][to].cnv)
	})

	t.Run("register existing", func(t *testing.T) {
		// --- Given ---
		type A struct{}
		type B struct{}
		cnv0 := func(from A) (to B, err error) { return to, nil }
		cnv1 := func(from A) (to B, err error) { return to, nil }
		Register(cnv0)

		// --- When ---
		have := Register(cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)
		from, to := reflect.TypeFor[A](), reflect.TypeFor[B]()
		assert.Same(t, cnv1, registry.m[from][to].cnv)
	})
}

func Test_Lookup(t *testing.T) {
	t.Run("lookup not registered", func(t *testing.T) {
		// --- Given ---
		type A struct{}
		type B struct{}

		// --- When ---
		have := Lookup[A, B]()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("lookup registered", func(t *testing.T) {
		// --- When ---
		have := Lookup[int, int]()

		// --- Then ---
		assert.Same(t, xcast.IntToInt, have)
	})
}

func Test_ConverterToCaster(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		cnv := ConverterToCaster(xcast.Float64ToInt)

		// --- Then ---
		have, err := cnv(float64(42))
		assert.NoError(t, err)
		assert.Equal(t, 42, have)
	})

	t.Run("invalid `from` type", func(t *testing.T) {
		// --- When ---
		cnv := ConverterToCaster(xcast.Float64ToInt)

		// --- Then ---
		have, err := cnv(42)
		assert.ErrorIs(t, xcast.ErrInvType, err)
		assert.ErrorEqual(t, "invalid type: expected float64, got int", err)
		assert.Equal(t, 0, have)
	})
}
