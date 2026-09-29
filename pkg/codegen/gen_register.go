// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"errors"
	"io"
)

// GenRegister represents the code generator for the function registering
// converters between values of the given types in the package-level registry.
type GenRegister struct {
	file // Represents a file with source code.
}

// NewGenRegister returns a new [GenRegister] instance.
func NewGenRegister(pkg string, ops ...Option) *GenRegister {
	return &GenRegister{file: newFile(pkg, ops...)}
}

// GenerateCode generates source code for the registerNumeric function, which
// registers the converter for every source-destination pair of the given
// types. Returns an error when no types are given.
func (gen *GenRegister) GenerateCode(w io.Writer, types []Type) error {
	if len(types) == 0 {
		return errors.New("no types to register converters for")
	}
	gen.reset()

	var names []string
	for _, src := range types {
		for _, dst := range types {
			names = append(names, src.Title()+"To"+dst.Title())
		}
	}
	data := map[string]any{"names": names}
	gen.addImport(cbRegisterFunc.Imports()...)
	if err := cbRegisterFunc.Render(gen.code, 0, data); err != nil {
		return err
	}
	if _, err := gen.WriteTo(w); err != nil {
		return err
	}
	return nil
}
