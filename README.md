[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/convert)](https://goreportcard.com/report/github.com/ctx42/convert)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/convert/pkg/convert)
[![Tests](https://github.com/ctx42/convert/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/ctx42/convert/actions/workflows/go.yml)

# convert

Lossless type conversions for Go: every conversion returns the exact value or
an error.

<!-- TOC -->
  * [Installation](#installation)
  * [Converters](#converters)
  * [Errors](#errors)
  * [Converter Registry](#converter-registry)
  * [Converter Types](#converter-types)
  * [AnyToXXX Converters](#anytoxxx-converters)
  * [Register Custom Converters](#register-custom-converters)
  * [Customize Converters](#customize-converters)
  * [Custom Registries](#custom-registries)
  * [License](#license)
<!-- TOC -->

**convert** is a lightweight Go library that performs safe type conversions
while preventing truncation, overflow, or unintended semantic changes between
values of different types.

## Installation

Requires Go 1.26 or later. Install using `go get`:

```bash
go get github.com/ctx42/convert
```

Import the package:

```go
import "github.com/ctx42/convert/pkg/convert"
```

## Converters

Use converter functions directly.

<!-- gmmce:pkg/convert/Example -->
```go
// Successful conversion.
ui8, err := convert.IntToUint8(42)
fmt.Printf("%[1]T(%[1]d), %v\n", ui8, err)

// Value too big for uint8.
ui8, err = convert.IntToUint8(420)
fmt.Printf("%[1]T(%[1]d), %v\n", ui8, err)

// Unsafe conversion.
f32, err := convert.IntToFloat32(convert.Float32SafeIntMax + 1)
fmt.Printf("%[1]T(%[1]g), %v\n", f32, err)

// Output:
// uint8(42), <nil>
// uint8(0), value out of range: from int to uint8
// float32(0), value out of safe range: from int to float32
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
- `uintptr`
- `time.Duration`

As well as converters implemented only between specific type pairs:

- `convert.BoolToBool`
- `convert.StringToDuration`
- `convert.StringToInt` through `convert.StringToUint64` - parse base-10
  strings exactly over the whole integer range.
- `convert.StringToFloat32`, `convert.StringToFloat64` - parse decimal
  strings; `StringToFloat32` rejects strings float32 cannot represent.
- `convert.StringToString`
- `convert.StringToTime(layout)` - registered by default with the
  `time.RFC3339Nano` layout.

All of these converters are automatically registered in the package-level
registry.

## Errors

A failed conversion returns a `convert.Error` wrapping one of the sentinel
errors, so match it with `errors.Is`:

| Error             | Returned when                                         |
|-------------------|-------------------------------------------------------|
| `ErrInvRange`     | Value outside the destination type range.             |
| `ErrInvSafeRange` | Value in range but would lose precision.              |
| `ErrFraction`     | Float with a fractional part converted to an integer. |
| `ErrInvValue`     | Invalid value, e.g. NaN, infinity, or a bad string.   |
| `ErrInvType`      | Value of the wrong type for the converter.            |
| `ErrUnkConv`      | No converter registered for the type pair.            |
| `ErrNilRegistry`  | Conversion with a nil registry.                       |

When parsing a string fails, the parser's error is kept as `Error.Cause` and is
matched by `errors.Is` and `errors.As` too.

<!-- gmmce:pkg/convert/Example_errors -->
```go
// Match the reason a conversion failed with errors.Is.
_, err := convert.IntToUint8(420)
fmt.Println(errors.Is(err, convert.ErrInvRange))

// A failed parse keeps the parser error as the cause.
var cnvErr convert.Error
_, err = convert.StringToDuration("abc")
if errors.As(err, &cnvErr) {
	fmt.Println(cnvErr.Cause)
}

// Output:
// true
// time: invalid duration "abc"
```

## Converter Registry

To get a converter function for a pair of types at runtime use `Lookup`
function.

<!-- gmmce:pkg/convert/ExampleLookup -->
```go
cnv := convert.Lookup[int, uint8]()

// Check cnv is not nil.

have, err := cnv(42)

// Check conversion error.

fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
// Output:
// output: uint8(42); error: <nil>
```

It returns a non-nil converter if the pair has been registered in the
package-level registry.

## Converter Types

All the converter functions provided by the package match the `SrcToDst` type.

```go
// SrcToDst represents a converter function that attempts lossless conversion
// of a value from the Src type to the Dst type. On success, it returns the
// converted value and a nil error. On failure (e.g., truncation, underflow,
// overflow, or semantic loss), it returns the zero value of Dst along with a
// non-nil error describing the issue.
type SrcToDst[Src, Dst any] func(Src) (Dst, error)

// AnyToAny is a non-generic version of [SrcToDst]. The behavior is exactly
// the same in terms of conversion and error handling.
type AnyToAny func(any) (any, error)
```

Additionally, the package defines the non-generic `AnyToAny` type. Converter
functions can be adapted to it with a helper.

<!-- gmmce:pkg/convert/ExampleToAnyAny -->
```go
cnv := convert.ToAnyAny(convert.Uint8ToUint8)

have, err := cnv("wrong")

fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
// Output:
// output: uint8(0); error: invalid type: expected uint8 got string
```

## AnyToXXX Converters

The package also provides a set of functions which convert a value of `any`
type to a given type:

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

<!-- gmmce:pkg/convert/ExampleRegister -->
```go
type A struct{ val int8 }
type B struct{ val int }

// Custom converter function matching [convert.SrcToDst] signature.
my := func(src A) (dst B, err error) {
	return B{val: int(src.val)}, nil
}

// Register a converter function between types A and B.
old := convert.Register(my)

// If there was already a converter for that source-destination type pair,
// it will be returned, nil otherwise.
_ = old

// Look up the registered converter.
cnv := convert.Lookup[A, B]()

// Run conversion.
have, err := cnv(A{42})

fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
// Output:
// output: convert_test.B({42}); error: <nil>
```

## Customize Converters

Some converters, like `convert.StringToTime`, are added to the package-level
registry with sane defaults, but you can customize them by overwriting the
default configuration.

<!-- gmmce:pkg/convert/ExampleRegister_overwrite -->
```go
// Replace the default string to time.Time converter.
def := convert.Register(convert.StringToTime(time.Kitchen))

// The default converter is returned in case you want to restore it.
defer convert.Register(def)

// Look up the registered converter.
cnv := convert.Lookup[string, time.Time]()

// Run conversion.
have, err := cnv("4:20AM")

fmt.Printf("output: %s; error: %v", have, err)
// Output:
// output: 0000-01-01 04:20:00 +0000 UTC; error: <nil>
```

## Custom Registries

Keep converters apart from the package-level registry by creating your own.
`RegisterConverter` and `LookupConverter` work on it, and the `AnyToXXX`
functions use it through the `WithRegistry` option.

<!-- gmmce:pkg/convert/ExampleNewRegistry -->
```go
reg := convert.NewRegistry()
convert.RegisterConverter(reg, convert.IntToUint8)

cnv := convert.LookupConverter[int, uint8](reg)
have, err := cnv(42)
fmt.Printf("output: %[1]T(%[1]d); error: %v\n", have, err)

v, err := convert.AnyToUint8(42, convert.WithRegistry(reg))
fmt.Printf("output: %[1]T(%[1]d); error: %v\n", v, err)

// Only converters registered in reg are available.
_, err = convert.AnyToUint8(int8(42), convert.WithRegistry(reg))
fmt.Println(err)

// Output:
// output: uint8(42); error: <nil>
// output: uint8(42); error: <nil>
// conversion undefined: from int8 to uint8
```

A nil registry makes `AnyToXXX` functions return `ErrNilRegistry`, and
`RegisterConverter` and `LookupConverter` return nil.

## License

MIT, see [LICENSE.md](LICENSE.md).
