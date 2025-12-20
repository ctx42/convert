// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint32ToUint8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint32
		want  uint8
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxUint8, math.MaxUint8, nil, ""},
		{
			"overflow",
			math.MaxUint8 + 1,
			0,
			ErrInvRange,
			"uint32 value out of range for uint8",
		},
	}

	for _, tc := range tt {
		t.Run("Uint32ToUint8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint32ToUint8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})

		t.Run("Uint32ToByte "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint32ToByte(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})
	}
}
