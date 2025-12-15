// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_StringToString(t *testing.T) {
	// --- When ---
	have, err := StringToString("abc")

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, "abc", have)
}
