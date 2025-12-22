// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package convert

import (
	"math"
	"net/http"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_AnyToInt64(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  int64
		err   error
		msg   string
	}{
		{"success from float64", 42.0, 42, nil, ""},
		{"success from uint8", uint8(42), 42, nil, ""},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			ErrUnkConv,
			"conversion undefined: from http.Cookie to int64",
		},
		{
			"error - overflow",
			uint64(math.MaxUint64),
			0,
			ErrInvRange,
			"uint64 value out of range for int64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToInt64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})
	}
}
