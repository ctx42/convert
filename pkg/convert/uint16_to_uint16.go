// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// Uint16ToUint16 converts a given numeric value to uint16.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint16ToUint16(value uint16) (uint16, error) { return value, nil }
