// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_IntToDuration_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int
		want  time.Duration
		err   error
		msg   string
	}{
		{"min", math.MinInt, math.MinInt, nil, ""},
		{"max", math.MaxInt, math.MaxInt, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := IntToDuration(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})
	}
}
