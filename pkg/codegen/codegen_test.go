// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewAction(t *testing.T) {
	t.Run("with value", func(t *testing.T) {
		// --- Given ---
		val := NewValue("value")

		// --- When ---
		have := NewAction(CheckIntSafeToFloatMin, val)

		// --- Then ---
		assert.Equal(t, have.name, CheckIntSafeToFloatMin)
		assert.Same(t, have.value, val)
	})

	t.Run("without value", func(t *testing.T) {
		// --- When ---
		have := NewAction(CheckIntSafeToFloatMin, nil)

		// --- Then ---
		assert.Equal(t, have.name, CheckIntSafeToFloatMin)
		assert.Nil(t, have.value)
	})
}

func Test_Action_Name(t *testing.T) {
	// --- Given ---
	act := &Action{name: CheckIntSafeToFloatMin}

	// --- When ---
	have := act.Name()

	// --- Then ---
	assert.Equal(t, have, CheckIntSafeToFloatMin)
}

func Test_Action_Imports(t *testing.T) {
	t.Run("without value", func(t *testing.T) {
		// --- Given ---
		act := NewAction(CheckIntSafeToFloatMin, nil)

		// --- When ---
		have := act.Imports()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("with value not requiring renderImports", func(t *testing.T) {
		// --- Given ---
		val := NewValue("value")
		act := NewAction(CheckIntSafeToFloatMin, val)

		// --- When ---
		have := act.Imports()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("with value requiring renderImports", func(t *testing.T) {
		// --- Given ---
		val := NewValue("pkg", "value")
		act := NewAction(CheckIntSafeToFloatMin, val)

		// --- When ---
		have := act.Imports()

		// --- Then ---
		assert.Equal(t, []string{"pkg"}, have)
	})
}

func Test_Action_Code(t *testing.T) {
	t.Run("without value", func(t *testing.T) {
		// --- Given ---
		act := NewAction(CheckIntSafeToFloatMin, nil)

		// --- When ---
		have := act.Code()

		// --- Then ---
		assert.Empty(t, have)
	})

	t.Run("with value", func(t *testing.T) {
		// --- Given ---
		val := NewValue("pkg", "value")
		act := NewAction(CheckIntSafeToFloatMin, val)

		// --- When ---
		have := act.Code()

		// --- Then ---
		assert.Equal(t, "pkg.value", have)
	})
}
