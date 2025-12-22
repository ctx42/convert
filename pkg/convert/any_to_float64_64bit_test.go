// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"net/http"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_AnyToFloat64(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  float64
		err   error
		msg   string
	}{
		{"success from float32", float32(42.0), 42.0, nil, ""},
		{"success from int", 42, 42.0, nil, ""},
		{
			"error - safe range",
			Float64SafeIntMin - 1,
			0.0,
			ErrInvSafeRange,
			"int value out of safe range for float64",
		},
		{
			"error - undefined conversion",
			http.Cookie{},
			0.0,
			ErrUnkConv,
			"conversion undefined: from http.Cookie to float64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToFloat64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float64(0), have)
		})
	}
}
