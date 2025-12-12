// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float32ToInt32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float32
		want  int32
		err   error
		msg   string
	}{
		{"fraction", 4.2, 0, ErrInvValue, "int32 requires non-fractional float32: invalid value"},
		{"underflow", Float32SafeIntMin - 1, 0, ErrInvSafeRange, "float32 value out of safe range for int32"},
		{"min", Float32SafeIntMin, Float32SafeIntMin, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", Float32SafeIntMax, Float32SafeIntMax, nil, ""},
		{"overflow", Float32SafeIntMax + 1, 0, ErrInvSafeRange, "float32 value out of safe range for int32"},
	}

	for _, tc := range tt {
		t.Run("Float32ToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float32ToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, float32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("Float32ToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float32ToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, float32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})
	}
}
