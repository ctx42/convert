// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/convert/pkg/xcast"
)

func Test_init(t *testing.T) {
	// --- Given ---
	types := xcast.SupportedTypes()

	// --- When ---
	for _, from := range types {
		for _, to := range types {
			assert.NotNil(t, registry.lookup(from, to), "%s -> %s", from, to)
		}
	}
}
