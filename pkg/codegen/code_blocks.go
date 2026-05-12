// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

// -----------------------------------------------------------------------------
// ---------------------------- Converter Functions ----------------------------
// -----------------------------------------------------------------------------

// cbSrcToDstFuncDef defines a conversion function between two types.
var cbSrcToDstFuncDef = MustCodeBlock("cbSrcToDstFuncDef", `
// {{.src.Title}}To{{.dst.Title}} safely converts {{.src.Doc}} value to {{.dst.Doc}}.
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
src := reflect.TypeOf(value)
wrp := ops.reg.lookup(src, typ{{.dst.Title}})
if wrp == nil {
	format := "%v: from %T to %v"
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
return ret.({{.dst.Code}}), nil // nolint: forcetypeassert
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
var cbIsWhole = MustCodeBlock("cbCheckIsWhole", `
import math

if {{ .var }} != math.Trunc({{ .var }}) {
	return 0, NewError(ErrFraction, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbIsFinite checks that a value is finite (not ±infinity).
var cbIsFinite = MustCodeBlock("cbChekIsFinite", `
import math

if math.IsInf({{.var}}, 0) {
	return 0, NewError(ErrInvValue, "{{.src.Code}}", "{{.dst.Code}}")
}
`)

// -----------------------------------------------------------------------------

// cbIsNumber checks a variable is a number (is not NaN).
var cbIsNumber = MustCodeBlock("cbChekIsNumber", `
import math

if math.IsNaN({{.var}}) {
	return 0, NewError(ErrInvValue, "{{.src.Code}}", "{{.dst.Code}}")
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
// typ{{.type.Title}} is reflected {{.type.Doc}}.
var typ{{.type.Title}} = reflect.TypeFor[{{.type.Code}}]()
`)

// -----------------------------------------------------------------------------
// ------------------------- Converter Functions Tests -------------------------
// -----------------------------------------------------------------------------

// cbTstSrcToDstTT defines tabular test cases for conversion functions between
// two types.
var cbTstSrcToDstTT = MustCodeBlock("cbTstSrcToDstTT", `
tt := []struct {
	testN string

	value {{.src.Code}}
	want  {{.dst.Code}}
	err   error
	msg   string
}{
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstTT defines tabular test cases for conversion functions between
// any type and a given type.
var cbTstAnyToDstTT = MustCodeBlock("cbTstAnyToDstTT", `
tt := []struct {
	testN string

	value any
	want  {{.dst.Code}}
	err   error
	msg   string
}{
`)

// -----------------------------------------------------------------------------

// cbTstSrcToDstLoop defines a tabular test loop for conversion functions
// between two types.
var cbTstSrcToDstLoop = MustCodeBlock("cbTstSrcToDstLoop", `
for _, tc := range tt {
	t.Run(tc.testN, func(t *testing.T) {
		// --- When ---
		have, err := {{.src.Title}}To{{.dst.Title}}(tc.value)

		// --- Then ---
		if tc.err == nil {
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
			assert.Equal(t, tc.value, {{.src.Code}}(have))
			return
		}

		assert.ErrorIs(t, tc.err, err)
		assert.ErrorEqual(t, tc.msg, err)
		assert.Equal(t, {{.dst.Code}}(0), have)
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstLoop defines a tabular test loop for conversion functions
// between any type two a given type.
var cbTstAnyToDstLoop = MustCodeBlock("cbTstAnyToDstLoop", `
for _, tc := range tt {
	t.Run(tc.testN, func(t *testing.T) {
		// --- When ---
		have, err := AnyTo{{.dst.Title}}(tc.value)

		// --- Then ---
		if tc.err == nil {
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
			return
		}

		assert.ErrorIs(t, tc.err, err)
		assert.ErrorEqual(t, tc.msg, err)
		assert.Equal(t, {{.dst.Code}}(0), have)
	})
}
`)

// -----------------------------------------------------------------------------

// cbTstSrcToDstSuccess is a success conversion test case between two types.
var cbTstSrcToDstSuccess = MustCodeBlock("cbTstSrcToDstSuccess", `
{"{{.name}}", {{.src_value}}, {{.dst_value}}, nil, ""},
`)

// -----------------------------------------------------------------------------

// cbTstErrUnderSafeRange is an out of the safe range underflow error test case.
var cbTstErrUnderSafeRange = MustCodeBlock("cbTstErrUnderSafeRange", `
{
	"error - safe underflow",
	{{.min.Code}} - 1,
	0,
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
	0,
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
	0,
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
	0,
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
	0,
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
	0,
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
	0,
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
	0,
	ErrInvValue,
	"invalid value: from {{.src.Code}} to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstFloatToInt is a test case for converting a float64 to an
// integer.
var cbTstAnyToDstFloatToInt = MustCodeBlock("cbTstAnyToDstFloatToInt", `
{"success from float64", 42.0, 42, nil, ""},
`)

// -----------------------------------------------------------------------------

// cbTstAnyToDstIntToFloat is a test case for converting an integer to a float64.
var cbTstAnyToDstIntToFloat = MustCodeBlock("cbTstAnyToDstIntToFloat", `
{"success from integer", 42, 42.0, nil, ""},
`)

// -----------------------------------------------------------------------------

// cbTstErrUndefinedConv is an error test case when an undefined conversion is
// attempted.
var cbTstErrUndefinedConv = MustCodeBlock("cbTstErrUndefinedConv", `
import github.com/ctx42/convert/internal/test

{
	"error - undefined conversion",
	test.Type{},
	0,
	ErrUnkConv,
	"conversion undefined: from test.Type to {{.dst.Code}}",
},
`)

// -----------------------------------------------------------------------------
