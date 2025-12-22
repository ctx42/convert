// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// AnyToByte converts the given value to byte using the package-level
// registry.
func AnyToByte(value any) (byte, error) {
	return AnyToByteUsing(registry, value)
}

// AnyToByteUsing converts the given value to byte using the provided
// registry.
func AnyToByteUsing(reg *Registry, value any) (byte, error) {
	return AnyToUint8Using(reg, value)
}
