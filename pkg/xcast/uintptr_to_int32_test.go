// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_UintptrToInt32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uintptr
		want  int32
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", uintptr(math.MaxInt32), math.MaxInt32, nil, ""},
		{
			"overflow",
			uintptr(math.MaxInt32) + 1,
			0,
			ErrInvRange,
			"uintptr value out of range for int32",
		},
	}

	for _, tc := range tt {
		t.Run("UintptrToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := UintptrToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uintptr(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("UintptrToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := UintptrToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uintptr(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})
	}
}
