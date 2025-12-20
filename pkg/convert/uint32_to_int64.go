// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// Uint32ToInt64 converts a given numeric value to int64.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint32ToInt64(value uint32) (int64, error) { return int64(value), nil }
