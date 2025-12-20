// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert_test

import (
	"fmt"
	"time"

	"github.com/ctx42/convert/pkg/convert"
)

func Example() {
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
}

func ExampleLookup() {
	cnv := convert.Lookup[int, uint8]()

	// Chack cnv is not nil.

	have, err := cnv(42)

	// Check conversion error.

	fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
	// Output:
	// output: uint8(42); error: <nil>
}

func ExampleLookup_time() {
	cnv := convert.Lookup[string, time.Time]()

	// Chack cnv is not nil.

	have, err := cnv("2000-01-02T03:04:05Z")

	// Check conversion error.

	fmt.Printf("output: %[1]s; error: %v", have, err)
	// Output:
	// output: 2000-01-02 03:04:05 +0000 UTC; error: <nil>
}

func ExampleLookup_error() {
	cnv := convert.Lookup[int, uint8]()

	// Chack conv is not nil.

	have, err := cnv(-42)

	// Check conversion error.

	fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
	// Output:
	// output: uint8(0); error: int value out of range for uint8
}

func ExampleRegister() {
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
}

func ExampleRegister_overwrite() {
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
}

func ExampleToAnyAny() {
	var cnv func(any) (any, error)

	cnv = convert.ToAnyAny(convert.Uint8ToUint8)

	have, err := cnv("wrong")

	fmt.Printf("output: %[1]T(%[1]d); error: %v", have, err)
	// Output:
	// output: uint8(0); error: invalid type: expected uint8, got string
}
