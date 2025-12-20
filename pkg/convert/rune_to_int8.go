// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// RuneToInt8 converts a given numeric value to int8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func RuneToInt8(value rune) (int8, error) { return Int32ToInt8(value) }
