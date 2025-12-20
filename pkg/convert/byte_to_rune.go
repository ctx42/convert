// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// ByteToRune converts a given numeric value to rune.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func ByteToRune(value byte) (rune, error) { return Uint8ToInt32(value) }
