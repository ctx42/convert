// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

//go:build amd64 || arm64 || mips64 || mips64le

package convert

// Uint64ToUintptr converts a given numeric value to uintptr.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Uint64ToUintptr(value uint64) (uintptr, error) {
	return uintptr(value), nil
}
