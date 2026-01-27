package safeconv

import "math"

// =============================================================================
// Float64 to Int* clamping truncation conversions
// =============================================================================

// ClampFloat64TruncToInt64 converts a float64 to int64 by truncating toward zero,
// clamping overflow/underflow to min/max of int64.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToInt64(x float64) (int64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrInfinity}
	}
	if x >= float64(math.MaxInt64) {
		return math.MaxInt64, nil
	}
	if x <= float64(math.MinInt64) {
		return math.MinInt64, nil
	}
	return int64(x), nil
}

// ClampFloat64TruncToInt32 converts a float64 to int32 by truncating toward zero,
// clamping overflow/underflow to min/max of int32.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToInt32(x float64) (int32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt32) {
		return math.MaxInt32, nil
	}
	if x < float64(math.MinInt32) {
		return math.MinInt32, nil
	}
	return int32(x), nil
}

// ClampFloat64TruncToInt16 converts a float64 to int16 by truncating toward zero,
// clamping overflow/underflow to min/max of int16.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToInt16(x float64) (int16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt16) {
		return math.MaxInt16, nil
	}
	if x < float64(math.MinInt16) {
		return math.MinInt16, nil
	}
	return int16(x), nil
}

// ClampFloat64TruncToInt8 converts a float64 to int8 by truncating toward zero,
// clamping overflow/underflow to min/max of int8.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToInt8(x float64) (int8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt8) {
		return math.MaxInt8, nil
	}
	if x < float64(math.MinInt8) {
		return math.MinInt8, nil
	}
	return int8(x), nil
}

// ClampFloat64TruncToInt converts a float64 to int by truncating toward zero,
// clamping overflow/underflow to min/max of int.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToInt(x float64) (int, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrInfinity}
	}
	if x >= float64(math.MaxInt) {
		return math.MaxInt, nil
	}
	if x <= float64(math.MinInt) {
		return math.MinInt, nil
	}
	return int(x), nil
}

// ClampFloat64TruncToUint64 converts a float64 to uint64 by truncating toward zero,
// clamping overflow to MaxUint64, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToUint64(x float64) (uint64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x >= float64(math.MaxUint64) {
		return math.MaxUint64, nil
	}
	return uint64(x), nil
}

// ClampFloat64TruncToUint32 converts a float64 to uint32 by truncating toward zero,
// clamping overflow to MaxUint32, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToUint32(x float64) (uint32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x > float64(math.MaxUint32) {
		return math.MaxUint32, nil
	}
	return uint32(x), nil
}

// ClampFloat64TruncToUint16 converts a float64 to uint16 by truncating toward zero,
// clamping overflow to MaxUint16, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToUint16(x float64) (uint16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x > float64(math.MaxUint16) {
		return math.MaxUint16, nil
	}
	return uint16(x), nil
}

// ClampFloat64TruncToUint8 converts a float64 to uint8 by truncating toward zero,
// clamping overflow to MaxUint8, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToUint8(x float64) (uint8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x > float64(math.MaxUint8) {
		return math.MaxUint8, nil
	}
	return uint8(x), nil
}

// ClampFloat64TruncToUint converts a float64 to uint by truncating toward zero,
// clamping overflow to max uint, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64TruncToUint(x float64) (uint, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x >= float64(^uint(0)) {
		return ^uint(0), nil
	}
	return uint(x), nil
}

// =============================================================================
// Float32 to Int* clamping truncation conversions
// =============================================================================

// ClampFloat32TruncToInt64 converts a float32 to int64 by truncating toward zero,
// clamping overflow/underflow to min/max of int64.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToInt64(x float32) (int64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrInfinity}
	}
	// float32 cannot represent values outside int64 range precisely
	return int64(x), nil
}

// ClampFloat32TruncToInt32 converts a float32 to int32 by truncating toward zero,
// clamping overflow/underflow to min/max of int32.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToInt32(x float32) (int32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrInfinity}
	}
	if x > float32(math.MaxInt32) {
		return math.MaxInt32, nil
	}
	if x < float32(math.MinInt32) {
		return math.MinInt32, nil
	}
	return int32(x), nil
}

// ClampFloat32TruncToInt16 converts a float32 to int16 by truncating toward zero,
// clamping overflow/underflow to min/max of int16.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToInt16(x float32) (int16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrInfinity}
	}
	if x > float32(math.MaxInt16) {
		return math.MaxInt16, nil
	}
	if x < float32(math.MinInt16) {
		return math.MinInt16, nil
	}
	return int16(x), nil
}

// ClampFloat32TruncToInt8 converts a float32 to int8 by truncating toward zero,
// clamping overflow/underflow to min/max of int8.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToInt8(x float32) (int8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrInfinity}
	}
	if x > float32(math.MaxInt8) {
		return math.MaxInt8, nil
	}
	if x < float32(math.MinInt8) {
		return math.MinInt8, nil
	}
	return int8(x), nil
}

// ClampFloat32TruncToInt converts a float32 to int by truncating toward zero,
// clamping overflow/underflow to min/max of int.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToInt(x float32) (int, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrInfinity}
	}
	// float32 max is ~3.4e38, which exceeds int64 max on 64-bit
	if float64(x) >= float64(math.MaxInt) {
		return math.MaxInt, nil
	}
	if float64(x) <= float64(math.MinInt) {
		return math.MinInt, nil
	}
	return int(x), nil
}

// ClampFloat32TruncToUint64 converts a float32 to uint64 by truncating toward zero,
// clamping overflow to MaxUint64, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToUint64(x float32) (uint64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	// float32 cannot represent MaxUint64 precisely
	return uint64(x), nil
}

// ClampFloat32TruncToUint32 converts a float32 to uint32 by truncating toward zero,
// clamping overflow to MaxUint32, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToUint32(x float32) (uint32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x > float32(math.MaxUint32) {
		return math.MaxUint32, nil
	}
	return uint32(x), nil
}

// ClampFloat32TruncToUint16 converts a float32 to uint16 by truncating toward zero,
// clamping overflow to MaxUint16, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToUint16(x float32) (uint16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x > float32(math.MaxUint16) {
		return math.MaxUint16, nil
	}
	return uint16(x), nil
}

// ClampFloat32TruncToUint8 converts a float32 to uint8 by truncating toward zero,
// clamping overflow to MaxUint8, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToUint8(x float32) (uint8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if x > float32(math.MaxUint8) {
		return math.MaxUint8, nil
	}
	return uint8(x), nil
}

// ClampFloat32TruncToUint converts a float32 to uint by truncating toward zero,
// clamping overflow to max uint, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32TruncToUint(x float32) (uint, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, nil
	}
	if float64(x) >= float64(^uint(0)) {
		return ^uint(0), nil
	}
	return uint(x), nil
}
