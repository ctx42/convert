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
