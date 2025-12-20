// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float64ToFloat64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float64
		want  float64
		err   error
		msg   string
	}{
		{"min", -math.MaxFloat64, -math.MaxFloat64, nil, ""},
		{"max", math.MaxFloat64, math.MaxFloat64, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float64ToFloat64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, float64(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float64(0), have)
		})
	}
}
