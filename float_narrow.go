package safeconv

import "math"

// Float64ToFloat32 converts a float64 to float32.
// Returns ErrOverflow if the value is too large for float32.
// NaN and Inf values are propagated (not treated as errors).
func Float64ToFloat32(x float64) (float32, error) {
	// NaN and Inf propagate (this is well-defined behavior)
	if math.IsNaN(x) {
		return float32(math.NaN()), nil
	}
	if math.IsInf(x, 1) {
		return float32(math.Inf(1)), nil
	}
	if math.IsInf(x, -1) {
		return float32(math.Inf(-1)), nil
	}

	// Check for overflow (value too large for float32)
	if x > math.MaxFloat32 || x < -math.MaxFloat32 {
		return 0, &ConversionError{From: "float64", To: "float32", Value: x, Err: ErrOverflow}
	}

	return float32(x), nil
}
