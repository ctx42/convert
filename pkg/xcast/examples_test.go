// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xcast_test

import (
	"fmt"

	"github.com/ctx42/convert/pkg/xcast"
)

func Example() {
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
}
