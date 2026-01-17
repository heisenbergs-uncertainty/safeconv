package safeconv

import "math"

// =============================================================================
// Float64 to Int* truncation conversions
// =============================================================================

// Float64TruncToInt64 converts a float64 to int64 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt64.
// Returns ErrUnderflow if the value is less than math.MinInt64.
func Float64TruncToInt64(x float64) (int64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt64) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt64) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrUnderflow}
	}
	return int64(x), nil
}

// Float64TruncToInt32 converts a float64 to int32 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt32.
// Returns ErrUnderflow if the value is less than math.MinInt32.
func Float64TruncToInt32(x float64) (int32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt32) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt32) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(x), nil
}

// Float64TruncToInt16 converts a float64 to int16 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
// Returns ErrUnderflow if the value is less than math.MinInt16.
func Float64TruncToInt16(x float64) (int16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt16) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt16) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(x), nil
}

// Float64TruncToInt8 converts a float64 to int8 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
// Returns ErrUnderflow if the value is less than math.MinInt8.
func Float64TruncToInt8(x float64) (int8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt8) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt8) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}

// Float64TruncToInt converts a float64 to int by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt.
// Returns ErrUnderflow if the value is less than math.MinInt.
func Float64TruncToInt(x float64) (int, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrUnderflow}
	}
	return int(x), nil
}

// =============================================================================
// Float64 to Uint* truncation conversions
// =============================================================================

// Float64TruncToUint64 converts a float64 to uint64 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint64.
func Float64TruncToUint64(x float64) (uint64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint64) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrOverflow}
	}
	return uint64(x), nil
}

// Float64TruncToUint32 converts a float64 to uint32 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint32.
func Float64TruncToUint32(x float64) (uint32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint32) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// Float64TruncToUint16 converts a float64 to uint16 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint16.
func Float64TruncToUint16(x float64) (uint16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint16) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Float64TruncToUint8 converts a float64 to uint8 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint8.
func Float64TruncToUint8(x float64) (uint8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint8) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Float64TruncToUint converts a float64 to uint by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint.
func Float64TruncToUint(x float64) (uint, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(x), nil
}

// =============================================================================
// Float32 to Int* truncation conversions
// =============================================================================

// Float32TruncToInt64 converts a float32 to int64 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt64.
// Returns ErrUnderflow if the value is less than math.MinInt64.
func Float32TruncToInt64(x float32) (int64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrInfinity}
	}
	if float64(x) > float64(math.MaxInt64) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrOverflow}
	}
	if float64(x) < float64(math.MinInt64) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrUnderflow}
	}
	return int64(x), nil
}

// Float32TruncToInt32 converts a float32 to int32 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt32.
// Returns ErrUnderflow if the value is less than math.MinInt32.
func Float32TruncToInt32(x float32) (int32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrInfinity}
	}
	if float64(x) > float64(math.MaxInt32) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrOverflow}
	}
	if float64(x) < float64(math.MinInt32) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(x), nil
}

// Float32TruncToInt16 converts a float32 to int16 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
// Returns ErrUnderflow if the value is less than math.MinInt16.
func Float32TruncToInt16(x float32) (int16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrInfinity}
	}
	if float64(x) > float64(math.MaxInt16) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrOverflow}
	}
	if float64(x) < float64(math.MinInt16) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(x), nil
}

// Float32TruncToInt8 converts a float32 to int8 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
// Returns ErrUnderflow if the value is less than math.MinInt8.
func Float32TruncToInt8(x float32) (int8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrInfinity}
	}
	if float64(x) > float64(math.MaxInt8) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrOverflow}
	}
	if float64(x) < float64(math.MinInt8) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}

// Float32TruncToInt converts a float32 to int by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the value exceeds math.MaxInt.
// Returns ErrUnderflow if the value is less than math.MinInt.
func Float32TruncToInt(x float32) (int, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrInfinity}
	}
	if float64(x) > float64(math.MaxInt) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrOverflow}
	}
	if float64(x) < float64(math.MinInt) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrUnderflow}
	}
	return int(x), nil
}

// =============================================================================
// Float32 to Uint* truncation conversions
// =============================================================================

// Float32TruncToUint64 converts a float32 to uint64 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint64.
func Float32TruncToUint64(x float32) (uint64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	if float64(x) > float64(math.MaxUint64) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrOverflow}
	}
	return uint64(x), nil
}

// Float32TruncToUint32 converts a float32 to uint32 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint32.
func Float32TruncToUint32(x float32) (uint32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if float64(x) > float64(math.MaxUint32) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// Float32TruncToUint16 converts a float32 to uint16 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint16.
func Float32TruncToUint16(x float32) (uint16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if float64(x) > float64(math.MaxUint16) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Float32TruncToUint8 converts a float32 to uint8 by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint8.
func Float32TruncToUint8(x float32) (uint8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if float64(x) > float64(math.MaxUint8) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Float32TruncToUint converts a float32 to uint by truncating toward zero.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the value is negative.
// Returns ErrOverflow if the value exceeds math.MaxUint.
func Float32TruncToUint(x float32) (uint, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrUnderflow}
	}
	if float64(x) > float64(math.MaxUint) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(x), nil
}
