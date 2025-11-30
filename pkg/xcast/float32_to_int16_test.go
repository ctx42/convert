// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float32ToInt16_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float32
		want  int16
		err   error
		msg   string
	}{
		{
			"fraction",
			4.2,
			0,
			ErrInvValue,
			"int16 requires non-fractional float32: invalid value",
		},
		{
			"underflow",
			math.MinInt16 - 1,
			0,
			ErrInvRange,
			"float32 value out of range for int16",
		},
		{"min", math.MinInt16, math.MinInt16, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt16, math.MaxInt16, nil, ""},
		{
			"overflow",
			math.MaxInt16 + 1,
			0,
			ErrInvRange,
			"float32 value out of range for int16",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float32ToInt16(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, float32(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int16(0), have)
		})
	}
}
