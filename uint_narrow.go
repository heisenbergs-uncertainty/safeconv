package safeconv

import "math"

// =============================================================================
// Uint64 to Uint* narrowing conversions
// =============================================================================

// Uint64ToUint32 converts a uint64 to uint32.
// Returns ErrOverflow if the value exceeds math.MaxUint32.
func Uint64ToUint32(x uint64) (uint32, error) {
	if x > math.MaxUint32 {
		return 0, &ConversionError{From: "uint64", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// Uint64ToUint16 converts a uint64 to uint16.
// Returns ErrOverflow if the value exceeds math.MaxUint16.
func Uint64ToUint16(x uint64) (uint16, error) {
	if x > math.MaxUint16 {
		return 0, &ConversionError{From: "uint64", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Uint64ToUint8 converts a uint64 to uint8.
// Returns ErrOverflow if the value exceeds math.MaxUint8.
func Uint64ToUint8(x uint64) (uint8, error) {
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "uint64", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Uint64ToUint converts a uint64 to uint.
// Returns ErrOverflow if the value exceeds the platform's max uint.
// Note: On 64-bit platforms, this always succeeds.
func Uint64ToUint(x uint64) (uint, error) {
	if x > uint64(^uint(0)) {
		return 0, &ConversionError{From: "uint64", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(x), nil
}

// =============================================================================
// Uint32 to Uint* conversions
// =============================================================================

// Uint32ToUint64 converts a uint32 to uint64. Always succeeds.
func Uint32ToUint64(x uint32) uint64 {
	return uint64(x)
}

// Uint32ToUint16 converts a uint32 to uint16.
// Returns ErrOverflow if the value exceeds math.MaxUint16.
func Uint32ToUint16(x uint32) (uint16, error) {
	if x > math.MaxUint16 {
		return 0, &ConversionError{From: "uint32", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Uint32ToUint8 converts a uint32 to uint8.
// Returns ErrOverflow if the value exceeds math.MaxUint8.
func Uint32ToUint8(x uint32) (uint8, error) {
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "uint32", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// =============================================================================
// Uint16 to Uint* conversions
// =============================================================================

// Uint16ToUint64 converts a uint16 to uint64. Always succeeds.
func Uint16ToUint64(x uint16) uint64 {
	return uint64(x)
}

// Uint16ToUint32 converts a uint16 to uint32. Always succeeds.
func Uint16ToUint32(x uint16) uint32 {
	return uint32(x)
}

// Uint16ToUint8 converts a uint16 to uint8.
// Returns ErrOverflow if the value exceeds math.MaxUint8.
func Uint16ToUint8(x uint16) (uint8, error) {
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "uint16", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// =============================================================================
// Uint8 to Uint* widening conversions (always succeed)
// =============================================================================

// Uint8ToUint64 converts a uint8 to uint64. Always succeeds.
func Uint8ToUint64(x uint8) uint64 {
	return uint64(x)
}

// Uint8ToUint32 converts a uint8 to uint32. Always succeeds.
func Uint8ToUint32(x uint8) uint32 {
	return uint32(x)
}

// Uint8ToUint16 converts a uint8 to uint16. Always succeeds.
func Uint8ToUint16(x uint8) uint16 {
	return uint16(x)
}

// =============================================================================
// Uint to Uint* conversions
// =============================================================================

// UintToUint64 converts a uint to uint64. Always succeeds.
func UintToUint64(x uint) uint64 {
	return uint64(x)
}

// UintToUint32 converts a uint to uint32.
// Returns ErrOverflow if the value exceeds math.MaxUint32.
func UintToUint32(x uint) (uint32, error) {
	if uint64(x) > math.MaxUint32 {
		return 0, &ConversionError{From: "uint", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// UintToUint16 converts a uint to uint16.
// Returns ErrOverflow if the value exceeds math.MaxUint16.
func UintToUint16(x uint) (uint16, error) {
	if x > math.MaxUint16 {
		return 0, &ConversionError{From: "uint", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// UintToUint8 converts a uint to uint8.
// Returns ErrOverflow if the value exceeds math.MaxUint8.
func UintToUint8(x uint) (uint8, error) {
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "uint", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}
