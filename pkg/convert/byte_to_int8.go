// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// ByteToInt8 converts a given numeric value to int8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func ByteToInt8(value byte) (int8, error) { return Uint8ToInt8(value) }
