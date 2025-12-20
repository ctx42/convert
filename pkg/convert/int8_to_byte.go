// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// Int8ToByte converts a given numeric value to byte.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int8ToByte(value int8) (byte, error) { return Int8ToUint8(value) }
