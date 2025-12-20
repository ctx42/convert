// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int32ToInt64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int32
		want  int64
		err   error
		msg   string
	}{
		{"min", math.MinInt32, math.MinInt32, nil, ""},
		{"max", math.MaxInt32, math.MaxInt32, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Int32ToInt64 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int32ToInt64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})

		t.Run("RuneToInt64 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := RuneToInt64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int32(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})
	}
}
