// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint8ToInt8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint8
		want  int8
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", uint8(math.MaxInt8), int8(math.MaxInt8), nil, ""},
		{
			"overflow",
			uint8(math.MaxInt8) + 1,
			int8(0),
			ErrInvRange,
			"uint8 value out of range for int8",
		},
	}

	for _, tc := range tt {
		t.Run("Uint8ToInt8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint8ToInt8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint8(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int8(0), have)
		})

		t.Run("ByteToInt8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := ByteToInt8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, byte(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int8(0), have)
		})
	}
}
