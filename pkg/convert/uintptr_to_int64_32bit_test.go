// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build 386 || arm || mips || mipsle || wasm

package convert

import (
	"math"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_UintptrToInt64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		value uintptr
		want  int64
		err   error
		msg   string
	}{
		{"min", 0, 0, nil, ""},
		{"max", math.MaxUint, math.MaxUint, nil, ""},
	}

	for _, tc := range tt {
		t.Run("UintptrToInt64 "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := UintptrToInt64(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, have)
				assert.Equal(t, tc.value, uintptr(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, int64(0), have)
		})

		t.Run("UintptrToDuration "+tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := UintptrToDuration(tc.value)

			// --- Then ---
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, time.Duration(tc.want), have)
				assert.Equal(t, tc.value, uintptr(have))
				return
			}

			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, time.Duration(0), have)
		})
	}
}
