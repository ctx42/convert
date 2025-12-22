// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

import (
	"net/http"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_AnyToUintptr(t *testing.T) {
	tt := []struct {
		testN string

		value any
		want  uintptr
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
			"int value out of range for uintptr",
		},
		{
			"error - undefined conversion",
			http.Cookie{},
			0,
			ErrUnkConv,
			"conversion undefined: from http.Cookie to uintptr",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := AnyToUintptr(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, uintptr(0), have)
		})
	}
}
