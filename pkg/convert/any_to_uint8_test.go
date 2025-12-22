// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"net/http"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_AnyToUint8(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  uint8
		err   error
		msg   string
	}{
		{"success from float64", 42.0, 42, nil, ""},
		{"success from int", 42, 42, nil, ""},
		{
			"error - underflow",
			-1,
			0,
			ErrInvRange,
			"int value out of range for uint8",
		},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			ErrUnkConv,
			"conversion undefined: from http.Cookie to uint8",
		},
	}

	for _, tc := range tt {
		t.Run("AnyToUint8 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToUint8(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint8(0), have)
		})

		t.Run("AnyToByte "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToByte(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, byte(0), have)
		})
	}
}
