// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"fmt"
	"math"
)

// Int64ToInt converts a given numeric value to int.
//
// It succeeds only if the numeric value can be exactly preserved without loss
// or truncation. Otherwise, it returns an error.
func Int64ToInt(value int64) (int, error) {
	if value < math.MinInt || value > math.MaxInt {
		return 0, fmt.Errorf("int64 %w for int", ErrInvRange)
	}
	return int(value), nil
}
