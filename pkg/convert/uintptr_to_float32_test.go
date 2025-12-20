// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_UintptrToFloat32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uintptr
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
			"uintptr value out of safe range for float32",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := UintptrToFloat32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uintptr(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float32(0), have)
		})
	}
}
