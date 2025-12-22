// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// AnyToRune converts the given value to rune using the package-level
// registry.
func AnyToRune(value any) (rune, error) {
	return AnyToRuneUsing(registry, value)
}

// AnyToRuneUsing converts the given value to rune using the provided
// registry.
func AnyToRuneUsing(reg *Registry, value any) (rune, error) {
	return AnyToInt32Using(reg, value)
}
