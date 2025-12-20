// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Int64ToInt8 converts a given numeric value to int8.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int64ToInt8(value int64) (int8, error) {
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, fmt.Errorf("int64 %w for int8", ErrInvRange)
	}
	return int8(value), nil
}
