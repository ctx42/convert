// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_DurationToDuration_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value time.Duration
		want  time.Duration
		err   error
		msg   string
	}{
		{"min", math.MinInt64, math.MinInt64, nil, ""},
		{"max", math.MaxInt64, math.MaxInt64, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DurationToDuration(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, have)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, time.Duration(0), have)
		})
	}
}
