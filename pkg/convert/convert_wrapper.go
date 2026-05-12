// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"reflect"
)

// wrapper encapsulates a generic [SrcToDst] converter.
type wrapper struct {
	src, dst reflect.Type // Source and destination types.
	cnv      any          // A generic converter.
	cst      AnyToAny     // A non-generic converter.
}

// wrap creates a new wrapper instance for a given [SrcToDst] function.
func wrap[Src, Dst any](conv SrcToDst[Src, Dst]) *wrapper {
	return &wrapper{
		src: reflect.TypeFor[Src](),
		dst: reflect.TypeFor[Dst](),
		cnv: conv,
		cst: ToAnyAny(conv),
	}
}
