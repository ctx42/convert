// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint8ToInt32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint8
		want  int32
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"max", math.MaxUint8, math.MaxUint8, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Uint8ToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint8ToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint8(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("Uint8ToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint8ToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint8(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, rune(0), have)
		})

		t.Run("ByteToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := ByteToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, byte(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("ByteToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := ByteToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, byte(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, rune(0), have)
		})
	}
}
