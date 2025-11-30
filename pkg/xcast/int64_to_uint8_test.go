// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int64ToUint8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int64
		want  uint8
		err   error
		msg   string
	}{
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"int64 value out of range for uint8",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxUint8, math.MaxUint8, nil, ""},
		{
			"overflow",
			math.MaxUint8 + 1,
			0,
			ErrInvRange,
			"int64 value out of range for uint8",
		},
	}

	for _, tc := range tt {
		t.Run("Int64ToUint8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int64ToUint8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, int64(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})

		t.Run("Int64ToByte "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int64ToByte(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, int64(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})
	}
}
