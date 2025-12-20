// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// ByteToUint8 converts a given numeric value to uint8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func ByteToUint8(value byte) (uint8, error) { return Uint8ToUint8(value) }
