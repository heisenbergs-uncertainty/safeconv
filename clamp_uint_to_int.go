package safeconv

import "math"

// =============================================================================
// Uint64 to Int* clamping conversions
// =============================================================================

// ClampUint64ToInt64 converts a uint64 to int64, clamping to valid range.
// Values > MaxInt64 clamp to MaxInt64.
func ClampUint64ToInt64(x uint64) int64 {
	if x > math.MaxInt64 {
		return math.MaxInt64
	}
	return int64(x)
}

// ClampUint64ToInt32 converts a uint64 to int32, clamping to valid range.
// Values > MaxInt32 clamp to MaxInt32.
func ClampUint64ToInt32(x uint64) int32 {
	if x > math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(x)
}

// ClampUint64ToInt16 converts a uint64 to int16, clamping to valid range.
// Values > MaxInt16 clamp to MaxInt16.
func ClampUint64ToInt16(x uint64) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	return int16(x)
}

// ClampUint64ToInt8 converts a uint64 to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8.
func ClampUint64ToInt8(x uint64) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	return int8(x)
}

// ClampUint64ToInt converts a uint64 to int, clamping to valid range.
// Values > MaxInt clamp to MaxInt.
func ClampUint64ToInt(x uint64) int {
	if x > uint64(math.MaxInt) {
		return math.MaxInt
	}
	return int(x)
}

// =============================================================================
// Uint32 to Int* clamping conversions
// =============================================================================

// ClampUint32ToInt32 converts a uint32 to int32, clamping to valid range.
// Values > MaxInt32 clamp to MaxInt32.
func ClampUint32ToInt32(x uint32) int32 {
	if x > math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(x)
}

// ClampUint32ToInt16 converts a uint32 to int16, clamping to valid range.
// Values > MaxInt16 clamp to MaxInt16.
func ClampUint32ToInt16(x uint32) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	return int16(x)
}

// ClampUint32ToInt8 converts a uint32 to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8.
func ClampUint32ToInt8(x uint32) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	return int8(x)
}

// =============================================================================
// Uint16 to Int* clamping conversions
// =============================================================================

// ClampUint16ToInt16 converts a uint16 to int16, clamping to valid range.
// Values > MaxInt16 clamp to MaxInt16.
func ClampUint16ToInt16(x uint16) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	return int16(x)
}

// ClampUint16ToInt8 converts a uint16 to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8.
func ClampUint16ToInt8(x uint16) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	return int8(x)
}

// =============================================================================
// Uint8 to Int* clamping conversions
// =============================================================================

// ClampUint8ToInt8 converts a uint8 to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8.
func ClampUint8ToInt8(x uint8) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	return int8(x)
}

// =============================================================================
// Uint to Int* clamping conversions
// =============================================================================

// ClampUintToInt64 converts a uint to int64, clamping to valid range.
// Values > MaxInt64 clamp to MaxInt64.
func ClampUintToInt64(x uint) int64 {
	if uint64(x) > math.MaxInt64 {
		return math.MaxInt64
	}
	return int64(x)
}

// ClampUintToInt32 converts a uint to int32, clamping to valid range.
// Values > MaxInt32 clamp to MaxInt32.
func ClampUintToInt32(x uint) int32 {
	if uint64(x) > math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(x)
}

// ClampUintToInt16 converts a uint to int16, clamping to valid range.
// Values > MaxInt16 clamp to MaxInt16.
func ClampUintToInt16(x uint) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	return int16(x)
}

// ClampUintToInt8 converts a uint to int8, clamping to valid range.
// Values > MaxInt8 clamp to MaxInt8.
func ClampUintToInt8(x uint) int8 {
	if x > math.MaxInt8 {
		return math.MaxInt8
	}
	return int8(x)
}

// ClampUintToInt converts a uint to int, clamping to valid range.
// Values > MaxInt clamp to MaxInt.
func ClampUintToInt(x uint) int {
	if x > uint(math.MaxInt) {
		return math.MaxInt
	}
	return int(x)
}
