// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int16ToInt32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int16
		want  int32
		err   error
		msg   string
	}{
		{"min", math.MinInt16, math.MinInt16, nil, ""},
		{"max", math.MaxInt16, math.MaxInt16, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Int16ToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int16ToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, int16(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("Int16ToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int16ToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, int16(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})
	}
}
