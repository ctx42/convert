// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_IntToInt32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int
		want  int32
		err   error
		msg   string
	}{
		{
			"underflow",
			math.MinInt32 - 1,
			0,
			ErrInvRange,
			"int value out of range for int32",
		},
		{"min", math.MinInt32, math.MinInt32, nil, ""},
		{"negative", -1, -1, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt32, math.MaxInt32, nil, ""},
		{
			"overflow",
			math.MaxInt32 + 1,
			0,
			ErrInvRange,
			"int value out of range for int32",
		},
	}

	for _, tc := range tt {
		t.Run("IntToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := IntToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("IntToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := IntToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})
	}
}
