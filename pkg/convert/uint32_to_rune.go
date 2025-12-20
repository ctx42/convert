// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// Uint32ToRune converts a given numeric value to rune.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint32ToRune(value uint32) (rune, error) { return Uint32ToInt32(value) }
