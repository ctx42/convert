// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

// Int8ToFloat32 converts a given numeric value to float32.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int8ToFloat32(value int8) (float32, error) { return float32(value), nil }
