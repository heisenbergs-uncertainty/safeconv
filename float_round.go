package safeconv

import "math"

// =============================================================================
// Float64 to Int* rounding conversions (round half away from zero)
// =============================================================================

// Float64RoundToInt64 converts a float64 to int64 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt64.
// Returns ErrUnderflow if the rounded value is less than math.MinInt64.
func Float64RoundToInt64(x float64) (int64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded > float64(math.MaxInt64) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt64) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrUnderflow}
	}
	return int64(rounded), nil
}

// Float64RoundToInt32 converts a float64 to int32 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt32.
// Returns ErrUnderflow if the rounded value is less than math.MinInt32.
func Float64RoundToInt32(x float64) (int32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded > float64(math.MaxInt32) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt32) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(rounded), nil
}

// Float64RoundToInt16 converts a float64 to int16 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt16.
// Returns ErrUnderflow if the rounded value is less than math.MinInt16.
func Float64RoundToInt16(x float64) (int16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded > float64(math.MaxInt16) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt16) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(rounded), nil
}

// Float64RoundToInt8 converts a float64 to int8 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt8.
// Returns ErrUnderflow if the rounded value is less than math.MinInt8.
func Float64RoundToInt8(x float64) (int8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded > float64(math.MaxInt8) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt8) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(rounded), nil
}

// Float64RoundToInt converts a float64 to int by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt.
// Returns ErrUnderflow if the rounded value is less than math.MinInt.
func Float64RoundToInt(x float64) (int, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded > float64(math.MaxInt) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrUnderflow}
	}
	return int(rounded), nil
}

// =============================================================================
// Float64 to Uint* rounding conversions (round half away from zero)
// =============================================================================

// Float64RoundToUint64 converts a float64 to uint64 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint64.
func Float64RoundToUint64(x float64) (uint64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded < 0 {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint64) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrOverflow}
	}
	return uint64(rounded), nil
}

// Float64RoundToUint32 converts a float64 to uint32 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint32.
func Float64RoundToUint32(x float64) (uint32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded < 0 {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint32) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(rounded), nil
}

// Float64RoundToUint16 converts a float64 to uint16 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint16.
func Float64RoundToUint16(x float64) (uint16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded < 0 {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint16) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(rounded), nil
}

// Float64RoundToUint8 converts a float64 to uint8 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint8.
func Float64RoundToUint8(x float64) (uint8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded < 0 {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint8) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(rounded), nil
}

// Float64RoundToUint converts a float64 to uint by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint.
func Float64RoundToUint(x float64) (uint, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(x)

	if rounded < 0 {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(rounded), nil
}

// =============================================================================
// Float32 to Int* rounding conversions (round half away from zero)
// =============================================================================

// Float32RoundToInt64 converts a float32 to int64 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt64.
// Returns ErrUnderflow if the rounded value is less than math.MinInt64.
func Float32RoundToInt64(x float32) (int64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded > float64(math.MaxInt64) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt64) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrUnderflow}
	}
	return int64(rounded), nil
}

// Float32RoundToInt32 converts a float32 to int32 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt32.
// Returns ErrUnderflow if the rounded value is less than math.MinInt32.
func Float32RoundToInt32(x float32) (int32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded > float64(math.MaxInt32) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt32) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(rounded), nil
}

// Float32RoundToInt16 converts a float32 to int16 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt16.
// Returns ErrUnderflow if the rounded value is less than math.MinInt16.
func Float32RoundToInt16(x float32) (int16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded > float64(math.MaxInt16) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt16) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(rounded), nil
}

// Float32RoundToInt8 converts a float32 to int8 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt8.
// Returns ErrUnderflow if the rounded value is less than math.MinInt8.
func Float32RoundToInt8(x float32) (int8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded > float64(math.MaxInt8) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt8) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(rounded), nil
}

// Float32RoundToInt converts a float32 to int by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrOverflow if the rounded value exceeds math.MaxInt.
// Returns ErrUnderflow if the rounded value is less than math.MinInt.
func Float32RoundToInt(x float32) (int, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded > float64(math.MaxInt) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrOverflow}
	}
	if rounded < float64(math.MinInt) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrUnderflow}
	}
	return int(rounded), nil
}

// =============================================================================
// Float32 to Uint* rounding conversions (round half away from zero)
// =============================================================================

// Float32RoundToUint64 converts a float32 to uint64 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint64.
func Float32RoundToUint64(x float32) (uint64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded < 0 {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint64) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrOverflow}
	}
	return uint64(rounded), nil
}

// Float32RoundToUint32 converts a float32 to uint32 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint32.
func Float32RoundToUint32(x float32) (uint32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded < 0 {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint32) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(rounded), nil
}

// Float32RoundToUint16 converts a float32 to uint16 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint16.
func Float32RoundToUint16(x float32) (uint16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded < 0 {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint16) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(rounded), nil
}

// Float32RoundToUint8 converts a float32 to uint8 by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint8.
func Float32RoundToUint8(x float32) (uint8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded < 0 {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint8) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(rounded), nil
}

// Float32RoundToUint converts a float32 to uint by rounding to nearest.
// Ties are rounded away from zero (half away from zero rounding).
// Returns ErrNaN if the value is NaN.
// Returns ErrInfinity if the value is infinite.
// Returns ErrUnderflow if the rounded value is negative.
// Returns ErrOverflow if the rounded value exceeds math.MaxUint.
func Float32RoundToUint(x float32) (uint, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrInfinity}
	}

	rounded := math.Round(float64(x))

	if rounded < 0 {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrUnderflow}
	}
	if rounded > float64(math.MaxUint) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(rounded), nil
}
