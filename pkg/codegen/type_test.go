// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NumericType_tabular(t *testing.T) {
	tt := []struct {
		testN string

		typ      func() Type
		value    *Value
		size     int
		bits     int
		signed   bool
		numeric  bool
		float    bool
		alias    *Value
		platform bool
	}{
		{
			testN:    "signed integer",
			typ:      func() Type { return NumericType[int]() },
			value:    NewValue("int"),
			size:     IntSize,
			bits:     IntSize - 1,
			signed:   true,
			numeric:  true,
			float:    false,
			alias:    nil,
			platform: true,
		},
		{
			testN:    "unsigned integer",
			typ:      func() Type { return NumericType[uint]() },
			value:    NewValue("uint"),
			size:     IntSize,
			bits:     IntSize,
			signed:   false,
			numeric:  true,
			float:    false,
			alias:    nil,
			platform: true,
		},
		{
			testN:    "uintptr",
			typ:      func() Type { return NumericType[uintptr]() },
			value:    NewValue("uintptr"),
			size:     IntSize,
			bits:     IntSize,
			signed:   false,
			numeric:  true,
			float:    false,
			alias:    nil,
			platform: true,
		},
		{
			testN:    "floating-point number",
			typ:      func() Type { return NumericType[float32]() },
			value:    NewValue("float32"),
			size:     32,
			bits:     24,
			signed:   true,
			numeric:  true,
			float:    true,
			alias:    nil,
			platform: false,
		},
		{
			testN:    "Uint64 type",
			typ:      func() Type { return NumericType[Uint64]() },
			value:    NewValue("github.com/ctx42/convert/pkg/codegen", "Uint64"),
			size:     64,
			bits:     64,
			signed:   false,
			numeric:  true,
			float:    false,
			alias:    nil,
			platform: false,
		},
		{
			testN:    "Int64 type",
			typ:      func() Type { return NumericType[Int64]() },
			value:    NewValue("github.com/ctx42/convert/pkg/codegen", "Int64"),
			size:     64,
			bits:     63,
			signed:   true,
			numeric:  true,
			float:    false,
			alias:    nil,
			platform: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.typ()

			// --- Then ---
			assert.Equal(t, tc.value, have.value)
			assert.Equal(t, tc.size, have.size)
			assert.Equal(t, tc.bits, have.bits)
			assert.Equal(t, tc.signed, have.signed)
			assert.Equal(t, tc.numeric, have.numeric)
			assert.Equal(t, tc.float, have.float)
			assert.Equal(t, tc.alias, have.alias)
			assert.Equal(t, tc.platform, have.platform)
			assert.Fields(t, 8, have)
		})
	}
}

func Test_Type_Alias(t *testing.T) {
	t.Run("set name only", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[uint8]()

		// --- When ---
		have := typ.Alias("byte")

		// --- Then ---
		assert.Equal(t, NewValue("byte"), have.value)
		assert.Equal(t, 8, have.size)
		assert.Equal(t, 8, have.bits)
		assert.True(t, have.numeric)
		assert.False(t, have.signed)
		assert.False(t, have.float)
		assert.Equal(t, NewValue("uint8"), have.alias)
		assert.False(t, have.platform)
		assert.Fields(t, 8, have)
	})

	t.Run("set package and name", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[int64]()

		// --- When ---
		have := typ.Alias("time", "Duration")

		// --- Then ---
		assert.Equal(t, NewValue("time", "Duration"), have.value)
		assert.Equal(t, 64, have.size)
		assert.Equal(t, 63, have.bits)
		assert.True(t, have.numeric)
		assert.True(t, have.signed)
		assert.False(t, have.float)
		assert.Equal(t, NewValue("int64"), have.alias)
		assert.False(t, have.platform)
		assert.Fields(t, 8, have)
	})

	t.Run("panics when no arguments", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[int64]()

		// --- When ---
		msg := assert.PanicMsg(t, func() { typ.Alias("") })

		// --- Then ---
		assert.Equal(t, "invalid number of arguments for Type.Alias", *msg)
	})

	t.Run("panics when more than two arguments", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[int64]()

		// --- When ---
		msg := assert.PanicMsg(t, func() { typ.Alias("pref", "two", "three") })

		// --- Then ---
		assert.Equal(t, "invalid number of arguments for Type.Alias", *msg)
	})
}

func Test_Type_IsAlias(t *testing.T) {
	t.Run("not an alias", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[uint8]()

		// --- When ---
		have := typ.IsAlias()

		// --- Then ---
		assert.False(t, have)
	})

	t.Run("is alias", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[uint8]().Alias("byte")

		// --- When ---
		have := typ.IsAlias()

		// --- Then ---
		assert.True(t, have)
	})
}

func Test_Type_Name(t *testing.T) {
	// --- Given ---
	typ := Type{value: NewValue("byte")}

	// --- When ---
	have := typ.Name()

	// --- Then ---
	assert.Equal(t, "byte", have)
}

func Test_Type_Title(t *testing.T) {
	// --- Given ---
	typ := Type{value: NewValue("byte")}

	// --- When ---
	have := typ.Title()

	// --- Then ---
	assert.Equal(t, "Byte", have)
}

func Test_Type_Doc(t *testing.T) {
	t.Run("build int type", func(t *testing.T) {
		// --- Given ---
		typ := Type{value: NewValue("byte")}

		// --- When ---
		have := typ.Doc()

		// --- Then ---
		assert.Equal(t, "byte", have)
	})

	t.Run("imported type", func(t *testing.T) {
		// --- Given ---
		typ := Type{value: NewValue("time", "Time")}

		// --- When ---
		have := typ.Doc()

		// --- Then ---
		assert.Equal(t, "[time.Time]", have)
	})
}

func Test_Type_Code(t *testing.T) {
	t.Run("without a package name", func(t *testing.T) {
		// --- Given ---
		typ := Type{value: NewValue("byte")}

		// --- When ---
		have := typ.Code()

		// --- Then ---
		assert.Equal(t, "byte", have)
	})

	t.Run("with a package name", func(t *testing.T) {
		// --- Given ---
		typ := Type{value: NewValue("time", "Duration")}

		// --- When ---
		have := typ.Code()

		// --- Then ---
		assert.Equal(t, "time.Duration", have)
	})
}

func Test_Type_Imports(t *testing.T) {
	t.Run("without import", func(t *testing.T) {
		// --- Given ---
		typ := Type{}

		// --- When ---
		have := typ.Imports()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("with import", func(t *testing.T) {
		// --- Given ---
		typ := Type{value: NewValue("time", "Duration")}

		// --- When ---
		have := typ.Imports()

		// --- Then ---
		assert.Equal(t, []string{"time"}, have)
	})
}

func Test_Type_IsNumeric(t *testing.T) {
	t.Run("numeric", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: true}

		// --- When ---
		have := typ.IsNumeric()

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("not numeric", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: false}

		// --- When ---
		have := typ.IsNumeric()

		// --- Then ---
		assert.False(t, have)
	})
}

func Test_Type_IsSigned(t *testing.T) {
	t.Run("signed", func(t *testing.T) {
		// --- Given ---
		typ := Type{signed: true}

		// --- When ---
		have := typ.IsSigned()

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("unsigned", func(t *testing.T) {
		// --- Given ---
		typ := Type{signed: false}

		// --- When ---
		have := typ.IsSigned()

		// --- Then ---
		assert.False(t, have)
	})
}

func Test_Type_IsUnsigned(t *testing.T) {
	t.Run("signed", func(t *testing.T) {
		// --- Given ---
		typ := Type{signed: true}

		// --- When ---
		have := typ.IsUnsigned()

		// --- Then ---
		assert.False(t, have)
	})

	t.Run("unsigned", func(t *testing.T) {
		// --- Given ---
		typ := Type{signed: false}

		// --- When ---
		have := typ.IsUnsigned()

		// --- Then ---
		assert.True(t, have)
	})
}

func Test_Type_IsFloat(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: true, float: true}

		// --- When ---
		have := typ.IsFloat()

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("integer", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: true, float: false}

		// --- When ---
		have := typ.IsFloat()

		// --- Then ---
		assert.False(t, have)
	})

	t.Run("not numeric", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: false, float: true}

		// --- When ---
		have := typ.IsFloat()

		// --- Then ---
		assert.False(t, have)
	})
}

func Test_Type_IsInteger(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: true, float: false}

		// --- When ---
		have := typ.IsInteger()

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("float", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: true, float: true}

		// --- When ---
		have := typ.IsInteger()

		// --- Then ---
		assert.False(t, have)
	})

	t.Run("not numeric", func(t *testing.T) {
		// --- Given ---
		typ := Type{numeric: false, float: false}

		// --- When ---
		have := typ.IsInteger()

		// --- Then ---
		assert.False(t, have)
	})
}

func Test_Type_IsPlatform(t *testing.T) {
	t.Run("platform", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[uint]()

		// --- When ---
		have := typ.IsPlatform()

		// --- Then ---
		assert.True(t, have)
	})

	t.Run("not platform", func(t *testing.T) {
		// --- Given ---
		typ := NumericType[uint64]()

		// --- When ---
		have := typ.IsPlatform()

		// --- Then ---
		assert.False(t, have)
	})
}

func Test_Type_Size(t *testing.T) {
	// --- Given ---
	typ := Type{size: 42}

	// --- When ---
	have := typ.Size()

	// --- Then ---
	assert.Equal(t, 42, have)
}

func Test_Type_onPlatform_tabular(t *testing.T) {
	tt := []struct {
		testN string

		typ  Type
		size int
		bits int
		want int
	}{
		{"int on 32-bit", NumericType[int](), 32, 31, 32},
		{"int on 64-bit", NumericType[int](), 64, 63, 64},
		{"uint on 32-bit", NumericType[uint](), 32, 32, 32},
		{"uintptr on 32-bit", NumericType[uintptr](), 32, 32, 32},
		{"int16 unchanged", NumericType[int16](), 32, 15, 16},
		{"float64 unchanged", NumericType[float64](), 32, 53, 64},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.typ.onPlatform(tc.size)

			// --- Then ---
			assert.Equal(t, tc.want, have.Size())
			assert.Equal(t, tc.bits, have.Bits())
			assert.Equal(t, tc.typ.Code(), have.Code())
			assert.Equal(t, tc.typ.IsPlatform(), have.IsPlatform())
		})
	}
}

func Test_Type_ConvActions_tabular(t *testing.T) {
	tt := []struct {
		testN string

		from    Type
		to      Type
		actions []Action
	}{
		{
			"nil actions when 'from' is not numeric",
			Type{numeric: true},
			Type{numeric: false},
			nil,
		},
		{
			"nil actions when 'to' is not numeric",
			Type{numeric: false},
			Type{numeric: true},
			nil,
		},

		{
			"int64 to float32",
			NumericType[int64](),
			NumericType[float32](),
			[]Action{
				NewAction(CheckIntSafeToFloatMin, NewValue("Float32SafeIntMin")),
				NewAction(CheckIntSafeToFloatMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int64 to float64",
			NumericType[int64](),
			NumericType[float64](),
			[]Action{
				NewAction(CheckIntSafeToFloatMin, NewValue("Float64SafeIntMin")),
				NewAction(CheckIntSafeToFloatMax, NewValue("Float64SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"int16 to int32",
			NumericType[int16](),
			NumericType[int32](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int32 to int32",
			NumericType[int32](),
			NumericType[int32](),
			[]Action{
				NewAction(CastNotNeeded, nil),
			},
		},
		{
			"int64 to int32",
			NumericType[int64](),
			NumericType[int32](),
			[]Action{
				NewAction(CheckUnderflows, NewValue("math", "MinInt32")),
				NewAction(CheckOverflows, NewValue("math", "MaxInt32")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"uint16 to int32",
			NumericType[uint16](),
			NumericType[int32](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint32 to int32",
			NumericType[uint32](),
			NumericType[int32](),
			[]Action{
				NewAction(CheckOverflows, NewValue("math", "MaxInt32")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint64 to int32",
			NumericType[uint64](),
			NumericType[int32](),
			[]Action{
				NewAction(CheckOverflows, NewValue("math", "MaxInt32")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"uint16 to uint32",
			NumericType[uint16](),
			NumericType[uint32](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint32 to uint32",
			NumericType[uint32](),
			NumericType[uint32](),
			[]Action{
				NewAction(CastNotNeeded, nil),
			},
		},
		{
			"uint64 to uint32",
			NumericType[uint64](),
			NumericType[uint32](),
			[]Action{
				NewAction(CheckOverflows, NewValue("math", "MaxUint32")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"int16 to uint32",
			NumericType[int16](),
			NumericType[uint32](),
			[]Action{
				NewAction(CheckIsNonNegative, nil),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int32 to uint32",
			NumericType[int32](),
			NumericType[uint32](),
			[]Action{
				NewAction(CheckIsNonNegative, nil),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int64 to uint32",
			NumericType[int64](),
			NumericType[uint32](),
			[]Action{
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckOverflows, NewValue("math", "MaxUint32")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"int to int",
			NumericType[int](),
			NumericType[int](),
			[]Action{
				NewAction(CastNotNeeded, nil),
			},
		},
		{
			"int to int64",
			NumericType[int](),
			NumericType[int64](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int64 to int",
			NumericType[int64](),
			NumericType[int](),
			[]Action{
				NewAction(CheckUnderflows, NewValue("math", "MinInt")),
				NewAction(CheckOverflows, NewValue("math", "MaxInt")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint to int64",
			NumericType[uint](),
			NumericType[int64](),
			[]Action{
				NewAction(CheckOverflows, NewValue("math", "MaxInt64")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint to int",
			NumericType[uint](),
			NumericType[int](),
			[]Action{
				NewAction(CheckOverflows, NewValue("math", "MaxInt")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int64 to uint",
			NumericType[int64](),
			NumericType[uint](),
			[]Action{
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckOverflows, NewValue("math", "MaxUint")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint64 to uintptr",
			NumericType[uint64](),
			NumericType[uintptr](),
			[]Action{
				NewAction(CheckOverflows, NewValue("MaxUintptr")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int to float32",
			NumericType[int](),
			NumericType[float32](),
			[]Action{
				NewAction(CheckIntSafeToFloatMin, NewValue("Float32SafeIntMin")),
				NewAction(CheckIntSafeToFloatMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int to float64",
			NumericType[int](),
			NumericType[float64](),
			[]Action{
				NewAction(CheckIntSafeToFloatMin, NewValue("Float64SafeIntMin")),
				NewAction(CheckIntSafeToFloatMax, NewValue("Float64SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float64 to int",
			NumericType[float64](),
			NumericType[int](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckUnderflows, NewValue("math", "MinInt")),
				NewAction(CheckOverflows, NewValue("math", "MaxInt")),
				NewAction(CheckFloatSafeToIntMin, NewValue("Float64SafeIntMin")),
				NewAction(CheckFloatSafeToIntMax, NewValue("Float64SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.from.ConvActions(tc.to)

			// --- Then ---
			assert.Equal(t, tc.actions, have)
		})
	}
}

func Test_Type_platformConvActions_tabular(t *testing.T) {
	tt := []struct {
		testN string

		from    Type
		to      Type
		size    int
		actions []Action
	}{
		{
			"int64 to int on 32-bit",
			NumericType[int64](),
			NumericType[int](),
			32,
			[]Action{
				NewAction(CheckUnderflows, NewValue("math", "MinInt32")),
				NewAction(CheckOverflows, NewValue("math", "MaxInt32")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int64 to int on 64-bit",
			NumericType[int64](),
			NumericType[int](),
			64,
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint to int64 on 32-bit",
			NumericType[uint](),
			NumericType[int64](),
			32,
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint to int64 on 64-bit",
			NumericType[uint](),
			NumericType[int64](),
			64,
			[]Action{
				NewAction(CheckOverflows, NewValue("math", "MaxInt64")),
				NewAction(CastDirectly, nil),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.from.platformConvActions(tc.to, tc.size)

			// --- Then ---
			assert.Equal(t, tc.actions, have)
		})
	}
}

func Test_Type_floatConvActions_tabular(t *testing.T) {
	// NOTE: When adding cases to this case
	// add tehm also to Test_Generator_Generate.

	tt := []struct {
		testN string

		from    Type
		to      Type
		actions []Action
	}{
		{
			"nil actions when 'from' is not numeric",
			Type{numeric: true},
			Type{numeric: false},
			nil,
		},
		{
			"nil actions when 'to' is not numeric",
			Type{numeric: false},
			Type{numeric: true},
			nil,
		},
		{
			"nil actions when neither 'from' nor 'to' are a floats",
			Type{numeric: true, float: false},
			Type{numeric: true, float: false},
			nil,
		},

		{
			"float32 to float64",
			NumericType[float32](),
			NumericType[float64](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float64 to float64",
			NumericType[float64](),
			NumericType[float64](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float32 to float32",
			NumericType[float32](),
			NumericType[float32](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float64 to float32",
			NumericType[float64](),
			NumericType[float32](),
			[]Action{
				NewAction(CheckFloatRange, NewValue("math", "MaxFloat32")),
				NewAction(CheckFloatExact, nil),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"float64 to uint64",
			NumericType[float64](),
			NumericType[uint64](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckFloatSafeToIntMax, NewValue("Float64SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float64 to uint32",
			NumericType[float64](),
			NumericType[uint32](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckOverflows, NewValue("math", "MaxUint32")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float64 to uint16",
			NumericType[float64](),
			NumericType[uint16](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckOverflows, NewValue("math", "MaxUint16")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"float32 to uint64",
			NumericType[float32](),
			NumericType[uint64](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckFloatSafeToIntMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float32 to uint32",
			NumericType[float32](),
			NumericType[uint32](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckFloatSafeToIntMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float32 to uint16",
			NumericType[float32](),
			NumericType[uint16](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckIsNonNegative, nil),
				NewAction(CheckOverflows, NewValue("math", "MaxUint16")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"float64 to int64",
			NumericType[float64](),
			NumericType[int64](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckFloatSafeToIntMin, NewValue("Float64SafeIntMin")),
				NewAction(CheckFloatSafeToIntMax, NewValue("Float64SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float64 to int32",
			NumericType[float64](),
			NumericType[int32](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckUnderflows, NewValue("math", "MinInt32")),
				NewAction(CheckOverflows, NewValue("math", "MaxInt32")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"float32 to int",
			NumericType[float32](),
			NumericType[int64](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckFloatSafeToIntMin, NewValue("Float32SafeIntMin")),
				NewAction(CheckFloatSafeToIntMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float32 to int32",
			NumericType[float32](),
			NumericType[int32](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckFloatSafeToIntMin, NewValue("Float32SafeIntMin")),
				NewAction(CheckFloatSafeToIntMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"float32 to int16",
			NumericType[float32](),
			NumericType[int16](),
			[]Action{
				NewAction(CastToFloat64, nil),
				NewAction(CheckIsNumber, nil),
				NewAction(CheckIsFinite, nil),
				NewAction(CheckIsWhole, nil),
				NewAction(CheckUnderflows, NewValue("math", "MinInt16")),
				NewAction(CheckOverflows, NewValue("math", "MaxInt16")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"uint64 to float64",
			NumericType[uint64](),
			NumericType[float64](),
			[]Action{
				NewAction(CheckIntSafeToFloatMax, NewValue("Float64SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint64 to float32",
			NumericType[uint64](),
			NumericType[float32](),
			[]Action{
				NewAction(CheckIntSafeToFloatMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"uint32 to float64",
			NumericType[uint32](),
			NumericType[float64](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"uint32 to float32",
			NumericType[uint64](),
			NumericType[float32](),
			[]Action{
				NewAction(CheckIntSafeToFloatMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"int64 to float64",
			NumericType[int64](),
			NumericType[float64](),
			[]Action{
				NewAction(CheckIntSafeToFloatMin, NewValue("Float64SafeIntMin")),
				NewAction(CheckIntSafeToFloatMax, NewValue("Float64SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int64 to float32",
			NumericType[int64](),
			NumericType[float32](),
			[]Action{
				NewAction(CheckIntSafeToFloatMin, NewValue("Float32SafeIntMin")),
				NewAction(CheckIntSafeToFloatMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},

		{
			"int32 to float64",
			NumericType[int32](),
			NumericType[float64](),
			[]Action{
				NewAction(CastDirectly, nil),
			},
		},
		{
			"int32 to float32",
			NumericType[int32](),
			NumericType[float32](),
			[]Action{
				NewAction(CheckIntSafeToFloatMin, NewValue("Float32SafeIntMin")),
				NewAction(CheckIntSafeToFloatMax, NewValue("Float32SafeIntMax")),
				NewAction(CastDirectly, nil),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.from.floatConvActions(tc.to)

			// --- Then ---
			assert.Equal(t, tc.actions, have)
		})
	}
}
