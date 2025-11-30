// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv

import (
	"reflect"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_wrap(t *testing.T) {
	// --- Given ---
	cnv := func(from uint) (uint8, error) { return uint8(from), nil }

	// --- When ---
	have := wrap(cnv)

	// --- Then ---
	assert.Equal(t, reflect.TypeFor[uint](), have.from)
	assert.Equal(t, reflect.TypeFor[uint8](), have.to)
	assert.Same(t, cnv, have.conv)
}
