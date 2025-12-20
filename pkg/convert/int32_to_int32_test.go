// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int32ToInt32_tabular(t *testing.T) {

	tt := []struct {
		testN string

		value int32
		want  int32
		err   error
		msg   string
	}{
		{"min", math.MaxInt32, math.MaxInt32, nil, ""},
		{"max", math.MaxInt32, math.MaxInt32, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Int32ToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int32ToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, have, tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("RuneToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := RuneToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, have, tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("Int32ToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int32ToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, have, tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("RuneToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := RuneToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, have, tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})
	}
}
