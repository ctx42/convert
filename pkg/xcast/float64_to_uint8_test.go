// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float64ToUint8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float64
		want  uint8
		err   error
		msg   string
	}{
		{"fraction", 42.5, 0, ErrInvValue, "uint8 requires non-fractional float64: invalid value"},
		{"negative", -1, 0, ErrInvRange, "float64 value out of range for uint8"},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxUint8, math.MaxUint8, nil, ""},
		{"overflow", math.MaxUint8 + 1, 0, ErrInvRange, "float64 value out of range for uint8"},
	}

	for _, tc := range tt {
		t.Run("Float64ToUint8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToUint8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, float64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})

		t.Run("Float64ToByte "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToByte(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, float64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})
	}
}
