// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xconv_test

import (
	"fmt"

	"github.com/ctx42/convert/pkg/xconv"
)

func ExampleLookup() {
	conv := xconv.Lookup[int, uint8]()

	// Chack conv is not nil.

	have, err := conv(42)

	// Check conversion error.

	fmt.Printf("output: %[1]T(%[1]d) error: %v", have, err)
	// Output:
	// output: uint8(42) error: <nil>
}

func ExampleLookup_error() {
	conv := xconv.Lookup[int, uint8]()

	// Chack conv is not nil.

	have, err := conv(-42)

	// Check conversion error.

	fmt.Printf("output: %[1]T(%[1]d) error: %v", have, err)
	// Output:
	// output: uint8(0) error: int value out of range for uint8
}

func ExampleRegister() {
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
}
