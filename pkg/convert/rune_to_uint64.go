// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// RuneToUint64 converts a given numeric value to uint64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func RuneToUint64(value rune) (uint64, error) { return Int32ToUint64(value) }
