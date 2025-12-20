// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"math"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Float32ToFloat32_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value float32
		want  float32
		err   error
		msg   string
	}{
		{"min", -math.MaxFloat32, -math.MaxFloat32, nil, ""},
		{"max", math.MaxFloat32, math.MaxFloat32, nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := Float32ToFloat32(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, have, tc.value)
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, float32(0), have)
		})
	}
}
