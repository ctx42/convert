// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int8ToUint16_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int8
		want  uint16
		err   error
		msg   string
	}{
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"int8 value out of range for uint16",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxInt8, math.MaxInt8, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int8ToUint16(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int8(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint16(0), have)
		})
	}
}
