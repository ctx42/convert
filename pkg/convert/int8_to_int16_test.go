// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Int8ToInt16_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value int8
		want  int16
		err   error
		msg   string
	}{
		{"min", math.MinInt8, math.MinInt8, nil, ""},
		{"max", math.MaxInt8, math.MaxInt8, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Int8ToInt16(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, int8(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int16(0), have)
		})
	}
}
