// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint32ToFloat32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint32
		want  float32
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", Float32SafeIntMax, Float32SafeIntMax, nil, ""},
		{
			"overflow",
			Float32SafeIntMax + 1,
			0,
			ErrInvSafeRange,
			"uint32 value out of safe range for float32",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint32ToFloat32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, uint32(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float32(0), have)
		})
	}
}
