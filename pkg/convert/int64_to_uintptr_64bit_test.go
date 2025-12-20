// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int64ToUintptr_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int64
		want  uintptr
		err   error
		msg   string
	}{
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"int64 value out of range for uintptr",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt64, math.MaxInt64, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int64ToUintptr(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uintptr(0), have)
		})
	}
}
