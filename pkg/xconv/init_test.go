// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/convert/pkg/xcast"
)

func Test_init(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		// --- Given ---
		types := xcast.SupportedTypes()

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
		assert.Same(t, xcast.StringToString, have.cnv)
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

		cnv := have.cnv.(Converter[string, time.Time])
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
		assert.Same(t, xcast.StringToDuration, have.cnv)
	})
}
