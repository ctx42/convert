[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/convert)](https://goreportcard.com/report/github.com/ctx42/convert)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/convert)
![Tests](https://github.com/ctx42/convert/actions/workflows/go.yml/badge.svg?branch=master)

<!-- TOC -->
  * [Installation](#installation)
  * [Example Usage](#example-usage)
  * [Introduction](#introduction)
  * [Built-in Converters](#built-in-converters)
  * [Dynamically Lookup Converters](#dynamically-lookup-converters)
  * [Register Custom Converters](#register-custom-converters)
  * [32bit vs. 64bit Systems](#32bit-vs-64bit-systems)
<!-- TOC -->

**convert** is a lightweight Go library for safe type conversions, preventing
truncation, overflow, or semantic loss from invalid casts. It supports 
conversions between common numeric types out of the box.

It's safe to use on 32-bit and 64-bit systems.   

## Installation

Install via `go get`:

```bash
go get github.com/ctx42/convert
```

## Example Usage

```go
// Successful cast.
ui8, err := xcast.IntToUint8(42)
fmt.Printf("xcast.IntToUint8 output: %[1]T(%[1]d) error: %v\n", ui8, err)

// Value too big for uint8.
ui8, err = xcast.IntToUint8(420)
fmt.Printf("xcast.IntToUint8 output: %[1]T(%[1]d) error: %v\n", ui8, err)

// Unsafe cast.
f32, err := xcast.IntToFloat32(xcast.Float32SafeIntMax + 1)
fmt.Printf("xcast.IntToUint8 output: %[1]T(%[1]g) error: %v\n", f32, err)

// Output:
// xcast.IntToUint8 output: uint8(42) error: <nil>
// xcast.IntToUint8 output: uint8(0) error: int value out of range for uint8
// xcast.IntToUint8 output: float32(0) error: int value out of safe range for float32
```

## Introduction

The main type used in the module is defined in the `xconv` package:

```go
// Converter represents a function that attempts lossless conversion from a
// source value of type From to a target value of type To. On success, it 
// returns the converted value and a nil error. On failure (e.g., truncation,
// overflow, or semantic loss), it returns the zero value of To along with a
// non-nil error describing the issue.
type Converter[From, To any] func(from From) (to To, err error)
```

All build in converters match the `Converter`.

Module is split into two packages:

- `xcast` - converter functions like `Float32ToInt8`.
- `xconvert` - registry of converters with the ability to register your own.

## Built-in Converters

Package `xcast` provides 225 `Converter` functions between build-in types:

- byte
- uint8
- uint16
- uint32
- uint64
- uint
- int8
- int16
- int32
- int64
- int
- float32
- float64
- rune
- uintqptr

## Dynamically Lookup Converters

Use `xconv` package to look up / register converts during runtime.

```go
conv := xconv.Lookup[int, uint8]()

// Chack conv is not nil.

have, err := conv(42)

// Check conversion error.

fmt.Printf("output: %[1]T(%[1]d) error: %v", have, err)
// Output:
// output: uint8(42) error: <nil>
```

## Register Custom Converters

```go
type A struct{ val int8 }
type B struct{ val int }

// Custom converter function matching [xconv.Converter] signature.
myConv := func(from A) (to B, err error) {
    return B{val: int(from.val)}, nil
}

// Register a converter function between types A and B.
old := xconv.Register(myConv)

// If there was already a converter for that source-destination type pair,
// it will be returned, nil otherwise.
_ = old

// Lookup converter registered converter.
conv := xconv.Lookup[A, B]()

// Run conversion.
have, err := conv(A{42})

fmt.Printf("output: %[1]T(%[1]d) error: %v", have, err)
// Output:
// output: xconv_test.B({42}) error: <nil>
```

## 32bit vs. 64bit Systems

In cases where its necessary module implements separate boundary checks for
32-bit dnd 64-bit systems. See files in [xcast](pkg/xcast) directory with
`_32bit` and `_64bit` strings in their fiolenames.