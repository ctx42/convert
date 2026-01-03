// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewValue_tabular(t *testing.T) {
	tt := []struct {
		testN string

		arg    string
		args   []string
		wPref  string
		wPkg   string
		wValue string
	}{
		{"name", "ConstantName", nil, "", "", "ConstantName"},
		{
			"package and name",
			"math",
			[]string{"MinInt32"},
			"",
			"math",
			"MinInt32",
		},
		{
			"prefix and package and name",
			"-",
			[]string{"math", "MaxFloat32"},
			"-",
			"math",
			"MaxFloat32",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := NewValue(tc.arg, tc.args...)

			// --- Then ---
			assert.Equal(t, tc.wPref, have.prefix)
			assert.Equal(t, tc.wPkg, have.pkg)
			assert.Equal(t, tc.wValue, have.value)
		})
	}
}

func Test_NewValue(t *testing.T) {
	t.Run("panics when invalid number of args", func(t *testing.T) {
		// --- When ---
		msg := assert.PanicMsg(t, func() { NewValue("a", "b", "c", "d") })

		// --- Then ---
		assert.NotEmpty(t, *msg)
	})
}

func Test_Value_Name(t *testing.T) {
	t.Run("nil instance", func(t *testing.T) {
		// --- Given ---
		var val *Value

		// --- When ---
		have := val.Name(false)

		// --- Then ---
		assert.Empty(t, have)
	})

	t.Run("name", func(t *testing.T) {
		// --- Given ---
		val := NewValue("byte")

		// --- When ---
		have := val.Name(false)

		// --- Then ---
		assert.Equal(t, "byte", have)
	})

	t.Run("titled name", func(t *testing.T) {
		// --- Given ---
		val := NewValue("byte")

		// --- When ---
		have := val.Name(true)

		// --- Then ---
		assert.Equal(t, "Byte", have)
	})
}

func Test_Value_Imports(t *testing.T) {
	t.Run("nil instance", func(t *testing.T) {
		// --- Given ---
		var val *Value

		// --- When ---
		have := val.Imports()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("without import", func(t *testing.T) {
		// --- Given ---
		val := NewValue("byte")

		// --- When ---
		have := val.Imports()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("with import", func(t *testing.T) {
		// --- Given ---
		val := NewValue("time", "Duration")

		// --- When ---
		have := val.Imports()

		// --- Then ---
		assert.Equal(t, []string{"time"}, have)
	})
}

func Test_Value_Code(t *testing.T) {
	t.Run("nil instance", func(t *testing.T) {
		// --- Given ---
		var val *Value

		// --- When ---
		have := val.Code()

		// --- Then ---
		assert.Empty(t, have)
	})

	t.Run("without a package name", func(t *testing.T) {
		// --- Given ---
		val := NewValue("byte")

		// --- When ---
		have := val.Code()

		// --- Then ---
		assert.Equal(t, "byte", have)
	})

	t.Run("with a package name", func(t *testing.T) {
		// --- Given ---
		val := NewValue("time", "Duration")

		// --- When ---
		have := val.Code()

		// --- Then ---
		assert.Equal(t, "time.Duration", have)
	})

	t.Run("with a package name and prefix", func(t *testing.T) {
		// --- Given ---
		val := NewValue("-", "math", "MaxFloat32")

		// --- When ---
		have := val.Code()

		// --- Then ---
		assert.Equal(t, "-math.MaxFloat32", have)
	})
}
