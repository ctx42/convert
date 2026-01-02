[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/convert)](https://goreportcard.com/report/github.com/ctx42/convert)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/convert)
![Tests](https://github.com/ctx42/convert/actions/workflows/go.yml/badge.svg?branch=master)

<!-- TOC -->
  * [Installation](#installation)
  * [Converters](#converters)
  * [Converter Registry](#converter-registry)
  * [Converter Types](#converter-types)
  * [AnyToXXX Converters.](#anytoxxx-converters)
  * [Register Custom Converters](#register-custom-converters)
  * [Customize Converters](#customize-converters)
  * [32bit vs. 64bit Systems](#32bit-vs-64bit-systems)
<!-- TOC -->

**convert** is a lightweight Go library for safe type conversions, preventing
truncation, overflow, or semantic loss when converting between types. 

It has support for 32-bit and 64-bit systems when converting numeric types.   

## Installation

Install using `go get`:

```bash
go get github.com/ctx42/convert
```

## Converters

Use converter functions directly.

```go
// Successful conversion.
ui8, err := convert.IntToUint8(42)
fmt.Printf("convert.IntToUint8 output: %[1]T(%[1]d) error: %v\n", ui8, err)

// Value too big for uint8.
ui8, err = convert.IntToUint8(420)
fmt.Printf("convert.IntToUint8 output: %[1]T(%[1]d) error: %v\n", ui8, err)

// Unsafe conversion.
f32, err := convert.IntToFloat32(convert.Float32SafeIntMax + 1)
fmt.Printf("convert.IntToUint8 output: %[1]T(%[1]g) error: %v\n", f32, err)

// Output:
// convert.IntToUint8 output: uint8(42) error: <nil>
// convert.IntToUint8 output: uint8(0) error: int value out of range for uint8
// convert.IntToUint8 output: float32(0) error: int value out of safe range for float32
```

Package `convert` provides more than 200 converter functions between numeric
types:

- `uint`
- `uint8`
- `uint16`
- `uint32`
- `uint64`
- `int`
- `int8`
- `int16`
- `int32`
- `int64`
- `float32`
- `float64`
- `byte`
- `rune`
- `uintqptr`
- `time.Duration`

As well as converters implemented only between specific type pairs:

- `convert.StringToDuration`
- `convert.StringToString`
- `convert.StringToTime` - string must be in `time.RFC3339Nano` format.

All of them pairs are automatically registered in the package-level registry.

## Converter Registry

To get a converter function for a pair of types at runtinme use `Lookup`
function.

```go
cnv := convert.Lookup[int, uint8]()

// Chack cnv is not nil.

have, err := cnv(42)

// Check conversion error.

fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
// Output:
// output: uint8(42); error: <nil>
```

It returns a non-nil converter if the pair has been registered in the 
package-level registry.

## Converter Types

All the converter functions provided by the package match generic `FromTo` type. 

```go
// FromTo represents a converter function that attempts lossless conversion of
// a value from the type "From" to the "To" type. On success, it returns the
// converted value and a nil error. On failure (e.g., truncation, underflow,
// overflow, or semantic loss), it returns the zero value of "To" along with a
// non-nil error describing the issue.
type FromTo[From, To any] func(from From) (to To, err error)

// AnyToAny is a non-generic version of [FromTo]. The behavior is exactly the
// same in terms of error handling.
type AnyToAny func(form any) (to any, err error)
```

Additionally, package defines non-generic `AnyToAny` type. Converter functions
can be adapted to it with a helper.

```go
var cnv func(any) (any, error)

cnv = convert.ToAnyAny(convert.Uint8ToUint8)

have, err := cnv("wrong")

fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
// Output:
// output: uint8(0); error: invalid type: expected uint8, got string
```

## AnyToXXX Converters.

Module also provides a set of functions which can convert `any` type to a given
type:

- `AnyToByte`
- `AnyToDuration`
- `AnyToFloat32`
- `AnyToFloat64`
- `AnyToInt`
- `AnyToInt8`
- `AnyToInt16`
- `AnyToInt32`
- `AnyToInt64`
- `AnyToUint`
- `AnyToUint8`
- `AnyToUint16`
- `AnyToUint32`
- `AnyToUint64`
- `AnyToUintptr`
- `AnyToRune`

## Register Custom Converters

```go
type A struct{ val int8 }
type B struct{ val int }

// Custom converter function matching [convert.Converter] signature.
my := func(from A) (to B, err error) {
    return B{val: int(from.val)}, nil
}

// Register a converter function between types A and B.
old := convert.Register(my)

// If there was already a converter for that source-destination type pair,
// it will be returned, nil otherwise.
_ = old

// Lookup converter registered converter.
cnv := convert.Lookup[A, B]()

// Run conversion.
have, err := cnv(A{42})

fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
// Output:
// output: convert_test.B({42}); error: <nil>
```

## Customize Converters

Some converters, like `convert.StringToTime`, are added to package-level
registry with sane defaults but you can customize them by overwriting the
defeult configuration.

```go
// Register a converter function between types A and B.
def := convert.Register(convert.StringToTime(time.Kitchen))

// The default converter is returned in case you want to restore it.
defer convert.Register(def)

// Lookup converter registered converter.
cnv := convert.Lookup[string, time.Time]()

// Run conversion.
have, err := cnv("4:20AM")

fmt.Printf("output: %s; error: %v", have, err)
// Output:
// output: 0000-01-01 04:20:00 +0000 UTC; error: <nil>
```

## 32bit vs. 64bit Systems

In cases where its necessary module implements separate boundary checks for
32-bit and 64-bit systems. See files in [convert](pkg/convert) directory with
`_32bit` and `_64bit` strings in their filenames.
