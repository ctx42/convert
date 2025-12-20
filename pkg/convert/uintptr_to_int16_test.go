// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_UintptrToInt16_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uintptr
		want  int16
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt16, math.MaxInt16, nil, ""},
		{"overflow", math.MaxInt16 + 1, 0, ErrInvRange, "uintptr value out of range for int16"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := UintptrToInt16(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uintptr(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int16(0), have)
		})
	}
}
