// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"net/http"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_AnyToInt32(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  int32
		err   error
		msg   string
	}{
		{"success from float64", 42.0, 42, nil, ""},
		{"success from uint32", uint32(42), 42, nil, ""},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			ErrUnkConv,
			"conversion undefined: from http.Cookie to int32",
		},
		{
			"error - overflow",
			uint64(math.MaxUint32 + 1),
			0,
			ErrInvRange,
			"uint64 value out of range for int32",
		},
	}

	for _, tc := range tt {
		t.Run("AnyToInt32 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToInt32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int32(0), have)
		})

		t.Run("AnyToRune "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToRune(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, rune(0), have)
		})
	}
}
