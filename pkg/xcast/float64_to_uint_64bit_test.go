// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package xcast

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float64ToUint_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float64
		want  uint
		err   error
		msg   string
	}{
		{
			"fraction",
			4.2,
			0,
			ErrInvValue,
			"uint requires non-fractional float64: invalid value",
		},
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"float64 value out of range for uint",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", Float64SafeIntMax, Float64SafeIntMax, nil, ""},
		{
			"overflow",
			Float64SafeIntMax + 1,
			0,
			ErrInvSafeRange,
			"float64 value out of safe range for uint",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToUint(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, float64(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint(0), have)
		})
	}
}
