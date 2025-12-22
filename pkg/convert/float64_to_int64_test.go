// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float64ToInt64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float64
		want  int64
		err   error
		msg   string
	}{
		{
			"fraction",
			4.2,
			0,
			ErrInvValue,
			"int64 requires non-fractional float64: invalid value",
		},
		{
			"underflow",
			Float64SafeIntMin - 1,
			0,
			ErrInvSafeRange,
			"float64 value out of safe range for int64",
		},
		{"min", Float64SafeIntMin, Float64SafeIntMin, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", Float64SafeIntMax, Float64SafeIntMax, nil, ""},
		{
			"overflow",
			Float64SafeIntMax + 1,
			0,
			ErrInvSafeRange,
			"float64 value out of safe range for int64",
		},
	}

	for _, tc := range tt {
		t.Run("Float64ToInt64 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToInt64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, float64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})

		t.Run("Float64ToDuration "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToDuration(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, time.Duration(tc.want), have)
				assert.Equal(t, tc.value, float64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, time.Duration(0), have)
		})
	}
}
