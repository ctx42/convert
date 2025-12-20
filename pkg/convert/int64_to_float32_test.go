// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int64ToFloat32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int64
		want  float32
		err   error
		msg   string
	}{
		{
			"underflow",
			Float32SafeIntMin - 1,
			0,
			ErrInvSafeRange,
			"int64 value out of safe range for float32",
		},
		{"min", Float32SafeIntMin, Float32SafeIntMin, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 42, 42, nil, ""},
		{"max", Float32SafeIntMax, Float32SafeIntMax, nil, ""},
		{
			"overflow",
			Float32SafeIntMax + 1,
			0,
			ErrInvSafeRange,
			"int64 value out of safe range for float32",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int64ToFloat32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float32(0), have)
		})
	}
}
