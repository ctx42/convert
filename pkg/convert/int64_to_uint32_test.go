// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int64ToUint32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int64
		want  uint32
		err   error
		msg   string
	}{
		{
			"underflow",
			-1,
			0,
			ErrInvRange,
			"int64 value out of range for uint32",
		},
		{"min", 0, 0, nil, ""},
		{"zero", 0, 0, nil, ""},
		{"positive", 1, 1, nil, ""},
		{"max", math.MaxUint32, math.MaxUint32, nil, ""},
		{
			"overflow",
			math.MaxUint32 + 1,
			0,
			ErrInvRange,
			"int64 value out of range for uint32",
		},
	}

	for _, tc := range tt {
		t.Run("Int64ToUint32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int64ToUint32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint32(0), have)
		})

		t.Run("DurationToUint32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DurationToUint32(time.Duration(tc.value))

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint32(0), have)
		})
	}
}
