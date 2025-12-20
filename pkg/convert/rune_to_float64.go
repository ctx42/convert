// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// RuneToFloat64 converts a given numeric value to float64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func RuneToFloat64(value rune) (float64, error) { return Int32ToFloat64(value) }
