// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

import (
	"unicode"
)

// Value represents Go code which can be rendered as a string.
type Value struct {
	prefix string // Prefix to put before the rendered value.
	pkg    string // Optional package name for the value.
	value  string // The value to be rendered.
}

// NewValue creates a new [Value] instance.
//
// Examples:
// - `NewValue("ConstantName")`
// - `NewValue("math", "MinInt32")`
// - `NewValue("-", "math", "MaxFloat32")`
func NewValue(arg string, args ...string) *Value {
	switch len(args) {
	case 0:
		return &Value{value: arg}
	case 1:
		return &Value{pkg: arg, value: args[0]}
	case 2:
		return &Value{prefix: arg, pkg: args[0], value: args[1]}
	default:
		panic("invalid number of arguments for NewValue")
	}
}

// Name returns the value's name, optionally titled.
func (val *Value) Name(titled bool) string {
	if val == nil {
		return ""
	}
	if titled && val.value != "" {
		runes := []rune(val.value)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}
	return val.value
}

// Imports returns the list of go imports required by the value. Returns nil if
// the value is nil or has an empty package field.
func (val *Value) Imports() []string {
	if val == nil || val.pkg == "" {
		return nil
	}
	return []string{val.pkg}
}

// Code returns the rendered value as a string.
func (val *Value) Code() string {
	if val == nil {
		return ""
	}
	ret := val.prefix
	if val.pkg != "" {
		ret += val.pkg + "."
	}
	ret += val.value
	return ret
}
