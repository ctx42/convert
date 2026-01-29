// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package codegen

// Int64 is an example signed type with a [Number] as a base type.
type Int64 int64

// Uint64 is an example unsigned type with a [Number] as a base type.
type Uint64 uint64

// Float64 is an example unsigned type with a [Number] as a base type.
type Float64 float64

// IntSize represents the size of the int on the current platform (32 or 64).
const IntSize = 32 << (^uint(0) >> 63)
