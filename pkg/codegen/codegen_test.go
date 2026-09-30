// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
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
		assert.Equal(t, CheckIntSafeToFloatMin, have.name)
		assert.Same(t, val, have.value)
	})

	t.Run("without value", func(t *testing.T) {
		// --- When ---
		have := NewAction(CheckIntSafeToFloatMin, nil)

		// --- Then ---
		assert.Equal(t, CheckIntSafeToFloatMin, have.name)
		assert.Nil(t, have.value)
	})
}

func Test_Action_Name(t *testing.T) {
	// --- Given ---
	act := &Action{name: CheckIntSafeToFloatMin}

	// --- When ---
	have := act.Name()

	// --- Then ---
	assert.Equal(t, CheckIntSafeToFloatMin, have)
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

	t.Run("with value not requiring imports", func(t *testing.T) {
		// --- Given ---
		val := NewValue("value")
		act := NewAction(CheckIntSafeToFloatMin, val)

		// --- When ---
		have := act.Imports()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("with value requiring imports", func(t *testing.T) {
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

func Test_Action_equal(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		// --- Given ---
		act := NewAction(CheckOverflows, NewValue("math", "MaxInt32"))
		other := NewAction(CheckOverflows, NewValue("math", "MaxInt32"))

		// --- When ---
		have := act.equal(other)

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("different name", func(t *testing.T) {
		// --- Given ---
		act := NewAction(CheckOverflows, NewValue("math", "MaxInt32"))
		other := NewAction(CheckUnderflows, NewValue("math", "MaxInt32"))

		// --- When ---
		have := act.equal(other)

		// --- Then ---
		assert.False(t, have)
	})

	t.Run("different value", func(t *testing.T) {
		// --- Given ---
		act := NewAction(CheckOverflows, NewValue("math", "MaxInt32"))
		other := NewAction(CheckOverflows, NewValue("math", "MaxInt64"))

		// --- When ---
		have := act.equal(other)

		// --- Then ---
		assert.False(t, have)
	})
}

func Test_mergeActions(t *testing.T) {
	t.Run("no lists", func(t *testing.T) {
		// --- Given ---
		target := NumericType[int32]()

		// --- When ---
		have := mergeActions(target)

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("deduplicates and orders", func(t *testing.T) {
		// --- Given ---
		target := NumericType[int32]()
		a32 := []Action{
			NewAction(CheckUnderflows, NewValue("math", "MinInt32")),
			NewAction(CastDirectly, nil),
		}
		a64 := []Action{
			NewAction(CheckIntSafeToFloatMin, NewValue("Float64SafeIntMin")),
			NewAction(CheckIsNonNegative, nil),
			NewAction(CastDirectly, nil),
		}

		// --- When ---
		have := mergeActions(target, a32, a64)

		// --- Then ---
		want := []Action{
			NewAction(CheckIsNonNegative, nil),
			NewAction(CheckUnderflows, NewValue("math", "MinInt32")),
			NewAction(CheckIntSafeToFloatMin, NewValue("Float64SafeIntMin")),
			NewAction(CastDirectly, nil),
		}
		assert.Equal(t, want, have)
	})

	t.Run("platform target range limits", func(t *testing.T) {
		// --- Given ---
		target := NumericType[int]()
		a32 := []Action{
			NewAction(CheckUnderflows, NewValue("math", "MinInt32")),
			NewAction(CheckOverflows, NewValue("math", "MaxInt32")),
			NewAction(CastDirectly, nil),
		}
		a64 := []Action{
			NewAction(CheckOverflows, NewValue("math", "MaxInt64")),
			NewAction(CastDirectly, nil),
		}

		// --- When ---
		have := mergeActions(target, a32, a64)

		// --- Then ---
		want := []Action{
			NewAction(CheckUnderflows, NewValue("math", "MinInt")),
			NewAction(CheckOverflows, NewValue("math", "MaxInt")),
			NewAction(CastDirectly, nil),
		}
		assert.Equal(t, want, have)
	})
}
