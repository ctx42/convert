// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package codegen provides generators for conversion functions and their tests.
package codegen

// Number represents numeric types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~uintptr
}

// ActionName identifies an action to perform when converting between types.
type ActionName string

// List of action names.
const (
	// CastNotNeeded no cast or conversion is needed.
	CastNotNeeded ActionName = "cast_not_needed"

	// CastDirectly cast the source value directly to the destination.
	CastDirectly ActionName = "cast_directly"

	// CastToFloat64 perform direct type cast from source to the destination.
	CastToFloat64 ActionName = "cast_to_float64"

	// CheckIsNonNegative verify the source value is >= 0.
	CheckIsNonNegative ActionName = "check_is_non_negative"

	// CheckUnderflows verify the source minimum fits in the destination type.
	CheckUnderflows ActionName = "check_underflows"

	// CheckOverflows verify the source maximum fits in the destination type.
	CheckOverflows ActionName = "check_overflows"

	// CheckIsFinite verify the source float is finite (is not infinity).
	CheckIsFinite ActionName = "check_is_finite"

	// CheckIsNumber verify the source float is a number (is not NaN).
	CheckIsNumber ActionName = "check_is_number"

	// CheckIsWhole verify the source float has no fractional part.
	CheckIsWhole ActionName = "check_is_whole"

	// CheckIntSafeToFloatMin verify the source integer minimum is
	// representable as float without loss of precision.
	CheckIntSafeToFloatMin ActionName = "check_int_safe_to_float_min"

	// CheckIntSafeToFloatMax verify the source integer maximum is
	// representable as float without loss of precision.
	CheckIntSafeToFloatMax ActionName = "check_int_safe_to_float_max"

	// CheckFloatSafeToIntMin verify the source float is representable as
	// integer without loss of precision.
	CheckFloatSafeToIntMin ActionName = "check_safe_float_int_min"

	// CheckFloatSafeToIntMax verify the source float is representable as
	// integer without loss of precision.
	CheckFloatSafeToIntMax ActionName = "check_safe_float_int_max"
)

// Action defines an action to take during conversion between types.
type Action struct {
	name  ActionName // The action.
	value *Value     // Optional [Value] associated with the action.
}

// NewAction returns a new [Action] instance with an optional value.
func NewAction(name ActionName, value *Value) Action {
	return Action{name: name, value: value}
}

// Name returns the action name.
func (act Action) Name() ActionName { return act.name }

// Imports returns the Go package imports required by the value. Returns nil if
// the value is not set or requires no Go package imports.
func (act Action) Imports() []string { return act.value.Imports() }

// Code returns the Go code representation of the value. Returns an empty
// string if the value is not set.
func (act Action) Code() string { return act.value.Code() }
