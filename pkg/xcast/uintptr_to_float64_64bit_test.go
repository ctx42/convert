// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package xcast

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_UintptrToFloat64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uintptr
		want  float64
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", Float64SafeIntMax, Float64SafeIntMax, nil, ""},
		{
			"overflow",
			Float64SafeIntMax + 1,
			Float64SafeIntMax + 1,
			ErrInvSafeRange,
			"uintptr value out of safe range for float64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := UintptrToFloat64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uintptr(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float64(0), have)
		})
	}
}
