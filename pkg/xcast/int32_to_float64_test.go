// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int32ToFloat64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int32
		want  float64
		err   error
		msg   string
	}{
		{"min", math.MinInt32, math.MinInt32, nil, ""},
		{"max", math.MaxInt32, math.MaxInt32, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Int32ToFloat64 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int32ToFloat64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, int32(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float64(0), have)
		})

		t.Run("RuneToFloat64 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := RuneToFloat64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, int32(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float64(0), have)
		})
	}
}
