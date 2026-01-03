// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToDuration_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src string
		dst time.Duration
		err error
		msg string
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
			"invalid value: from string to time.Duration",
		},
		{
			"error - not matching format",
			"abc",
			0,
			ErrInvValue,
			"invalid value: from string to time.Duration",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToDuration(tc.src)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.dst, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, time.Duration(0), have)
		})
	}
}
