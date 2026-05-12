// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"reflect"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_wrap(t *testing.T) {
	// --- Given ---
	cnv := func(src uint) (uint8, error) { return uint8(src + 2), nil }

	// --- When ---
	have := wrap(cnv)

	// --- Then ---
	assert.Equal(t, reflect.TypeFor[uint](), have.src)
	assert.Equal(t, reflect.TypeFor[uint8](), have.dst)
	assert.Same(t, cnv, have.cnv)

	val, err := have.cst(uint(42))
	assert.NoError(t, err)
	assert.Equal(t, uint8(44), val)
}
