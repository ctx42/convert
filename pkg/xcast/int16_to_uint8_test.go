// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int16ToUint8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int16
		want  uint8
		err   error
		msg   string
	}{
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"int16 value out of range for uint8",
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
			"int16 value out of range for uint8",
		},
	}

	for _, tc := range tt {
		t.Run("Int16ToUint8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int16ToUint8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int16(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})

		t.Run("Int16ToByte "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int16ToByte(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int16(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})
	}
}
