// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package convert

import (
	"math"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int64ToInt_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int64
		want  int
		err   error
		msg   string
	}{
		{
			"underflow",
			math.MinInt - 1,
			0,
			ErrInvRange,
			"int64 value out of range for int",
		},
		{"min", math.MinInt, math.MinInt, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt, math.MaxInt, nil, ""},
		{
			"overflow",
			math.MaxInt + 1,
			0,
			ErrInvRange,
			"int64 value out of range for int",
		},
	}

	for _, tc := range tt {
		t.Run("Int64ToInt "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int64ToInt(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int(0), have)
		})

		t.Run("DurationToInt "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DurationToInt(time.Duration(tc.value))

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int(0), have)
		})
	}
}
