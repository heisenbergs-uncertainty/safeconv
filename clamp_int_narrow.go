package safeconv

import "math"

// =============================================================================
// Int64 to Int* clamping narrowing conversions
// =============================================================================

// ClampInt64ToInt32 converts an int64 to int32, clamping to valid range.
// Values > MaxInt32 clamp to MaxInt32, values < MinInt32 clamp to MinInt32.
func ClampInt64ToInt32(x int64) int32 {
	if x > math.MaxInt32 {
		return math.MaxInt32
	}
	if x < math.MinInt32 {
		return math.MinInt32
	}
	return int32(x)
}

// ClampInt64ToInt16 converts an int64 to int16, clamping to valid range.
// Values > MaxInt16 clamp to MaxInt16, values < MinInt16 clamp to MinInt16.
func ClampInt64ToInt16(x int64) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	if x < math.MinInt16 {
		return math.MinInt16
	}
	return int16(x)
}

// ClampInt64ToInt8 converts an int64 to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8, values < MinInt8 clamp to MinInt8.
func ClampInt64ToInt8(x int64) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	if x < math.MinInt8 {
		return math.MinInt8
	}
	return int8(x)
}

// ClampInt64ToInt converts an int64 to int, clamping to valid range.
// Values > MaxInt clamp to MaxInt, values < MinInt clamp to MinInt.
func ClampInt64ToInt(x int64) int {
	if x > math.MaxInt {
		return math.MaxInt
	}
	if x < math.MinInt {
		return math.MinInt
	}
	return int(x)
}

// =============================================================================
// Int32 to Int* clamping narrowing conversions
// =============================================================================

// ClampInt32ToInt16 converts an int32 to int16, clamping to valid range.
// Values > MaxInt16 clamp to MaxInt16, values < MinInt16 clamp to MinInt16.
func ClampInt32ToInt16(x int32) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	if x < math.MinInt16 {
		return math.MinInt16
	}
	return int16(x)
}

// ClampInt32ToInt8 converts an int32 to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8, values < MinInt8 clamp to MinInt8.
func ClampInt32ToInt8(x int32) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	if x < math.MinInt8 {
		return math.MinInt8
	}
	return int8(x)
}

// =============================================================================
// Int16 to Int* clamping narrowing conversions
// =============================================================================

// ClampInt16ToInt8 converts an int16 to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8, values < MinInt8 clamp to MinInt8.
func ClampInt16ToInt8(x int16) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	if x < math.MinInt8 {
		return math.MinInt8
	}
	return int8(x)
}

// =============================================================================
// Int to Int* clamping narrowing conversions
// =============================================================================

// ClampIntToInt32 converts an int to int32, clamping to valid range.
// Values > MaxInt32 clamp to MaxInt32, values < MinInt32 clamp to MinInt32.
func ClampIntToInt32(x int) int32 {
	if x > math.MaxInt32 {
		return math.MaxInt32
	}
	if x < math.MinInt32 {
		return math.MinInt32
	}
	return int32(x)
}

// ClampIntToInt16 converts an int to int16, clamping to valid range.
// Values > MaxInt16 clamp to MaxInt16, values < MinInt16 clamp to MinInt16.
func ClampIntToInt16(x int) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	if x < math.MinInt16 {
		return math.MinInt16
	}
	return int16(x)
}

// ClampIntToInt8 converts an int to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8, values < MinInt8 clamp to MinInt8.
func ClampIntToInt8(x int) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	if x < math.MinInt8 {
		return math.MinInt8
	}
	return int8(x)
}
