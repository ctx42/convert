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
		assert.Same(t, cnv, registry.m[from][to].conv)
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
		assert.Same(t, cnv1, registry.m[from][to].conv)
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
