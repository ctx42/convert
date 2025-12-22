// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"net/http"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_AnyToUint32(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  uint32
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
			"int value out of range for uint32",
		},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			ErrUnkConv,
			"conversion undefined: from http.Cookie to uint32",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToUint32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uint32(0), have)
		})
	}
}
