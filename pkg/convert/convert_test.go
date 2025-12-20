// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
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
		assert.Same(t, IntToInt, have)
	})
}

func Test_ToAnyAny(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		cnv := ToAnyAny(Float64ToInt)

		// --- Then ---
		have, err := cnv(float64(42))
		assert.NoError(t, err)
		assert.Equal(t, 42, have)
	})

	t.Run("invalid `from` type", func(t *testing.T) {
		// --- When ---
		cnv := ToAnyAny(Float64ToInt)

		// --- Then ---
		have, err := cnv(42)
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorEqual(t, "invalid type: expected float64, got int", err)
		assert.Equal(t, 0, have)
	})
}

func Test_init(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		// --- Given ---
		types := SupportedTypes()

		// --- When ---
		for _, from := range types {
			for _, to := range types {
				assert.NotNil(t, registry.lookup(from, to), "%s -> %s", from, to)
			}
		}
	})

	t.Run("StringToString", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[string]()
		to := reflect.TypeFor[string]()

		// --- When ---
		have := registry.lookup(from, to)

		// --- Then ---
		assert.Equal(t, from, have.from)
		assert.Equal(t, to, have.to)
		assert.Same(t, StringToString, have.cnv)
	})

	t.Run("StringToTime", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[string]()
		to := reflect.TypeFor[time.Time]()

		// --- When ---
		have := registry.lookup(from, to)

		// --- Then ---
		assert.Equal(t, from, have.from)
		assert.Equal(t, to, have.to)

		cnv := have.cnv.(FromTo[string, time.Time])
		hTim, err := cnv("2000-01-02T03:04:05Z")
		assert.NoError(t, err)
		wTim := time.Date(2000, time.January, 2, 3, 4, 5, 0, time.UTC)
		assert.Exact(t, wTim, hTim)
	})

	t.Run("StringToDuration", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[string]()
		to := reflect.TypeFor[time.Duration]()

		// --- When ---
		have := registry.lookup(from, to)

		// --- Then ---
		assert.Equal(t, from, have.from)
		assert.Equal(t, to, have.to)
		assert.Same(t, StringToDuration, have.cnv)
	})
}
