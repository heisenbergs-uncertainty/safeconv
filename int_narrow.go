package safeconv

import "math"

// =============================================================================
// Int64 to Int* narrowing conversions
// =============================================================================

// Int64ToInt32 converts an int64 to int32.
// Returns ErrOverflow if the value exceeds math.MaxInt32.
// Returns ErrUnderflow if the value is less than math.MinInt32.
func Int64ToInt32(x int64) (int32, error) {
	if x > math.MaxInt32 {
		return 0, &ConversionError{From: "int64", To: "int32", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt32 {
		return 0, &ConversionError{From: "int64", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(x), nil
}

// Int64ToInt16 converts an int64 to int16.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
// Returns ErrUnderflow if the value is less than math.MinInt16.
func Int64ToInt16(x int64) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "int64", To: "int16", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt16 {
		return 0, &ConversionError{From: "int64", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(x), nil
}

// Int64ToInt8 converts an int64 to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
// Returns ErrUnderflow if the value is less than math.MinInt8.
func Int64ToInt8(x int64) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "int64", To: "int8", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt8 {
		return 0, &ConversionError{From: "int64", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}

// Int64ToInt converts an int64 to int.
// Returns ErrOverflow if the value exceeds math.MaxInt.
// Returns ErrUnderflow if the value is less than math.MinInt.
// Note: On 64-bit platforms, this always succeeds.
func Int64ToInt(x int64) (int, error) {
	if x > math.MaxInt {
		return 0, &ConversionError{From: "int64", To: "int", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt {
		return 0, &ConversionError{From: "int64", To: "int", Value: x, Err: ErrUnderflow}
	}
	return int(x), nil
}

// =============================================================================
// Int32 to Int* narrowing conversions
// =============================================================================

// Int32ToInt64 converts an int32 to int64. Always succeeds.
func Int32ToInt64(x int32) int64 {
	return int64(x)
}

// Int32ToInt16 converts an int32 to int16.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
// Returns ErrUnderflow if the value is less than math.MinInt16.
func Int32ToInt16(x int32) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "int32", To: "int16", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt16 {
		return 0, &ConversionError{From: "int32", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(x), nil
}

// Int32ToInt8 converts an int32 to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
// Returns ErrUnderflow if the value is less than math.MinInt8.
func Int32ToInt8(x int32) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "int32", To: "int8", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt8 {
		return 0, &ConversionError{From: "int32", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}

// =============================================================================
// Int16 to Int* narrowing conversions
// =============================================================================

// Int16ToInt64 converts an int16 to int64. Always succeeds.
func Int16ToInt64(x int16) int64 {
	return int64(x)
}

// Int16ToInt32 converts an int16 to int32. Always succeeds.
func Int16ToInt32(x int16) int32 {
	return int32(x)
}

// Int16ToInt8 converts an int16 to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
// Returns ErrUnderflow if the value is less than math.MinInt8.
func Int16ToInt8(x int16) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "int16", To: "int8", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt8 {
		return 0, &ConversionError{From: "int16", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}

// =============================================================================
// Int8 to Int* widening conversions (always succeed)
// =============================================================================

// Int8ToInt64 converts an int8 to int64. Always succeeds.
func Int8ToInt64(x int8) int64 {
	return int64(x)
}

// Int8ToInt32 converts an int8 to int32. Always succeeds.
func Int8ToInt32(x int8) int32 {
	return int32(x)
}

// Int8ToInt16 converts an int8 to int16. Always succeeds.
func Int8ToInt16(x int8) int16 {
	return int16(x)
}

// =============================================================================
// Int to Int* conversions
// =============================================================================

// IntToInt64 converts an int to int64. Always succeeds.
func IntToInt64(x int) int64 {
	return int64(x)
}

// IntToInt32 converts an int to int32.
// Returns ErrOverflow if the value exceeds math.MaxInt32.
// Returns ErrUnderflow if the value is less than math.MinInt32.
func IntToInt32(x int) (int32, error) {
	if x > math.MaxInt32 {
		return 0, &ConversionError{From: "int", To: "int32", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt32 {
		return 0, &ConversionError{From: "int", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(x), nil
}

// IntToInt16 converts an int to int16.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
// Returns ErrUnderflow if the value is less than math.MinInt16.
func IntToInt16(x int) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "int", To: "int16", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt16 {
		return 0, &ConversionError{From: "int", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(x), nil
}

// IntToInt8 converts an int to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
// Returns ErrUnderflow if the value is less than math.MinInt8.
func IntToInt8(x int) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "int", To: "int8", Value: x, Err: ErrOverflow}
	}
	if x < math.MinInt8 {
		return 0, &ConversionError{From: "int", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}
