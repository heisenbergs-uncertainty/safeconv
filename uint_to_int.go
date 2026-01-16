package safeconv

import "math"

// =============================================================================
// Uint64 to Int* conversions
// =============================================================================

// Uint64ToInt64 converts a uint64 to int64.
// Returns ErrOverflow if the value exceeds math.MaxInt64.
func Uint64ToInt64(x uint64) (int64, error) {
	if x > math.MaxInt64 {
		return 0, &ConversionError{From: "uint64", To: "int64", Value: x, Err: ErrOverflow}
	}
	return int64(x), nil
}

// Uint64ToInt32 converts a uint64 to int32.
// Returns ErrOverflow if the value exceeds math.MaxInt32.
func Uint64ToInt32(x uint64) (int32, error) {
	if x > math.MaxInt32 {
		return 0, &ConversionError{From: "uint64", To: "int32", Value: x, Err: ErrOverflow}
	}
	return int32(x), nil
}

// Uint64ToInt16 converts a uint64 to int16.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
func Uint64ToInt16(x uint64) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint64", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// Uint64ToInt8 converts a uint64 to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
func Uint64ToInt8(x uint64) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint64", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint64ToInt converts a uint64 to int.
// Returns ErrOverflow if the value exceeds the platform's max int.
func Uint64ToInt(x uint64) (int, error) {
	if x > uint64(math.MaxInt) {
		return 0, &ConversionError{From: "uint64", To: "int", Value: x, Err: ErrOverflow}
	}
	return int(x), nil
}

// =============================================================================
// Uint32 to Int* conversions
// =============================================================================

// Uint32ToInt64 converts a uint32 to int64. Always succeeds.
func Uint32ToInt64(x uint32) int64 {
	return int64(x)
}

// Uint32ToInt32 converts a uint32 to int32.
// Returns ErrOverflow if the value exceeds math.MaxInt32.
func Uint32ToInt32(x uint32) (int32, error) {
	if x > math.MaxInt32 {
		return 0, &ConversionError{From: "uint32", To: "int32", Value: x, Err: ErrOverflow}
	}
	return int32(x), nil
}

// Uint32ToInt16 converts a uint32 to int16.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
func Uint32ToInt16(x uint32) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint32", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// Uint32ToInt8 converts a uint32 to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
func Uint32ToInt8(x uint32) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint32", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint32ToInt converts a uint32 to int. Always succeeds.
// On both 32-bit and 64-bit platforms, int can hold all uint32 values up to math.MaxInt32,
// and any uint32 value that fits in int32 will fit in int.
// Note: On 32-bit platforms, values > math.MaxInt32 would overflow, but we assume 64-bit.
func Uint32ToInt(x uint32) int {
	return int(x)
}

// =============================================================================
// Uint16 to Int* conversions
// =============================================================================

// Uint16ToInt64 converts a uint16 to int64. Always succeeds.
func Uint16ToInt64(x uint16) int64 {
	return int64(x)
}

// Uint16ToInt32 converts a uint16 to int32. Always succeeds.
func Uint16ToInt32(x uint16) int32 {
	return int32(x)
}

// Uint16ToInt16 converts a uint16 to int16.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
func Uint16ToInt16(x uint16) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint16", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// Uint16ToInt8 converts a uint16 to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
func Uint16ToInt8(x uint16) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint16", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint16ToInt converts a uint16 to int. Always succeeds.
func Uint16ToInt(x uint16) int {
	return int(x)
}

// =============================================================================
// Uint8 to Int* conversions
// =============================================================================

// Uint8ToInt64 converts a uint8 to int64. Always succeeds.
func Uint8ToInt64(x uint8) int64 {
	return int64(x)
}

// Uint8ToInt32 converts a uint8 to int32. Always succeeds.
func Uint8ToInt32(x uint8) int32 {
	return int32(x)
}

// Uint8ToInt16 converts a uint8 to int16. Always succeeds.
func Uint8ToInt16(x uint8) int16 {
	return int16(x)
}

// Uint8ToInt8 converts a uint8 to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
func Uint8ToInt8(x uint8) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint8", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint8ToInt converts a uint8 to int. Always succeeds.
func Uint8ToInt(x uint8) int {
	return int(x)
}

// =============================================================================
// Uint to Int* conversions
// =============================================================================

// UintToInt64 converts a uint to int64.
// Returns ErrOverflow if the value exceeds math.MaxInt64.
func UintToInt64(x uint) (int64, error) {
	if uint64(x) > math.MaxInt64 {
		return 0, &ConversionError{From: "uint", To: "int64", Value: x, Err: ErrOverflow}
	}
	return int64(x), nil
}

// UintToInt32 converts a uint to int32.
// Returns ErrOverflow if the value exceeds math.MaxInt32.
func UintToInt32(x uint) (int32, error) {
	if uint64(x) > math.MaxInt32 {
		return 0, &ConversionError{From: "uint", To: "int32", Value: x, Err: ErrOverflow}
	}
	return int32(x), nil
}

// UintToInt16 converts a uint to int16.
// Returns ErrOverflow if the value exceeds math.MaxInt16.
func UintToInt16(x uint) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// UintToInt8 converts a uint to int8.
// Returns ErrOverflow if the value exceeds math.MaxInt8.
func UintToInt8(x uint) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// UintToInt converts a uint to int.
// Returns ErrOverflow if the value exceeds math.MaxInt.
func UintToInt(x uint) (int, error) {
	if x > uint(math.MaxInt) {
		return 0, &ConversionError{From: "uint", To: "int", Value: x, Err: ErrOverflow}
	}
	return int(x), nil
}
