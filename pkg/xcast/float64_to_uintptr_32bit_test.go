// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float64ToUintptr_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float64
		want  uintptr
		err   error
		msg   string
	}{
		{
			"fraction",
			4.2,
			0,
			ErrInvValue,
			"uintptr requires non-fractional float64: invalid value",
		},
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"float64 value out of range for uintptr",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxUint, math.MaxUint, nil, ""},
		{
			"overflow",
			math.MaxUint + 1,
			0,
			ErrInvSafeRange,
			"float64 value out of safe range for uintptr",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToUintptr(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, float64(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uintptr(0), have)
		})
	}
}
