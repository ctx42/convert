// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float64ToInt32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float64
		want  int32
		err   error
		msg   string
	}{
		{
			"fraction",
			4.2,
			0,
			ErrInvValue,
			"int32 requires non-fractional float64: invalid value",
		},
		{
			"underflow",
			Float64SafeIntMin - 1,
			0,
			ErrInvRange,
			"float64 value out of range for int32",
		},
		{"min", math.MinInt32, math.MinInt32, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt32, math.MaxInt32, nil, ""},
		{
			"overflow",
			Float64SafeIntMax,
			0,
			ErrInvRange,
			"float64 value out of range for int32",
		},
	}

	for _, tc := range tt {
		t.Run("Float64ToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, float64(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("Float64ToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, float64(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})
	}
}
