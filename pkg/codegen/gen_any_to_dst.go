// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"fmt"
	"io"
)

// GenAnyToDst represents the code generator for converters between values of
// any type to a given type and their tests.
type GenAnyToDst struct {
	file      // Represents a file with source code.
	dst  Type // Destination type.
}

// NewGenAnyToDst returns a new [GenAnyToDst] instance.
func NewGenAnyToDst(pkg string, ops ...Option) *GenAnyToDst {
	return &GenAnyToDst{file: newFile(pkg, ops...)}
}

// GenerateCode generates source code for the conversion function.
func (gen *GenAnyToDst) GenerateCode(w io.Writer, dst Type) error {
	gen.reset()
	gen.dst = dst
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
func (gen *GenAnyToDst) GenerateTest(w io.Writer, dst Type) error {
	gen.reset()
	gen.dst = dst
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
func (gen *GenAnyToDst) convFunc() error {
	data := map[string]any{"type": gen.dst}
	gen.addImport(cbReflectType.Imports()...)
	if err := cbReflectType.Render(gen.code, 0, data); err != nil {
		return err
	}
	gen.writeCode("\n")

	data = map[string]any{"dst": gen.dst}
	gen.addImport(cbAnyToDstFuncDef.Imports()...)
	if err := cbAnyToDstFuncDef.Render(gen.code, 0, data); err != nil {
		return err
	}

	if err := gen.convFuncBody(); err != nil {
		return err
	}
	gen.writeCode("}\n")
	return nil
}

// convFuncBody generates code for the numeric conversion function body.
func (gen *GenAnyToDst) convFuncBody() error {
	if !gen.dst.IsNumeric() {
		format := "unsupported conversion between any and %s"
		return fmt.Errorf(format, gen.dst.Code())
	}
	data := map[string]any{"dst": gen.dst}
	gen.addImport(cbAnyToDstFuncBody.Imports()...)
	if err := cbAnyToDstFuncBody.Render(gen.code, 1, data); err != nil {
		return err
	}
	return nil
}

// testFunc generates code for the conversion function tests.
func (gen *GenAnyToDst) testFunc() error {
	gen.addImport("testing", "github.com/ctx42/testing/pkg/assert")
	gen.writeCode("func ")
	gen.writeCode("Test_AnyTo%s_tabular", gen.dst.Title())
	gen.writeCode("(t *testing.T) {\n")
	if err := gen.testFuncBody(); err != nil {
		return err
	}
	gen.writeCode("}\n")
	return nil
}

// testFuncBody generates code for the conversion function tests.
func (gen *GenAnyToDst) testFuncBody() error {
	data := map[string]any{"dst": gen.dst}
	gen.addImport(cbTstAnyToDstTT.Imports()...)
	if err := cbTstAnyToDstTT.Render(gen.code, 1, data); err != nil {
		return err
	}
	if err := gen.testCases(); err != nil {
		return err
	}
	gen.writeCode("\t}\n\n")
	gen.addImport(cbTstAnyToDstLoop.Imports()...)
	return cbTstAnyToDstLoop.Render(gen.code, 1, data)
}

// testCases generates code for the conversion function test cases.
func (gen *GenAnyToDst) testCases() error {
	// Conversion success.
	tpl := cbTstAnyToDstFloatToInt
	if gen.dst.IsFloat() {
		tpl = cbTstAnyToDstIntToFloat
	}
	gen.addImport(tpl.Imports()...)
	if err := tpl.Render(gen.code, 2, nil); err != nil {
		return err
	}

	// Undefined conversion error.
	tpl = cbTstErrUndefinedConv
	data := map[string]any{"dst": gen.dst}
	gen.addImport(tpl.Imports()...)
	if err := tpl.Render(gen.code, 2, data); err != nil {
		return err
	}

	// Conversion error for unsigned types.
	if gen.dst.IsUnsigned() {
		tpl = cbTstErrInvalidRange
		data = map[string]any{
			"name":  "range",
			"src":   NumericType[int](),
			"dst":   gen.dst,
			"value": "-1",
		}
		gen.addImport(tpl.Imports()...)
		if err := tpl.Render(gen.code, 2, data); err != nil {
			return err
		}
	}

	// Conversion error for safe range overflow.
	if gen.dst.IsFloat() {
		mv := MaxSafeFloat(gen.dst.Size())
		tpl = cbTstErrOverSafeRange
		data = map[string]any{
			"src": NumericType[int](),
			"dst": gen.dst,
			"max": mv,
		}
		gen.addImport(mv.Imports()...)
		gen.addImport(tpl.Imports()...)
		if err := tpl.Render(gen.code, 2, data); err != nil {
			return err
		}
	}

	// Conversion error when destination must be a whole number.
	if gen.dst.IsInteger() {
		tpl = cbTstErrIsWhole
		data = map[string]any{"src": NumericType[float64](), "dst": gen.dst}
		gen.addImport(tpl.Imports()...)
		if err := tpl.Render(gen.code, 2, data); err != nil {
			return err
		}
	}

	return nil
}
