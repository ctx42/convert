// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float32ToInt8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float32
		want  int8
		err   error
		msg   string
	}{
		{"fraction", 4.2, 0, ErrInvValue, "int8 requires non-fractional float32: invalid value"},
		{"underflow", math.MinInt8 - 1, 0, ErrInvRange, "float32 value out of range for int8"},
		{"min", math.MinInt8, math.MinInt8, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt8, math.MaxInt8, nil, ""},
		{"overflow", math.MaxInt8 + 1, 0, ErrInvRange, "float32 value out of range for int8"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float32ToInt8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, float32(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int8(0), have)
		})
	}
}
