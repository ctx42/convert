// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv

import (
	"reflect"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewRegistry(t *testing.T) {
	// --- When ---
	have := NewRegistry()

	// --- Then ---
	assert.NotNil(t, have)
	assert.Nil(t, have.m)
}

func Test_Registry_register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[uint]()
		to := reflect.TypeFor[uint8]()
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		wrp := wrap(cnv)
		reg := &Registry{}

		// --- When ---
		have := reg.register(wrp)

		// --- Then ---
		assert.Nil(t, have)
		assert.Same(t, wrp, reg.m[from][to])
	})

	t.Run("register already registered pair", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[uint]()
		to := reflect.TypeFor[uint8]()
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		wrp0 := wrap(cnv)
		wrp1 := wrap(cnv)
		reg := &Registry{
			m: map[reflect.Type]map[reflect.Type]*wrapper{from: {to: wrp0}},
		}

		// --- When ---
		have := reg.register(wrp1)

		// --- Then ---
		assert.Same(t, have, wrp0)
		assert.Same(t, wrp1, reg.m[from][to])
	})
}

func Test_Registry_lookup(t *testing.T) {
	t.Run("empty registry", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[uint]()
		to := reflect.TypeFor[uint8]()
		reg := &Registry{}

		// --- When ---
		have := reg.lookup(from, to)

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("existing converter", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[uint]()
		to := reflect.TypeFor[uint8]()
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		wrp := wrap(cnv)
		reg := &Registry{
			m: map[reflect.Type]map[reflect.Type]*wrapper{from: {to: wrp}},
		}

		// --- When ---
		have := reg.lookup(from, to)

		// --- Then ---
		assert.Same(t, wrp, have)
	})

	t.Run("not existing converter - from", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[uint]()
		to := reflect.TypeFor[uint8]()
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		wrp := wrap(cnv)
		reg := &Registry{
			m: map[reflect.Type]map[reflect.Type]*wrapper{from: {to: wrp}},
		}

		// --- When ---
		have := reg.lookup(reflect.TypeFor[int](), to)

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("not existing converter - to", func(t *testing.T) {
		// --- Given ---
		from := reflect.TypeFor[uint]()
		to := reflect.TypeFor[uint8]()
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		wrp := wrap(cnv)
		reg := &Registry{
			m: map[reflect.Type]map[reflect.Type]*wrapper{from: {to: wrp}},
		}

		// --- When ---
		have := reg.lookup(from, reflect.TypeFor[int]())

		// --- Then ---
		assert.Nil(t, have)
	})
}

func Test_RegisterConverter(t *testing.T) {
	t.Run("nil converter", func(t *testing.T) {
		// --- Given ---
		reg := &Registry{}

		// --- When ---
		have := RegisterConverter[int, int](reg, nil)

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("register new converter", func(t *testing.T) {
		// --- Given ---
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		reg := &Registry{}

		// --- When ---
		have := RegisterConverter(reg, cnv)

		// --- Then ---
		assert.Nil(t, have)
		from, to := reflect.TypeFor[uint](), reflect.TypeFor[uint8]()
		assert.Same(t, cnv, reg.m[from][to].cnv)
	})

	t.Run("register already registered converter", func(t *testing.T) {
		// --- Given ---
		cnv0 := func(from uint) (uint8, error) { return uint8(from), nil }
		cnv1 := func(from uint) (uint8, error) { return uint8(from), nil }
		reg := NewRegistry()
		reg.register(wrap(cnv0))

		// --- When ---
		have := RegisterConverter(reg, cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)
		from, to := reflect.TypeFor[uint](), reflect.TypeFor[uint8]()
		assert.Same(t, cnv1, reg.m[from][to].cnv)
	})

	t.Run("it should never happen", func(t *testing.T) {
		// --- Given ---
		wrongTypes := func(from float32) (int, error) { return 0, nil }
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		reg := &Registry{
			m: map[reflect.Type]map[reflect.Type]*wrapper{
				reflect.TypeFor[uint](): {
					reflect.TypeFor[uint8](): wrap(wrongTypes),
				},
			},
		}

		// --- When ---
		have := RegisterConverter(reg, cnv)

		// --- Then ---
		assert.Nil(t, have)
		from, to := reflect.TypeFor[uint](), reflect.TypeFor[uint8]()
		assert.Same(t, cnv, reg.m[from][to].cnv)
	})
}

func Test_LookupConverter(t *testing.T) {
	t.Run("registered", func(t *testing.T) {
		// --- Given ---
		cnv := func(from uint) (uint8, error) { return uint8(from), nil }
		reg := NewRegistry()
		reg.register(wrap(cnv))

		// --- When ---
		have := LookupConverter[uint, uint8](reg)

		// --- Then ---
		assert.Same(t, cnv, have)
	})

	t.Run("not registered", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := LookupConverter[uint, uint8](reg)

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("it should never happen", func(t *testing.T) {
		// --- Given ---
		wrongTypes := func(from float32) (int, error) { return 0, nil }
		reg := &Registry{
			m: map[reflect.Type]map[reflect.Type]*wrapper{
				reflect.TypeFor[uint](): {
					reflect.TypeFor[uint8](): wrap(wrongTypes),
				},
			},
		}

		// --- When ---
		have := LookupConverter[uint, uint8](reg)

		// --- Then ---
		assert.Nil(t, have)
	})
}
