// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint32ToInt_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint32
		want  int
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"max", math.MaxInt, math.MaxInt, nil, ""},
		{
			"overflows",
			math.MaxInt + 1,
			0,
			ErrInvRange,
			"uint32 value out of range for int",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint32ToInt(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int(0), have)
		})
	}
}
