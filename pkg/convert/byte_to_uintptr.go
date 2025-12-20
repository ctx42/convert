// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// ByteToUintptr converts a given numeric value to uintptr.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func ByteToUintptr(value byte) (uintptr, error) { return Uint8ToUintptr(value) }
