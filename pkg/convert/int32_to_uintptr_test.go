// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int32ToUintptr_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int32
		want  uintptr
		err   error
		msg   string
	}{
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"int32 value out of range for uintptr",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt32, math.MaxInt32, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Int32ToUintptr "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int32ToUintptr(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uintptr(0), have)
		})

		t.Run("RuneToUintptr "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := RuneToUintptr(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uintptr(0), have)
		})
	}
}
