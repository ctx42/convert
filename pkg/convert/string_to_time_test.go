// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToTime(t *testing.T) {
	t.Run("error - parse error cause", func(t *testing.T) {
		// --- Given ---
		cnv := StringToTime(time.Kitchen)

		// --- When ---
		have, err := cnv("abc")

		// --- Then ---
		assert.ErrorEqual(t, "invalid value: from string to time.Time", err)
		assert.ErrorIs(t, ErrInvValue, err)
		var pErr *time.ParseError
		assert.ErrorAs(t, &pErr, err)
		assert.Zero(t, have)
	})
}

func Test_StringToTime_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src    string
		format string
		dst    time.Time
	}{
		{
			"RFC3339",
			"2000-01-02T03:04:05Z",
			time.RFC3339,
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			"kitchen time",
			"3:42PM",
			time.Kitchen,
			time.Date(0000, 1, 1, 15, 42, 0, 0, time.UTC),
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			cnv := StringToTime(tc.format)

			// --- When ---
			have, err := cnv(tc.src)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exact(t, tc.dst, have)
		})
	}
}

func Test_StringToTime_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src    string
		format string
		err    error
		msg    string
	}{
		{
			"error - empty string",
			"",
			time.RFC3339,
			ErrInvValue,
			"invalid value: from string to time.Time",
		},
		{
			"error - not matching format",
			"2000-01-02T03:04:05Z",
			time.Kitchen,
			ErrInvValue,
			"invalid value: from string to time.Time",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			cnv := StringToTime(tc.format)

			// --- When ---
			have, err := cnv(tc.src)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Zero(t, have)
		})
	}
}
