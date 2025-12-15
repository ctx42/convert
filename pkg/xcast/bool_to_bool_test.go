// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_BoolToBool(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		// --- When ---
		have, err := BoolToBool(true)

		// --- Then ---
		assert.NoError(t, err)
		assert.True(t, have)
	})

	t.Run("false", func(t *testing.T) {
		// --- When ---
		have, err := BoolToBool(false)

		// --- Then ---
		assert.NoError(t, err)
		assert.False(t, have)
	})
}
