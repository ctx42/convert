// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int64ToUint_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int64
		want  uint
		err   error
		msg   string
	}{
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"int64 value out of range for uint",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxUint, math.MaxUint, nil, ""},
		{
			"overflow",
			math.MaxUint + 1,
			0,
			ErrInvRange,
			"int64 value out of range for uint",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int64ToUint(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, int64(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint(0), have)
		})
	}
}
