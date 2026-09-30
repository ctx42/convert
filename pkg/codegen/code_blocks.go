// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

// -----------------------------------------------------------------------------
// ---------------------------- Converter Functions ----------------------------
// -----------------------------------------------------------------------------

// cbSrcToDstFuncDef defines a conversion function between two types.
var cbSrcToDstFuncDef = MustCodeBlock("cbSrcToDstFuncDef", `
{{.doc}}
func {{.src.Title}}To{{.dst.Title}}(src {{.src.Code}}) (dst {{.dst.Code}}, err error) {
`)

// -----------------------------------------------------------------------------

// cbAnyToDstFuncDef defines a conversion function between any type and a given
// type.
var cbAnyToDstFuncDef = MustCodeBlock("cbAnyToDstFuncDef", `
// AnyTo{{.dst.Title}} converts the given value to {{.dst.Doc}}
// using the package-level registry.
func AnyTo{{.dst.Title}}(value any, opts ...Option) ({{.dst.Code}}, error) {
`)

// -----------------------------------------------------------------------------

// cbAnyToDstFuncBody is a code block for the body of the function converting
// between any type and a given type.
var cbAnyToDstFuncBody = MustCodeBlock("cbAnyToDstFuncBody", `
import reflect

ops := NewOptions(opts...)
format := "%v: from %T to %v"
if ops.reg == nil {
	return 0, NewError(ErrNilRegistry, value, "{{.dst.Code}}").Format(format)
}
src := reflect.TypeOf(value)
wrp := ops.reg.lookup(src, typ{{.dst.Title}})
if wrp == nil {
	return 0, NewError(ErrUnkConv, value, "{{.dst.Code}}").Format(format)
}
ret, err := wrp.cst(value)
if err != nil {
{{- if .dst.IsAlias}}
	return 0, ChangeErrDstName(err, "{{.dst.Code}}")
{{- else}}
	return 0, err
{{- end}}
}
return ret.({{.dst.Code}}), nil //nolint:forcetypeassert
`)

// -----------------------------------------------------------------------------

// cbRegisterFunc defines a function registering the given converters in the
// package-level registry.
var cbRegisterFunc = MustCodeBlock("cbRegisterFunc", `
// registerNumeric registers the generated numeric converters in the
// package-level registry.
func registerNumeric() {
{{- range .names}}
	Register({{.}})
{{- end}}
}
`)

// -----------------------------------------------------------------------------

// cbCondition evaluates a specific condition.
var cbCondition = MustCodeBlock("cbCondition", `
if {{.var}} {{.cond}} {{.value.Code}} {
	return 0, NewError({{.error}}, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbIsWhole verifies that a value is a whole number (no fractional part).
var cbIsWhole = MustCodeBlock("cbIsWhole", `
import math

if {{ .var }} != math.Trunc({{ .var }}) {
	return 0, NewError(ErrFraction, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbIsFinite checks that a value is finite (not ±infinity).
var cbIsFinite = MustCodeBlock("cbIsFinite", `
import math

if math.IsInf({{.var}}, 0) {
	return 0, NewError(ErrInvValue, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbIsNumber checks a variable is a number (is not NaN).
var cbIsNumber = MustCodeBlock("cbIsNumber", `
import math

if math.IsNaN({{.var}}) {
	return 0, NewError(ErrInvValue, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbFloatRange checks that a finite floating-point value is within the
// destination floating-point type range.
var cbFloatRange = MustCodeBlock("cbFloatRange", `
import math

if !math.IsInf({{.var}}, 0) && math.Abs({{.var}}) > {{.value.Code}} {
	return 0, NewError(ErrInvRange, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbFloatExact checks that a floating-point value is exactly representable by
// the destination floating-point type. NaN is never equal to itself, so it is
// excluded from the check.
var cbFloatExact = MustCodeBlock("cbFloatExact", `
import math

if !math.IsNaN({{.var}}) && {{.src.Code}}({{.dst.Code}}({{.var}})) != {{.var}} {
	return 0, NewError(ErrInvSafeRange, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbCastToFloat64 casts a variable to float64 if needed.
var cbCastToFloat64 = MustCodeBlock("cbCastToFloat64", `
{{if eq .src.Name "float64" -}} 
f64 := {{.var}}
{{else -}}
f64 := float64({{.var}})
{{end -}}
`)

// -----------------------------------------------------------------------------

// cbReflectType defines a variable representing [reflect.Type] of a type.
var cbReflectType = MustCodeBlock("cbReflectType", `
// typ{{.type.Title}} is the [reflect.Type] of {{.type.Doc}}.
var typ{{.type.Title}} = reflect.TypeFor[{{.type.Code}}]()
`)

// -----------------------------------------------------------------------------
// ------------------------- Converter Functions Tests -------------------------
// -----------------------------------------------------------------------------

// cbTstFloatToFloat defines a test for a conversion function between two
// floating-point types, checking NaN and infinities are preserved.
var cbTstFloatToFloat = MustCodeBlock("cbTstFloatToFloat", `
import math

func Test_{{.src.Title}}To{{.dst.Title}}(t *testing.T) {
	t.Run("NaN", func(t *testing.T) {
		// --- Given ---
		src := {{.nan}}

		// --- When ---
		have, err := {{.src.Title}}To{{.dst.Title}}(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.True(t, math.IsNaN(float64(have)))
	})

	t.Run("positive infinity", func(t *testing.T) {
		// --- Given ---
		src := {{.pinf}}

		// --- When ---
		have, err := {{.src.Title}}To{{.dst.Title}}(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.True(t, math.IsInf(float64(have), 1))
	})

	t.Run("negative infinity", func(t *testing.T) {
		// --- Given ---
		src := {{.ninf}}

		// --- When ---
		have, err := {{.src.Title}}To{{.dst.Title}}(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.True(t, math.IsInf(float64(have), -1))
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstErrPrecision is an error test case when a value is not exactly
// representable by the destination type.
var cbTstErrPrecision = MustCodeBlock("cbTstErrPrecision", `
{
	"error - precision loss",
	0.1,
	ErrInvSafeRange,
	"value out of safe range: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstSrcToDstTT defines tabular success test cases for conversion functions
// between two types.
var cbTstSrcToDstTT = MustCodeBlock("cbTstSrcToDstTT", `
tt := []struct {
	testN string

	value {{.src.Code}}
	want  {{.dst.Code}}
}{
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDst defines a test for a conversion function between any type
// and a given type, run with a nil registry.
var cbTstAnyToDst = MustCodeBlock("cbTstAnyToDst", `
func Test_AnyTo{{.dst.Title}}(t *testing.T) {
	t.Run("error - nil registry", func(t *testing.T) {
		// --- Given ---
		opt := WithRegistry(nil)

		// --- When ---
		have, err := AnyTo{{.dst.Title}}(42, opt)

		// --- Then ---
		assert.ErrorIs(t, ErrNilRegistry, err)
		assert.ErrorEqual(t, "nil registry: from int to {{.dst.Code}}", err)
		assert.Equal(t, {{.dst.Code}}(0), have)
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstTT defines tabular success test cases for conversion functions
// between any type and a given type.
var cbTstAnyToDstTT = MustCodeBlock("cbTstAnyToDstTT", `
tt := []struct {
	testN string

	value any
	want  {{.dst.Code}}
}{
`)

// -----------------------------------------------------------------------------

// cbTstErrTT defines tabular error test cases for conversion functions.
var cbTstErrTT = MustCodeBlock("cbTstErrTT", `
tt := []struct {
	testN string

	value {{.carrier}}
	err   error
	msg   string
}{
`)

// -----------------------------------------------------------------------------

// cbTstPlatformSkip skips a test on platforms other than the ones with the
// given word size.
var cbTstPlatformSkip = MustCodeBlock("cbTstPlatformSkip", `
import strconv

if strconv.IntSize != {{.size}} {
	t.Skip("{{.size}}-bit platforms only")
}
`)

// -----------------------------------------------------------------------------

// cbTstSrcToDstLoop defines a tabular success test loop for conversion
// functions between two types.
var cbTstSrcToDstLoop = MustCodeBlock("cbTstSrcToDstLoop", `
for _, tc := range tt {
	t.Run(tc.testN, func(t *testing.T) {
		// --- When ---
		have, err := {{.src.Title}}To{{.dst.Title}}(tc.value)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, tc.want, have)
		assert.Equal(t, tc.value, {{.src.Code}}(have))
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstSrcToDstErrLoop defines a tabular error test loop for conversion
// functions between two types. Test case values stored in a carrier type
// other than the source type are converted before the call.
var cbTstSrcToDstErrLoop = MustCodeBlock("cbTstSrcToDstErrLoop", `
for _, tc := range tt {
	t.Run(tc.testN, func(t *testing.T) {
{{- if ne .carrier .src.Code}}
		// --- Given ---
		src := {{.src.Code}}(tc.value)

		// --- When ---
		have, err := {{.src.Title}}To{{.dst.Title}}(src)
{{- else}}
		// --- When ---
		have, err := {{.src.Title}}To{{.dst.Title}}(tc.value)
{{- end}}

		// --- Then ---
		assert.ErrorIs(t, tc.err, err)
		assert.ErrorEqual(t, tc.msg, err)
		assert.Equal(t, {{.dst.Code}}(0), have)
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstLoop defines a tabular success test loop for conversion
// functions between any type and a given type.
var cbTstAnyToDstLoop = MustCodeBlock("cbTstAnyToDstLoop", `
for _, tc := range tt {
	t.Run(tc.testN, func(t *testing.T) {
		// --- When ---
		have, err := AnyTo{{.dst.Title}}(tc.value)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, tc.want, have)
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstErrLoop defines a tabular error test loop for conversion
// functions between any type and a given type.
var cbTstAnyToDstErrLoop = MustCodeBlock("cbTstAnyToDstErrLoop", `
for _, tc := range tt {
	t.Run(tc.testN, func(t *testing.T) {
		// --- When ---
		have, err := AnyTo{{.dst.Title}}(tc.value)

		// --- Then ---
		assert.ErrorIs(t, tc.err, err)
		assert.ErrorEqual(t, tc.msg, err)
		assert.Equal(t, {{.dst.Code}}(0), have)
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstSrcToDstSuccess is a success conversion test case between two types.
var cbTstSrcToDstSuccess = MustCodeBlock("cbTstSrcToDstSuccess", `
{"{{.name}}", {{.src_value}}, {{.dst_value}}},
`)

// -----------------------------------------------------------------------------

// cbTstErrUnderSafeRange is an out of the safe range underflow error test case.
var cbTstErrUnderSafeRange = MustCodeBlock("cbTstErrUnderSafeRange", `
{
	"error - safe underflow",
	{{.min.Code}} - 1,
	ErrInvSafeRange,
	"value out of safe range: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstErrOverSafeRange is an out of the safe range overflow error test case.
var cbTstErrOverSafeRange = MustCodeBlock("cbTstErrOverSafeRange", `
{
	"error - safe overflow",
	{{.max.Code}} + 1,
	ErrInvSafeRange,
	"value out of safe range: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstErrUnderflow is an out-of-the range underflow error test case.
var cbTstErrUnderflow = MustCodeBlock("cbTstErrUnderflow", `
{
	"error - underflow",
	{{.min.Code}} - 1,
	ErrInvRange,
	"value out of range: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstErrOverflow is an out-of-the range overflow error test case.
var cbTstErrOverflow = MustCodeBlock("cbTstErrOverflow", `
{
	"error - overflow",
	{{.max.Code}} + 1,
	ErrInvRange,
	"value out of range: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstErrIsWhole is an error test case when a fractional value is given
// instead of a whole number.
var cbTstErrIsWhole = MustCodeBlock("cbTstErrIsWhole", `
{
	"error - fraction",
	4.2,
	ErrFraction,
	"must be a whole number: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstErrInvalidValue is an error test case when an invalid value is given.
var cbTstErrInvalidValue = MustCodeBlock("cbTstErrInvalidValue", `
{
	"error - {{.name}}",
	{{.value}},
	ErrInvValue,
	"invalid value: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstErrInvalidRange is an error test case when an invalid out-of-range
// value is given.
var cbTstErrInvalidRange = MustCodeBlock("cbTstErrInvalidRange", `
{
	"error - {{.name}}",
	{{.value}},
	ErrInvRange,
	"value out of range: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstErrInfinite is an error test case when an invalid infinity value is
// given.
var cbTstErrInfinite = MustCodeBlock("cbTstErrInfinite", `
import math

{
	"error - {{if ge (.sign) 0}}positive{{else}}negative{{end}} infinity",
	{{if eq .src.Size 64 -}} 
	math.Inf({{.sign}}),
	{{- else -}}
	float{{.src.Size}}(math.Inf({{.sign}})),
	{{- end}}
	ErrInvValue,
	"invalid value: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstFloatToInt is a test case for converting a float64 to an
// integer.
var cbTstAnyToDstFloatToInt = MustCodeBlock("cbTstAnyToDstFloatToInt", `
{"success from float64", 42.0, 42},
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstIntToFloat is a test case for converting an integer to a float64.
var cbTstAnyToDstIntToFloat = MustCodeBlock("cbTstAnyToDstIntToFloat", `
{"success from integer", 42, 42.0},
`)

// -----------------------------------------------------------------------------

// cbTstErrUndefinedConv is an error test case when an undefined conversion is
// attempted.
var cbTstErrUndefinedConv = MustCodeBlock("cbTstErrUndefinedConv", `
import github.com/ctx42/convert/internal/test

{
	"error - undefined conversion",
	test.Type{},
	ErrUnkConv,
	"conversion undefined: from test.Type to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------
