// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package xconv

import (
	"math"
	"net/http"
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/convert/pkg/xcast"
)

func Test_AnyToInt(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  int
		err   error
		msg   string
	}{
		{"success from float64", 42.0, 42, nil, ""},
		{"success from uint8", uint8(42), 42, nil, ""},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			xcast.ErrUndConv,
			"conversion undefined: from http.Cookie to int",
		},
		{
			"error - overflow",
			uint64(math.MaxUint64),
			0,
			xcast.ErrInvRange,
			"uint64 value out of range for int",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToInt(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, 0, have)
		})
	}
}

func Test_AnyToInt64(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  int64
		err   error
		msg   string
	}{
		{"success from float64", 42.0, 42, nil, ""},
		{"success from uint8", uint8(42), 42, nil, ""},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			xcast.ErrUndConv,
			"conversion undefined: from http.Cookie to int64",
		},
		{
			"error - overflow",
			uint64(math.MaxUint64),
			0,
			xcast.ErrInvRange,
			"uint64 value out of range for int64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToInt64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})
	}
}

func Test_AnyToUint(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  uint
		err   error
		msg   string
	}{
		{"success from float64", 42.0, 42, nil, ""},
		{"success from int", 42, 42, nil, ""},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			xcast.ErrUndConv,
			"conversion undefined: from http.Cookie to uint",
		},
		{
			"error - underflow",
			-1,
			0,
			xcast.ErrInvRange,
			"int value out of range for uint",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToUint(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint(0), have)
		})
	}
}

func Test_AnyToUint64(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  uint64
		err   error
		msg   string
	}{
		{"success from float64", 42.0, 42, nil, ""},
		{"success from int", 42, 42, nil, ""},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			xcast.ErrUndConv,
			"conversion undefined: from http.Cookie to uint64",
		},
		{
			"error - underflow",
			-1,
			0,
			xcast.ErrInvRange,
			"int value out of range for uint64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToUint64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint64(0), have)
		})
	}
}

func Test_AnyToFloat64(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  float64
		err   error
		msg   string
	}{
		{"success from float32", float32(42.0), 42.0, nil, ""},
		{"success from int", 42, 42.0, nil, ""},
		{
			"error - undefined conversion",
			http.Cookie{},
			0.0,
			xcast.ErrUndConv,
			"conversion undefined: from http.Cookie to float64",
		},
		{
			"error - safe range",
			int64(xcast.Float64SafeIntMax + 1),
			0.0,
			xcast.ErrInvSafeRange,
			"int64 value out of safe range for float64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToFloat64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float64(0), have)
		})
	}
}
