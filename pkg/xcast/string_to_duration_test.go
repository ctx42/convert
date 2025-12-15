// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToDuration_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value string
		want  time.Duration
		err   error
		msg   string
	}{
		{
			"success",
			"4h2s",
			4*time.Hour + 2*time.Second,
			nil,
			"",
		},
		{
			"error - empty string",
			"",
			0,
			ErrInvValue,
			"cannot convert (parse) an empty string to time.Duration: invalid value",
		},
		{
			"error - not matching format",
			"abc",
			0,
			ErrInvValue,
			`cannot convert (parse) "abc" as time.Duration: invalid value`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToDuration(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, time.Duration(0), have)
		})
	}
}
