// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package codegen

import (
	"fmt"
	"reflect"
	"strings"
)

// IsSigned checks if the provided [Number] is a signed type.
func IsSigned[T Number]() bool {
	typ := reflect.TypeFor[T]()
	switch typ.Kind() {
	case reflect.Uint:
		return false

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return false

	case reflect.Uintptr:
		return false

	case reflect.Int:
		return true

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true

	case reflect.Float32, reflect.Float64:
		return true

	default:
		return false
	}
}

// IsFloat checks if the provided [Number] is a floating-point type.
func IsFloat[T Number]() bool {
	typ := reflect.TypeFor[T]()
	switch typ.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// article returns the English indefinite article for the given word: "an"
// before a word starting with a, e, i or o, and "a" otherwise. Go type names
// starting with u, like uint, are pronounced with a consonant sound.
func article(word string) string {
	if word != "" && strings.ContainsRune("aeioAEIO", rune(word[0])) {
		return "an"
	}
	return "a"
}

// comment returns the text as Go line comments, word-wrapped so no line,
// including the "// " prefix, is longer than width. A word longer than the
// width is put on its own line.
func comment(text string, width int) string {
	var lines []string
	line := "//"
	for _, word := range strings.Fields(text) {
		if line != "//" && len(line)+1+len(word) > width {
			lines = append(lines, line)
			line = "//"
		}
		line += " " + word
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}

// MinValue returns a [Value] representing the minimum value for the given type.
func MinValue(typ Type) *Value {
	if typ.IsPlatform() {
		return MinPlatformInteger(typ)
	}
	if typ.IsInteger() {
		return MinInteger(typ.size, typ.IsSigned())
	}
	return MinSafeFloat(typ.Size())
}

// MaxValue returns a [Value] representing the maximum value for the given type.
func MaxValue(typ Type) *Value {
	if typ.IsPlatform() {
		return MaxPlatformInteger(typ)
	}
	if typ.IsInteger() {
		return MaxInteger(typ.size, typ.IsSigned())
	}
	return MaxSafeFloat(typ.Size())
}

// MinInteger returns a [Value] representing a minimum integer value of the
// given size.
func MinInteger(size int, signed bool) *Value {
	if signed {
		return NewValue("math", fmt.Sprintf("MinInt%d", size))
	}
	return NewValue("0")
}

// MaxInteger returns a [Value] representing a maximum integer value of the
// given size.
func MaxInteger(size int, signed bool) *Value {
	if signed {
		return NewValue("math", fmt.Sprintf("MaxInt%d", size))
	}
	return NewValue("math", fmt.Sprintf("MaxUint%d", size))
}

// MinPlatformInteger returns a [Value] representing the minimum value of the
// given platform-sized integer type, valid on both 32-bit and 64-bit
// platforms.
func MinPlatformInteger(typ Type) *Value {
	if typ.IsSigned() {
		return NewValue("math", "MinInt")
	}
	return NewValue("0")
}

// MaxPlatformInteger returns a [Value] representing the maximum value of the
// given platform-sized integer type, valid on both 32-bit and 64-bit
// platforms. For uintptr it is
// [github.com/ctx42/convert/pkg/convert.MaxUintptr].
func MaxPlatformInteger(typ Type) *Value {
	switch {
	case typ.Name() == "uintptr":
		return NewValue("MaxUintptr")
	case typ.IsSigned():
		return NewValue("math", "MaxInt")
	default:
		return NewValue("math", "MaxUint")
	}
}

// MinSafeFloat returns a [Value] representing a minimum safe integer value for
// the given floating-point size.
//
// See:
//   - [github.com/ctx42/convert/pkg/convert.Float32SafeIntMin]
//   - [github.com/ctx42/convert/pkg/convert.Float64SafeIntMin]
func MinSafeFloat(size int) *Value {
	return NewValue(fmt.Sprintf("Float%dSafeIntMin", size))
}

// MaxSafeFloat returns a [Value] representing a maximum safe integer value for
// the given floating-point size.
//
// See:
//   - [github.com/ctx42/convert/pkg/convert.Float32SafeIntMax]
//   - [github.com/ctx42/convert/pkg/convert.Float64SafeIntMax]
func MaxSafeFloat(size int) *Value {
	return NewValue(fmt.Sprintf("Float%dSafeIntMax", size))
}
