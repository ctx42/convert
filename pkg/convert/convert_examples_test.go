// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/ctx42/convert/pkg/convert"
)

func Example() {
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
}

func Example_errors() {
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
}

func ExampleLookup() {
	cnv := convert.Lookup[int, uint8]()

	// Check cnv is not nil.

	have, err := cnv(42)

	// Check conversion error.

	fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
	// Output:
	// output: uint8(42); error: <nil>
}

func ExampleLookup_time() {
	cnv := convert.Lookup[string, time.Time]()

	// Check cnv is not nil.

	have, err := cnv("2000-01-02T03:04:05Z")

	// Check conversion error.

	fmt.Printf("output: %[1]s; error: %v", have, err)
	// Output:
	// output: 2000-01-02 03:04:05 +0000 UTC; error: <nil>
}

func ExampleLookup_error() {
	cnv := convert.Lookup[int, uint8]()

	// Check cnv is not nil.

	have, err := cnv(-42)

	// Check conversion error.

	fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
	// Output:
	// output: uint8(0); error: value out of range: from int to uint8
}

func ExampleRegister() {
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
}

func ExampleRegister_overwrite() {
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
}

func ExampleNewRegistry() {
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
}

func ExampleToAnyAny() {
	cnv := convert.ToAnyAny(convert.Uint8ToUint8)

	have, err := cnv("wrong")

	fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
	// Output:
	// output: uint8(0); error: invalid type: expected uint8 got string
}
