// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToTime_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src    string
		format string
		dst    time.Time
		err    error
		msg    string
	}{
		{
			"RFC3339",
			"2000-01-02T03:04:05Z",
			time.RFC3339,
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			nil,
			"",
		},
		{
			"kitchen time",
			"3:42PM",
			time.Kitchen,
			time.Date(0000, 1, 1, 15, 42, 0, 0, time.UTC),
			nil,
			"",
		},
		{
			"error - empty string",
			"",
			time.RFC3339,
			time.Date(0000, 1, 1, 15, 42, 0, 0, time.UTC),
			ErrInvValue,
			"invalid value: from string to time.Time",
		},
		{
			"error - not matching format",
			"2000-01-02T03:04:05Z",
			time.Kitchen,
			time.Date(0000, 1, 1, 15, 42, 0, 0, time.UTC),
			ErrInvValue,
			"invalid value: from string to time.Time",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToTime(tc.format)(tc.src)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Exact(t, tc.dst, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Zero(t, have)
		})
	}
}
