// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"fmt"
	"io"
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
				"var":   "src",
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
				"var":   "src",
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
				"var":   "src",
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
				"var":   "src",
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

// testFunc generates code for the conversion function tests.
func (gen *GenSrcToDst) testFunc() error {
	gen.addImport("testing", "github.com/ctx42/testing/pkg/assert")
	gen.writeCode("func ")

	format := "Test_%sTo%s_tabular"
	gen.writeCode(format, gen.src.Title(), gen.dst.Title())

	gen.writeCode("(t *testing.T) {\n")
	if err := gen.testFuncBody(); err != nil {
		return err
	}
	gen.writeCode("}\n")
	return nil
}

// testFuncBody generates code for the conversion function tests.
func (gen *GenSrcToDst) testFuncBody() error {
	data := map[string]any{"src": gen.src, "dst": gen.dst}
	gen.addImport(cbTstSrcToDstTT.Imports()...)
	if err := cbTstSrcToDstTT.Render(gen.code, 1, data); err != nil {
		return err
	}
	if err := gen.testCases(); err != nil {
		return err
	}
	gen.writeCode("\t}\n\n")

	data = map[string]any{"src": gen.src, "dst": gen.dst}
	gen.addImport(cbTstSrcToDstLoop.Imports()...)
	return cbTstSrcToDstLoop.Render(gen.code, 1, data)
}

// testCases generates code for the conversion function test cases.
//
// nolint: cyclop, gocognit
func (gen *GenSrcToDst) testCases() error {
	for _, act := range gen.src.ConvActions(gen.dst) {
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
