// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int32ToInt16_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int32
		want  int16
		err   error
		msg   string
	}{
		{
			"underflow",
			math.MinInt16 - 1,
			0,
			ErrInvRange,
			"int32 value out of range for int16",
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
			"int32 value out of range for int16",
		},
	}

	for _, tc := range tt {
		t.Run("Int32ToInt16 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int32ToInt16(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int16(0), have)
		})

		t.Run("RuneToInt16 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := RuneToInt16(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int16(0), have)
		})
	}
}
