// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float64ToUint16_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float64
		want  uint16
		err   error
		msg   string
	}{
		{
			"fraction",
			4.2,
			0,
			ErrInvValue,
			"uint16 requires non-fractional float64: invalid value",
		},
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"float64 value out of range for uint16",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxUint16, math.MaxUint16, nil, ""},
		{
			"overflow",
			math.MaxUint16 + 1,
			0,
			ErrInvRange,
			"float64 value out of range for uint16",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToUint16(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, float64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint16(0), have)
		})
	}
}
