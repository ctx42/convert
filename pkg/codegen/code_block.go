// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"text/template"
)

// CodeBlock represents a Go code block.
type CodeBlock struct {
	imps []string           // Go package imports required by the code block.
	tpl  *template.Template // Parsed text template.
}

// MustCodeBlock works the same way as [NewCodeBlock] but panics on error.
func MustCodeBlock(name, tpl string) *CodeBlock {
	cb, err := NewCodeBlock(name, tpl)
	if err != nil {
		panic(err)
	}
	return cb
}

// NewCodeBlock creates a new instance of [CodeBlock] with the given name from
// a Go text template string.
func NewCodeBlock(name, tpl string) (*CodeBlock, error) {
	cb := &CodeBlock{tpl: template.New(name).Option("missingkey=error")}
	scn := bufio.NewScanner(strings.NewReader(tpl))
	var code []string
	for scn.Scan() {
		line, found := strings.CutPrefix(scn.Text(), "import ")
		if found {
			cb.imps = append(cb.imps, line)
			continue
		}
		code = append(code, line)
	}
	code = append(code, "")

	tpl = strings.Join(code, "\n")
	tpl = strings.TrimLeft(tpl, " \t\n\r")

	var err error
	if cb.tpl, err = cb.tpl.Parse(tpl); err != nil {
		return nil, err
	}
	return cb, nil
}

// Imports returns the Go package imports required by the text template.
func (cb *CodeBlock) Imports() []string { return cb.imps }

// Render writes the rendered code block with given data to the writer, giving
// each line the specified tab indent.
func (cb *CodeBlock) Render(w io.Writer, ind int, data map[string]any) error {
	rendered := &strings.Builder{}
	if err := cb.tpl.Execute(rendered, data); err != nil {
		return err
	}

	indented := &bytes.Buffer{}
	scn := bufio.NewScanner(strings.NewReader(rendered.String()))
	for scn.Scan() {
		line := scn.Text()
		if line != "" {
			indented.WriteString(strings.Repeat("\t", ind))
			indented.WriteString(line)
		}
		indented.WriteString("\n")
	}
	if _, err := indented.WriteTo(w); err != nil {
		return err
	}
	return nil
}
