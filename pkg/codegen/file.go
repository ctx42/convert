// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"bytes"
	"fmt"
	"io"
	"slices"
	"strings"
)

// file represents a file with source code.
type file struct {
	ops  Options       // Generation options.
	pkg  string        // Package name.
	imps []string      // Go package imports required by the code.
	code *bytes.Buffer // Source code.
}

// newFile returns a new instance of the file.
func newFile(pkg string, opts ...Option) file {
	return file{
		pkg:  pkg,
		code: &bytes.Buffer{},
		imps: make([]string, 0, 20),
		ops:  NewOptions(opts...),
	}
}

// writeCode writes formatted string to the code buffer.
func (fil *file) writeCode(format string, args ...any) {
	_, _ = fmt.Fprintf(fil.code, format, args...)
}

// addImport registers Go import path(s) required by the source code.
func (fil *file) addImport(imps ...string) {
	fil.imps = append(fil.imps, imps...)
}

// renderImports returns the formatted Go imports code block with sorted,
// de-duplicated, non-empty import paths. Returns an empty string when there
// are no import paths to render.
func (fil *file) renderImports() string {
	imps := slices.DeleteFunc(slices.Clone(fil.imps), func(imp string) bool {
		return imp == ""
	})
	if len(imps) == 0 {
		return ""
	}
	slices.Sort(imps)
	imps = slices.Compact(imps)

	buf := &strings.Builder{}
	buf.WriteString("import (\n")
	for _, imp := range imps {
		buf.WriteString("\t\"")
		buf.WriteString(imp)
		buf.WriteString("\"\n")
	}
	buf.WriteString(")\n\n")
	return buf.String()
}

// tabularTest writes a tabular test function with the given name: an optional
// header code block, the table declaration, the table rows and the loop over
// them, all rendered with the same data. A test function written after other
// code is separated from it with an empty line.
func (fil *file) tabularTest(
	name string,
	header, table, loop *CodeBlock,
	rows string,
	data map[string]any,
) error {

	if fil.code.Len() > 0 {
		fil.writeCode("\n")
	}
	fil.writeCode("func %s(t *testing.T) {\n", name)
	if header != nil {
		fil.addImport(header.Imports()...)
		if err := header.Render(fil.code, 1, data); err != nil {
			return err
		}
		fil.writeCode("\n")
	}
	fil.addImport(table.Imports()...)
	if err := table.Render(fil.code, 1, data); err != nil {
		return err
	}
	fil.code.WriteString(rows)
	fil.writeCode("\t}\n\n")
	fil.addImport(loop.Imports()...)
	if err := loop.Render(fil.code, 1, data); err != nil {
		return err
	}
	fil.writeCode("}\n")
	return nil
}

// reset resets code buffer and imports slice.
func (fil *file) reset() {
	fil.code.Reset()
	fil.imps = fil.imps[:0]
}

// WriteTo writes formatted file content to the provided writer.
func (fil *file) WriteTo(w io.Writer) (int64, error) {
	buf := &bytes.Buffer{}
	buf.WriteString(fil.ops.copyright)
	buf.WriteString(fil.ops.generatedBy)
	buf.WriteString("package ")
	buf.WriteString(fil.pkg)
	buf.WriteString("\n\n")
	buf.WriteString(fil.renderImports())
	buf.WriteString(fil.code.String())
	return buf.WriteTo(w)
}
