package safeconv

import "math"

// =============================================================================
// Float64 to Int* clamping rounding conversions
// =============================================================================

// ClampFloat64RoundToInt64 converts a float64 to int64 by rounding to nearest,
// clamping overflow/underflow to min/max of int64.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToInt64(x float64) (int64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded >= float64(math.MaxInt64) {
		return math.MaxInt64, nil
	}
	if rounded <= float64(math.MinInt64) {
		return math.MinInt64, nil
	}
	return int64(rounded), nil
}

// ClampFloat64RoundToInt32 converts a float64 to int32 by rounding to nearest,
// clamping overflow/underflow to min/max of int32.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToInt32(x float64) (int32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded > float64(math.MaxInt32) {
		return math.MaxInt32, nil
	}
	if rounded < float64(math.MinInt32) {
		return math.MinInt32, nil
	}
	return int32(rounded), nil
}

// ClampFloat64RoundToInt16 converts a float64 to int16 by rounding to nearest,
// clamping overflow/underflow to min/max of int16.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToInt16(x float64) (int16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded > float64(math.MaxInt16) {
		return math.MaxInt16, nil
	}
	if rounded < float64(math.MinInt16) {
		return math.MinInt16, nil
	}
	return int16(rounded), nil
}

// ClampFloat64RoundToInt8 converts a float64 to int8 by rounding to nearest,
// clamping overflow/underflow to min/max of int8.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToInt8(x float64) (int8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded > float64(math.MaxInt8) {
		return math.MaxInt8, nil
	}
	if rounded < float64(math.MinInt8) {
		return math.MinInt8, nil
	}
	return int8(rounded), nil
}

// ClampFloat64RoundToInt converts a float64 to int by rounding to nearest,
// clamping overflow/underflow to min/max of int.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToInt(x float64) (int, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded >= float64(math.MaxInt) {
		return math.MaxInt, nil
	}
	if rounded <= float64(math.MinInt) {
		return math.MinInt, nil
	}
	return int(rounded), nil
}

// ClampFloat64RoundToUint64 converts a float64 to uint64 by rounding to nearest,
// clamping overflow to MaxUint64, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToUint64(x float64) (uint64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded < 0 {
		return 0, nil
	}
	if rounded >= float64(math.MaxUint64) {
		return math.MaxUint64, nil
	}
	return uint64(rounded), nil
}

// ClampFloat64RoundToUint32 converts a float64 to uint32 by rounding to nearest,
// clamping overflow to MaxUint32, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToUint32(x float64) (uint32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded < 0 {
		return 0, nil
	}
	if rounded > float64(math.MaxUint32) {
		return math.MaxUint32, nil
	}
	return uint32(rounded), nil
}

// ClampFloat64RoundToUint16 converts a float64 to uint16 by rounding to nearest,
// clamping overflow to MaxUint16, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToUint16(x float64) (uint16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded < 0 {
		return 0, nil
	}
	if rounded > float64(math.MaxUint16) {
		return math.MaxUint16, nil
	}
	return uint16(rounded), nil
}

// ClampFloat64RoundToUint8 converts a float64 to uint8 by rounding to nearest,
// clamping overflow to MaxUint8, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToUint8(x float64) (uint8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded < 0 {
		return 0, nil
	}
	if rounded > float64(math.MaxUint8) {
		return math.MaxUint8, nil
	}
	return uint8(rounded), nil
}

// ClampFloat64RoundToUint converts a float64 to uint by rounding to nearest,
// clamping overflow to max uint, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat64RoundToUint(x float64) (uint, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(x)
	if rounded < 0 {
		return 0, nil
	}
	if rounded >= float64(^uint(0)) {
		return ^uint(0), nil
	}
	return uint(rounded), nil
}

// =============================================================================
// Float32 to Int* clamping rounding conversions
// =============================================================================

// ClampFloat32RoundToInt64 converts a float32 to int64 by rounding to nearest,
// clamping overflow/underflow to min/max of int64.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToInt64(x float32) (int64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	return int64(rounded), nil
}

// ClampFloat32RoundToInt32 converts a float32 to int32 by rounding to nearest,
// clamping overflow/underflow to min/max of int32.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToInt32(x float32) (int32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded > float64(math.MaxInt32) {
		return math.MaxInt32, nil
	}
	if rounded < float64(math.MinInt32) {
		return math.MinInt32, nil
	}
	return int32(rounded), nil
}

// ClampFloat32RoundToInt16 converts a float32 to int16 by rounding to nearest,
// clamping overflow/underflow to min/max of int16.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToInt16(x float32) (int16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded > float64(math.MaxInt16) {
		return math.MaxInt16, nil
	}
	if rounded < float64(math.MinInt16) {
		return math.MinInt16, nil
	}
	return int16(rounded), nil
}

// ClampFloat32RoundToInt8 converts a float32 to int8 by rounding to nearest,
// clamping overflow/underflow to min/max of int8.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToInt8(x float32) (int8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded > float64(math.MaxInt8) {
		return math.MaxInt8, nil
	}
	if rounded < float64(math.MinInt8) {
		return math.MinInt8, nil
	}
	return int8(rounded), nil
}

// ClampFloat32RoundToInt converts a float32 to int by rounding to nearest,
// clamping overflow/underflow to min/max of int.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToInt(x float32) (int, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded >= float64(math.MaxInt) {
		return math.MaxInt, nil
	}
	if rounded <= float64(math.MinInt) {
		return math.MinInt, nil
	}
	return int(rounded), nil
}

// ClampFloat32RoundToUint64 converts a float32 to uint64 by rounding to nearest,
// clamping overflow to MaxUint64, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToUint64(x float32) (uint64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded < 0 {
		return 0, nil
	}
	return uint64(rounded), nil
}

// ClampFloat32RoundToUint32 converts a float32 to uint32 by rounding to nearest,
// clamping overflow to MaxUint32, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToUint32(x float32) (uint32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded < 0 {
		return 0, nil
	}
	if rounded > float64(math.MaxUint32) {
		return math.MaxUint32, nil
	}
	return uint32(rounded), nil
}

// ClampFloat32RoundToUint16 converts a float32 to uint16 by rounding to nearest,
// clamping overflow to MaxUint16, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToUint16(x float32) (uint16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded < 0 {
		return 0, nil
	}
	if rounded > float64(math.MaxUint16) {
		return math.MaxUint16, nil
	}
	return uint16(rounded), nil
}

// ClampFloat32RoundToUint8 converts a float32 to uint8 by rounding to nearest,
// clamping overflow to MaxUint8, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToUint8(x float32) (uint8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded < 0 {
		return 0, nil
	}
	if rounded > float64(math.MaxUint8) {
		return math.MaxUint8, nil
	}
	return uint8(rounded), nil
}

// ClampFloat32RoundToUint converts a float32 to uint by rounding to nearest,
// clamping overflow to max uint, negative values to 0.
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
func ClampFloat32RoundToUint(x float32) (uint, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrInfinity}
	}
	rounded := math.Round(float64(x))
	if rounded < 0 {
		return 0, nil
	}
	if rounded >= float64(^uint(0)) {
		return ^uint(0), nil
	}
	return uint(rounded), nil
}
