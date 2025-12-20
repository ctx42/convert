// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"reflect"
)

// wrapper encapsulates a generic [FromTo] converter.
type wrapper struct {
	from, to reflect.Type // Source and destination types.
	cnv      any          // A generic converter.
	cst      AnyToAny     // A non-generic converter.
}

// wrap creates a new wrapper instance for a given [FromTo] function.
func wrap[From, To any](conv FromTo[From, To]) *wrapper {
	return &wrapper{
		from: reflect.TypeFor[From](),
		to:   reflect.TypeFor[To](),
		cnv:  conv,
		cst:  ToAnyAny(conv),
	}
}
