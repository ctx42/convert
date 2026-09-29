// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToDuration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := StringToDuration("4h2s")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 4*time.Hour+2*time.Second, have)
	})

	t.Run("error - parse error cause", func(t *testing.T) {
		// --- When ---
		have, err := StringToDuration("abc")

		// --- Then ---
		wMsg := "invalid value: from string to time.Duration"
		assert.ErrorEqual(t, wMsg, err)
		assert.ErrorIs(t, ErrInvValue, err)
		var cnvErr Error
		assert.ErrorAs(t, &cnvErr, err)
		assert.ErrorEqual(t, "time: invalid duration \"abc\"", cnvErr.Cause)
		assert.Equal(t, time.Duration(0), have)
	})
}

func Test_StringToDuration_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		src string
		err error
		msg string
	}{
		{
			"error - empty string",
			"",
			ErrInvValue,
			"invalid value: from string to time.Duration",
		},
		{
			"error - not matching format",
			"abc",
			ErrInvValue,
			"invalid value: from string to time.Duration",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := StringToDuration(tc.src)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, time.Duration(0), have)
		})
	}
}
