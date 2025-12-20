// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_IntToInt8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int
		want  int8
		err   error
		msg   string
	}{
		{
			"underflow",
			math.MinInt8 - 1,
			0,
			ErrInvRange,
			"int value out of range for int8",
		},
		{"min", math.MinInt8, math.MinInt8, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt8, math.MaxInt8, nil, ""},
		{
			"overflow",
			math.MaxInt8 + 1,
			0,
			ErrInvRange,
			"int value out of range for int8",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := IntToInt8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int8(0), have)
		})
	}
}
