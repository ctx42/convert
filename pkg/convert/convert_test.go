// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
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
		cnv := func(src A) (dst B, err error) { return dst, nil }

		// --- When ---
		have := Register(cnv)

		// --- Then ---
		assert.Nil(t, have)
		src, dst := reflect.TypeFor[A](), reflect.TypeFor[B]()
		assert.Same(t, cnv, registry.m[src][dst].cnv)
	})

	t.Run("register existing", func(t *testing.T) {
		// --- Given ---
		type A struct{}
		type B struct{}
		cnv0 := func(src A) (dst B, err error) { return dst, nil }
		cnv1 := func(src A) (dst B, err error) { return dst, nil }
		Register(cnv0)

		// --- When ---
		have := Register(cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)
		src, dst := reflect.TypeFor[A](), reflect.TypeFor[B]()
		assert.Same(t, cnv1, registry.m[src][dst].cnv)
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
		assert.ErrorEqual(t, "invalid type: expected float64 got int", err)
		assert.Equal(t, 0, have)
	})
}

func Test_NewOptions(t *testing.T) {
	t.Run("without options", func(t *testing.T) {
		// --- When ---
		have := NewOptions()

		// --- Then ---
		assert.Same(t, registry, have.reg)
	})

	t.Run("with options", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := NewOptions(WithRegistry(reg))

		// --- Then ---
		assert.Same(t, reg, have.reg)
	})
}

func Test_WithRegistry(t *testing.T) {
	// --- Given ---
	ops := &Options{}
	reg := NewRegistry()

	// --- When ---
	WithRegistry(reg)(ops)

	// --- Then ---
	assert.Same(t, reg, ops.reg)
}

func Test_init(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		// --- Given ---
		types := SupportedTypes()

		// --- When ---
		for _, src := range types {
			for _, dst := range types {
				assert.NotNil(t, registry.lookup(src, dst), "%s -> %s", src, dst)
			}
		}
	})

	t.Run("BoolToBool", func(t *testing.T) {
		// --- Given ---
		src := reflect.TypeFor[bool]()
		dst := reflect.TypeFor[bool]()

		// --- When ---
		have := registry.lookup(src, dst)

		// --- Then ---
		assert.Equal(t, src, have.src)
		assert.Equal(t, dst, have.dst)
		assert.Same(t, BoolToBool, have.cnv)
	})

	t.Run("StringToString", func(t *testing.T) {
		// --- Given ---
		src := reflect.TypeFor[string]()
		dst := reflect.TypeFor[string]()

		// --- When ---
		have := registry.lookup(src, dst)

		// --- Then ---
		assert.Equal(t, src, have.src)
		assert.Equal(t, dst, have.dst)
		assert.Same(t, StringToString, have.cnv)
	})

	t.Run("StringToTime", func(t *testing.T) {
		// --- Given ---
		src := reflect.TypeFor[string]()
		dst := reflect.TypeFor[time.Time]()

		// --- When ---
		have := registry.lookup(src, dst)

		// --- Then ---
		assert.Equal(t, src, have.src)
		assert.Equal(t, dst, have.dst)

		cnv := have.cnv.(SrcToDst[string, time.Time])
		hTim, err := cnv("2000-01-02T03:04:05Z")
		assert.NoError(t, err)
		wTim := time.Date(2000, time.January, 2, 3, 4, 5, 0, time.UTC)
		assert.Exact(t, wTim, hTim)
	})

	t.Run("StringToDuration", func(t *testing.T) {
		// --- Given ---
		src := reflect.TypeFor[string]()
		dst := reflect.TypeFor[time.Duration]()

		// --- When ---
		have := registry.lookup(src, dst)

		// --- Then ---
		assert.Equal(t, src, have.src)
		assert.Equal(t, dst, have.dst)
		assert.Same(t, StringToDuration, have.cnv)
	})
}
