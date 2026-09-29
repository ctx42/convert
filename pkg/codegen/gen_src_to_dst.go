// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"fmt"
	"io"
	"slices"
)

// GenSrcToDst represents the code generator for converters between values of
// two types and their tests.
type GenSrcToDst struct {
	file      // Represents a file with source code.
	src  Type // Source type.
	dst  Type // Destination type.
}

// NewGenSrcToDst returns a new [GenSrcToDst] instance.
func NewGenSrcToDst(pkg string, ops ...Option) *GenSrcToDst {
	return &GenSrcToDst{file: newFile(pkg, ops...)}
}

// GenerateCode generates source code for the conversion function.
func (gen *GenSrcToDst) GenerateCode(w io.Writer, src, dst Type) error {
	gen.reset()
	gen.src = src
	gen.dst = dst
	gen.addImport(src.Imports()...)
	gen.addImport(dst.Imports()...)
	if err := gen.convFunc(); err != nil {
		return err
	}
	if _, err := gen.WriteTo(w); err != nil {
		return err
	}
	return nil
}

// GenerateTest generates tests for the conversion function.
func (gen *GenSrcToDst) GenerateTest(w io.Writer, src, dst Type) error {
	gen.reset()
	gen.src = src
	gen.dst = dst
	gen.addImport(src.Imports()...)
	gen.addImport(dst.Imports()...)
	if err := gen.testFunc(); err != nil {
		return err
	}
	if _, err := gen.WriteTo(w); err != nil {
		return err
	}
	return nil
}

// convFunc generates code for the conversion function.
func (gen *GenSrcToDst) convFunc() error {
	data := map[string]any{"src": gen.src, "dst": gen.dst}
	gen.addImport(cbSrcToDstFuncDef.Imports()...)
	if err := cbSrcToDstFuncDef.Render(gen.code, 0, data); err != nil {
		return err
	}
	if err := gen.convFuncBody(gen.src.ConvActions(gen.dst)); err != nil {
		return err
	}
	gen.writeCode("}\n")
	return nil
}

// convFuncBody generates code for the numeric conversion function body.
//
// nolint: cyclop, gocognit
func (gen *GenSrcToDst) convFuncBody(actions []Action) error {
	if len(actions) == 0 {
		format := "no conversion actions found for %s to %s"
		return fmt.Errorf(format, gen.src.Code(), gen.dst.Code())
	}

	for _, act := range actions {
		gen.addImport(act.Imports()...)

		switch act.Name() {
		case CastNotNeeded:
			gen.writeCode("\treturn src, nil\n")

		case CastDirectly:
			gen.writeCode("\treturn %s(src), nil\n", gen.dst.Code())

		case CastToFloat64:
			data := map[string]any{"src": gen.src, "var": "src"}
			gen.addImport(cbCastToFloat64.Imports()...)
			if err := cbCastToFloat64.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckIsNonNegative:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   "src",
				"cond":  "<",
				"value": NewValue("0"),
				"error": "ErrInvRange",
			}
			gen.addImport(cbCondition.Imports()...)
			if err := cbCondition.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckUnderflows:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   gen.cmpVar(act.Name()),
				"cond":  "<",
				"value": act,
				"error": "ErrInvRange",
			}
			gen.addImport(cbCondition.Imports()...)
			if err := cbCondition.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckOverflows:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   gen.cmpVar(act.Name()),
				"cond":  ">",
				"value": act,
				"error": "ErrInvRange",
			}
			gen.addImport(cbCondition.Imports()...)
			if err := cbCondition.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckIsFinite:
			data := map[string]any{"src": gen.src, "dst": gen.dst, "var": "f64"}
			gen.addImport(cbIsFinite.Imports()...)
			if err := cbIsFinite.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckIsNumber:
			data := map[string]any{"var": "f64", "src": gen.src, "dst": gen.dst}
			gen.addImport(cbIsNumber.Imports()...)
			if err := cbIsNumber.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckIsWhole:
			data := map[string]any{"src": gen.src, "dst": gen.dst, "var": "f64"}
			gen.addImport(cbIsWhole.Imports()...)
			if err := cbIsWhole.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckIntSafeToFloatMin:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   gen.cmpVar(act.Name()),
				"cond":  "<",
				"value": act,
				"error": "ErrInvSafeRange",
			}
			gen.addImport(cbCondition.Imports()...)
			if err := cbCondition.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckIntSafeToFloatMax:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   gen.cmpVar(act.Name()),
				"cond":  ">",
				"value": act,
				"error": "ErrInvSafeRange",
			}
			gen.addImport(cbCondition.Imports()...)
			if err := cbCondition.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckFloatSafeToIntMin:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   "f64",
				"cond":  "<",
				"value": act,
				"error": "ErrInvSafeRange",
			}
			gen.addImport(cbCondition.Imports()...)
			if err := cbCondition.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckFloatRange:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   "src",
				"value": act,
			}
			gen.addImport(cbFloatRange.Imports()...)
			if err := cbFloatRange.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckFloatExact:
			data := map[string]any{"src": gen.src, "dst": gen.dst, "var": "src"}
			gen.addImport(cbFloatExact.Imports()...)
			if err := cbFloatExact.Render(gen.code, 1, data); err != nil {
				return err
			}

		case CheckFloatSafeToIntMax:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"var":   "f64",
				"cond":  ">",
				"value": act,
				"error": "ErrInvSafeRange",
			}
			gen.addImport(cbCondition.Imports()...)
			if err := cbCondition.Render(gen.code, 1, data); err != nil {
				return err
			}

		default:
			return fmt.Errorf("unsupported action: %s", act.Name())
		}
	}
	return nil
}

// cmpVar returns the expression representing the source value in the range
// check performed by the named action. When the source or the destination is
// a platform-sized type, an integer source is converted to a 64-bit type, so
// the limit it is compared with fits it on both 32-bit and 64-bit platforms.
func (gen *GenSrcToDst) cmpVar(name ActionName) string {
	src, dst := gen.src, gen.dst
	if src.IsFloat() || (!src.IsPlatform() && !dst.IsPlatform()) {
		return "src"
	}
	wide := NumericType[int64]()
	if src.IsUnsigned() || (name == CheckOverflows && dst.IsUnsigned()) {
		wide = NumericType[uint64]()
	}
	isWide := !src.IsPlatform() && src.Size() == 64
	if isWide && src.IsSigned() == wide.IsSigned() {
		return "src"
	}
	return wide.Code() + "(src)"
}

// testActions returns the conversion actions to test on every platform, and
// the ones to test only on 32-bit and only on 64-bit platforms.
func (gen *GenSrcToDst) testActions() (common, only32, only64 []Action) {
	if !gen.src.IsPlatform() && !gen.dst.IsPlatform() {
		return gen.src.ConvActions(gen.dst), nil, nil
	}
	a32 := gen.src.platformConvActions(gen.dst, 32)
	a64 := gen.src.platformConvActions(gen.dst, 64)
	for _, act := range a64 {
		if slices.ContainsFunc(a32, act.equal) {
			common = append(common, act)
			continue
		}
		only64 = append(only64, act)
	}
	for _, act := range a32 {
		if !slices.ContainsFunc(common, act.equal) {
			only32 = append(only32, act)
		}
	}
	return common, only32, only64
}

// testFunc generates code for the conversion function tests.
func (gen *GenSrcToDst) testFunc() error {
	common, only32, only64 := gen.testActions()

	gen.addImport("testing", "github.com/ctx42/testing/pkg/assert")
	if err := gen.floatTestFunc(); err != nil {
		return err
	}
	gen.writeCode("func ")

	format := "Test_%sTo%s_tabular"
	gen.writeCode(format, gen.src.Title(), gen.dst.Title())

	gen.writeCode("(t *testing.T) {\n")
	if err := gen.testFuncBody(common); err != nil {
		return err
	}
	gen.writeCode("}\n")

	if err := gen.platformTestFunc(32, only32); err != nil {
		return err
	}
	return gen.platformTestFunc(64, only64)
}

// floatTestFunc generates code for the conversion function test checking NaN
// and infinities are preserved when converting between two different
// floating-point types. Generates nothing for other conversions.
func (gen *GenSrcToDst) floatTestFunc() error {
	if !gen.src.IsFloat() || !gen.dst.IsFloat() {
		return nil
	}
	if gen.src.Code() == gen.dst.Code() {
		return nil
	}

	special := func(expr string) string {
		if gen.src.Size() == 64 {
			return expr
		}
		return gen.src.Code() + "(" + expr + ")"
	}
	data := map[string]any{
		"src":  gen.src,
		"dst":  gen.dst,
		"nan":  special("math.NaN()"),
		"pinf": special("math.Inf(1)"),
		"ninf": special("math.Inf(-1)"),
	}
	gen.addImport(cbTstFloatToFloat.Imports()...)
	if err := cbTstFloatToFloat.Render(gen.code, 0, data); err != nil {
		return err
	}
	gen.writeCode("\n")
	return nil
}

// testFuncBody generates code for the conversion function tests.
func (gen *GenSrcToDst) testFuncBody(actions []Action) error {
	data := map[string]any{"src": gen.src, "dst": gen.dst}
	gen.addImport(cbTstSrcToDstTT.Imports()...)
	if err := cbTstSrcToDstTT.Render(gen.code, 1, data); err != nil {
		return err
	}
	if err := gen.testCases(actions); err != nil {
		return err
	}
	gen.writeCode("\t}\n\n")

	gen.addImport(cbTstSrcToDstLoop.Imports()...)
	return cbTstSrcToDstLoop.Render(gen.code, 1, data)
}

// platformTestFunc generates code for the conversion function tests run only
// on platforms with the given word size in bits. Generates nothing when there
// are no actions to test.
func (gen *GenSrcToDst) platformTestFunc(size int, actions []Action) error {
	if len(actions) == 0 {
		return nil
	}

	// Test case values are stored in a 64-bit type, so values out of range
	// on the other platform compile.
	carrier := gen.src
	if gen.src.IsPlatform() {
		carrier = NumericType[int64]()
		if gen.src.IsUnsigned() {
			carrier = NumericType[uint64]()
		}
	}

	data := map[string]any{
		"src":     gen.src,
		"dst":     gen.dst,
		"size":    size,
		"carrier": carrier,
	}
	gen.writeCode("\n")
	gen.addImport(cbTstSrcToDstPlatformTT.Imports()...)
	if err := cbTstSrcToDstPlatformTT.Render(gen.code, 0, data); err != nil {
		return err
	}
	if err := gen.testCases(actions); err != nil {
		return err
	}
	gen.writeCode("\t}\n\n")

	gen.addImport(cbTstSrcToDstPlatformLoop.Imports()...)
	if err := cbTstSrcToDstPlatformLoop.Render(gen.code, 1, data); err != nil {
		return err
	}
	gen.writeCode("}\n")
	return nil
}

// testCases generates code for the conversion function test cases.
//
// nolint: cyclop, gocognit
func (gen *GenSrcToDst) testCases(actions []Action) error {
	for _, act := range actions {
		gen.addImport(act.Imports()...)

		switch act.Name() {
		case CastNotNeeded:
			val := MinValue(gen.src)
			gen.addImport(val.Imports()...)
			data := map[string]any{
				"name":      "min",
				"src_value": val.Code(),
				"dst_value": val.Code(),
			}
			gen.addImport(cbTstSrcToDstSuccess.Imports()...)
			if err := cbTstSrcToDstSuccess.Render(gen.code, 2, data); err != nil {
				return err
			}

			val = MaxValue(gen.src)
			gen.addImport(val.Imports()...)
			data = map[string]any{
				"name":      "max",
				"src_value": val.Code(),
				"dst_value": val.Code(),
			}
			gen.addImport(cbTstSrcToDstSuccess.Imports()...)

			if err := cbTstSrcToDstSuccess.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CastDirectly:
			data := map[string]any{
				"name":      "success",
				"src_value": 42,
				"dst_value": 42,
			}
			gen.addImport(cbTstSrcToDstSuccess.Imports()...)
			if err := cbTstSrcToDstSuccess.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CastToFloat64:
			// Nothing to do.

		case CheckIsNonNegative:
			data := map[string]any{
				"src":   gen.src,
				"dst":   gen.dst,
				"name":  "negative",
				"value": -1,
			}
			block := cbTstErrInvalidRange
			gen.addImport(block.Imports()...)
			if err := block.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CheckUnderflows:
			data := map[string]any{"min": act, "src": gen.src, "dst": gen.dst}
			gen.addImport(cbTstErrUnderflow.Imports()...)
			if err := cbTstErrUnderflow.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CheckOverflows:
			data := map[string]any{"max": act, "src": gen.src, "dst": gen.dst}
			gen.addImport(cbTstErrOverflow.Imports()...)
			if err := cbTstErrOverflow.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CheckIsFinite:
			data := map[string]any{"src": gen.src, "dst": gen.dst, "sign": -1}
			gen.addImport(cbTstErrInfinite.Imports()...)
			if err := cbTstErrInfinite.Render(gen.code, 2, data); err != nil {
				return err
			}

			data["sign"] = 1
			gen.addImport(cbTstErrInfinite.Imports()...)
			if err := cbTstErrInfinite.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CheckIsNumber:
			nan := "math.NaN()"
			if gen.src.Size() != 64 {
				nan = fmt.Sprintf("float%d(math.NaN())", gen.src.Size())
			}
			data := map[string]any{
				"name":  "must be a number",
				"value": nan,
				"src":   gen.src,
				"dst":   gen.dst,
			}
			tpl := cbTstErrInvalidValue
			gen.addImport(tpl.Imports()...)
			if err := tpl.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CheckIsWhole:
			data := map[string]any{"src": gen.src, "dst": gen.dst}
			gen.addImport(cbTstErrIsWhole.Imports()...)
			if err := cbTstErrIsWhole.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CheckFloatRange:
			gen.addImport(act.Imports()...)
			for _, row := range [][2]string{
				{"overflow", "math.MaxFloat64"},
				{"negative overflow", "-math.MaxFloat64"},
			} {
				data := map[string]any{
					"src":   gen.src,
					"dst":   gen.dst,
					"name":  row[0],
					"value": row[1],
				}
				block := cbTstErrInvalidRange
				gen.addImport(block.Imports()...)
				if err := block.Render(gen.code, 2, data); err != nil {
					return err
				}
			}

		case CheckFloatExact:
			data := map[string]any{"src": gen.src, "dst": gen.dst}
			gen.addImport(cbTstErrPrecision.Imports()...)
			if err := cbTstErrPrecision.Render(gen.code, 2, data); err != nil {
				return err
			}
			for _, row := range [][2]string{
				{"fraction", "0.5"},
				{"large", "1e10"},
				{"max", "math.MaxFloat32"},
			} {
				data := map[string]any{
					"name":      row[0],
					"src_value": row[1],
					"dst_value": row[1],
				}
				block := cbTstSrcToDstSuccess
				gen.addImport(block.Imports()...)
				if err := block.Render(gen.code, 2, data); err != nil {
					return err
				}
			}

		case CheckIntSafeToFloatMin, CheckFloatSafeToIntMin:
			data := map[string]any{"min": act, "src": gen.src, "dst": gen.dst}
			block := cbTstErrUnderSafeRange
			gen.addImport(block.Imports()...)
			if err := block.Render(gen.code, 2, data); err != nil {
				return err
			}

		case CheckIntSafeToFloatMax, CheckFloatSafeToIntMax:
			data := map[string]any{"max": act, "src": gen.src, "dst": gen.dst}
			block := cbTstErrOverSafeRange
			gen.addImport(block.Imports()...)
			if err := block.Render(gen.code, 2, data); err != nil {
				return err
			}
		}
	}

	return nil
}
