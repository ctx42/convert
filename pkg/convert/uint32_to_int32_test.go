// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint32ToInt32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint32
		want  int32
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt32, math.MaxInt32, nil, ""},
		{
			"overflow",
			math.MaxInt32 + 1,
			0,
			ErrInvRange,
			"uint32 value out of range for int32",
		},
	}

	for _, tc := range tt {
		t.Run("Uint32ToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint32ToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("Uint32ToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint32ToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})
	}
}
