// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint8ToInt_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint8
		want  int
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"max", math.MaxUint8, math.MaxUint8, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Uint8ToInt "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint8ToInt(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uint8(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, 0, have)
		})

		t.Run("ByteToInt "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := ByteToInt(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, byte(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, 0, have)
		})
	}
}
