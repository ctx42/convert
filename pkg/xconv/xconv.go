// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package xconv provides utilities for lossless type conversions.
package xconv

// registry represents package level [Registry].
var registry = &Registry{}

// Register adds the provided [Converter] to the package-level [Registry].
// If a converter for the same source-destination type pair already exists,
// it is replaced, and the previous converter is returned; otherwise nil is
// returned.
func Register[From, To any](conv Converter[From, To]) Converter[From, To] {
	return RegisterConverter(registry, conv)
}

// Lookup returns the [Converter] for the given source-destination type pair
// from the package-level [Registry]. Returns nil if no converter was
// registered for the given source-destination type pair.
func Lookup[From, To any]() Converter[From, To] {
	return LookupConverter[From, To](registry)
}

// Converter represents a function that attempts lossless conversion from a
// source value of type From to a target value of type To.
// On success, it returns the converted value and a nil error.
// On failure (e.g., truncation, overflow, or semantic loss), it returns the
// zero value of To along with a non-nil error describing the issue.
type Converter[From, To any] func(from From) (to To, err error)
