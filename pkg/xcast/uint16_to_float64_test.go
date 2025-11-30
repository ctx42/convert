// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Uint16ToFloat64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uint16
		want  float64
		err   error
		msg   string
	}{
		{"min", math.MaxUint16, math.MaxUint16, nil, ""},
		{"max", math.MaxUint16, math.MaxUint16, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Uint16ToFloat64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, uint16(have), tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float64(0), have)
		})
	}
}
