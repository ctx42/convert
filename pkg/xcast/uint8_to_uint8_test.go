// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint8ToUint8_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint8
		want  uint8
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"max", math.MaxUint8, math.MaxUint8, nil, ""},
	}

	for _, tc := range tt {
		t.Run("Uint8ToUint8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint8ToUint8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, uint8(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})

		t.Run("ByteToByte "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := ByteToByte(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, byte(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})
	}
}
