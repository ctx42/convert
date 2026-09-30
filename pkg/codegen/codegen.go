// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package codegen provides generators for conversion functions and their tests.
package codegen

import (
	"slices"
)

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
	// CastNotNeeded means no cast or conversion is needed.
	CastNotNeeded ActionName = "cast_not_needed"

	// CastDirectly casts the source value directly to the destination type.
	CastDirectly ActionName = "cast_directly"

	// CastToFloat64 converts the source value to a float64 local variable.
	CastToFloat64 ActionName = "cast_to_float64"

	// CheckIsNonNegative checks the source value is not negative.
	CheckIsNonNegative ActionName = "check_is_non_negative"

	// CheckUnderflows checks the source value is not below the destination type
	// minimum.
	CheckUnderflows ActionName = "check_underflows"

	// CheckOverflows checks the source value is not above the destination type
	// maximum.
	CheckOverflows ActionName = "check_overflows"

	// CheckIsFinite checks the source float is finite (is not infinity).
	CheckIsFinite ActionName = "check_is_finite"

	// CheckIsNumber checks the source float is a number (is not NaN).
	CheckIsNumber ActionName = "check_is_number"

	// CheckIsWhole checks the source float has no fractional part.
	CheckIsWhole ActionName = "check_is_whole"

	// CheckIntSafeToFloatMin checks the source integer is not below the
	// smallest integer the destination float represents exactly.
	CheckIntSafeToFloatMin ActionName = "check_int_safe_to_float_min"

	// CheckIntSafeToFloatMax checks the source integer is not above the
	// largest integer the destination float represents exactly.
	CheckIntSafeToFloatMax ActionName = "check_int_safe_to_float_max"

	// CheckFloatSafeToIntMin checks the source float is not below the
	// smallest integer its floating-point type represents exactly.
	CheckFloatSafeToIntMin ActionName = "check_safe_float_int_min"

	// CheckFloatSafeToIntMax checks the source float is not above the
	// largest integer its floating-point type represents exactly.
	CheckFloatSafeToIntMax ActionName = "check_safe_float_int_max"

	// CheckFloatRange checks the finite source float is within the range of
	// the destination floating-point type.
	CheckFloatRange ActionName = "check_float_range"

	// CheckFloatExact checks the source float is exactly representable by
	// the destination floating-point type.
	CheckFloatExact ActionName = "check_float_exact"
)

// actionOrder lists action names in the order the generated code performs
// them. Range checks precede safe range checks, so when actions derived for
// 32-bit and 64-bit platforms are merged, a value outside the destination
// range on the current platform fails the range check first.
var actionOrder = []ActionName{
	CastToFloat64,
	CheckIsNumber,
	CheckIsFinite,
	CheckIsWhole,
	CheckFloatRange,
	CheckFloatExact,
	CheckIsNonNegative,
	CheckUnderflows,
	CheckOverflows,
	CheckIntSafeToFloatMin,
	CheckIntSafeToFloatMax,
	CheckFloatSafeToIntMin,
	CheckFloatSafeToIntMax,
	CastNotNeeded,
	CastDirectly,
}

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

// equal returns true if both actions have the same name and value code.
func (act Action) equal(other Action) bool {
	return act.name == other.name && act.Code() == other.Code()
}

// mergeActions merges action lists derived for 32-bit and 64-bit platforms
// into one list valid on both, ordered by actionOrder. Range checks against a
// platform-sized target use its platform-dependent limits.
func mergeActions(target Type, lists ...[]Action) []Action {
	var merged []Action
	for _, actions := range lists {
		for _, act := range actions {
			if target.IsPlatform() {
				switch act.name {
				case CheckUnderflows:
					act.value = MinPlatformInteger(target)
				case CheckOverflows:
					act.value = MaxPlatformInteger(target)
				}
			}
			sameName := func(m Action) bool { return m.name == act.name }
			if !slices.ContainsFunc(merged, sameName) {
				merged = append(merged, act)
			}
		}
	}
	slices.SortStableFunc(merged, func(a, b Action) int {
		return slices.Index(actionOrder, a.name) -
			slices.Index(actionOrder, b.name)
	})
	return merged
}
