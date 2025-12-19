// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv

import (
	"reflect"
)

// wrapper encapsulates a concrete [Converter] function for a specific
// source-destination type pair.
type wrapper struct {
	from, to reflect.Type // Source and destination types.
	cnv      any          // A [Converter] matching from and to types.
	cst      Caster       // A non-generic [Converter].
}

// wrap creates a new wrapper instance for a given [Converter] function.
func wrap[From, To any](conv Converter[From, To]) *wrapper {
	return &wrapper{
		from: reflect.TypeFor[From](),
		to:   reflect.TypeFor[To](),
		cnv:  conv,
		cst:  ConverterToCaster(conv),
	}
}
