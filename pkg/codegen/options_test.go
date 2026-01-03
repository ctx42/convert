// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewOptions(t *testing.T) {
	// --- Given ---
	opt0 := func(ops *Options) { ops.copyright += "a" }
	opt1 := func(ops *Options) { ops.copyright += "b" }
	opt2 := func(ops *Options) { ops.copyright += "c" }

	// --- When ---
	have := NewOptions(opt0, opt1, opt2)

	// --- Then ---
	assert.Equal(t, "abc", have.copyright)
}

func Test_WithCopyright(t *testing.T) {
	// --- Given ---
	ops := &Options{}

	// --- When ---
	WithCopyright("copyright")(ops)

	// --- Then ---
	assert.Equal(t, "copyright", ops.copyright)
}

func Test_WithGeneratedBy(t *testing.T) {
	// --- Given ---
	ops := &Options{}

	// --- When ---
	WithGeneratedBy("abc")(ops)

	// --- Then ---
	assert.Equal(t, "abc", ops.generatedBy)
}

func Test_WithVerbose(t *testing.T) {
	// --- Given ---
	ops := &Options{}

	// --- When ---
	WithVerbose(ops)

	// --- Then ---
	assert.True(t, ops.verbose)
}
