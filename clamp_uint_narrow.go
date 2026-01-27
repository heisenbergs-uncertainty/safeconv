package safeconv

import "math"

// =============================================================================
// Uint64 to Uint* clamping narrowing conversions
// =============================================================================

// ClampUint64ToUint32 converts a uint64 to uint32, clamping to valid range.
// Values > MaxUint32 clamp to MaxUint32.
func ClampUint64ToUint32(x uint64) uint32 {
	if x > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(x)
}

// ClampUint64ToUint16 converts a uint64 to uint16, clamping to valid range.
// Values > MaxUint16 clamp to MaxUint16.
func ClampUint64ToUint16(x uint64) uint16 {
	if x > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(x)
}

// ClampUint64ToUint8 converts a uint64 to uint8, clamping to valid range.
// Values > MaxUint8 clamp to MaxUint8.
func ClampUint64ToUint8(x uint64) uint8 {
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}

// ClampUint64ToUint converts a uint64 to uint, clamping to valid range.
// Values > max uint clamp to max uint.
func ClampUint64ToUint(x uint64) uint {
	if x > uint64(^uint(0)) {
		return ^uint(0)
	}
	return uint(x)
}

// =============================================================================
// Uint32 to Uint* clamping narrowing conversions
// =============================================================================

// ClampUint32ToUint16 converts a uint32 to uint16, clamping to valid range.
// Values > MaxUint16 clamp to MaxUint16.
func ClampUint32ToUint16(x uint32) uint16 {
	if x > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(x)
}

// ClampUint32ToUint8 converts a uint32 to uint8, clamping to valid range.
// Values > MaxUint8 clamp to MaxUint8.
func ClampUint32ToUint8(x uint32) uint8 {
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}

// =============================================================================
// Uint16 to Uint* clamping narrowing conversions
// =============================================================================

// ClampUint16ToUint8 converts a uint16 to uint8, clamping to valid range.
// Values > MaxUint8 clamp to MaxUint8.
func ClampUint16ToUint8(x uint16) uint8 {
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}

// =============================================================================
// Uint to Uint* clamping narrowing conversions
// =============================================================================

// ClampUintToUint32 converts a uint to uint32, clamping to valid range.
// Values > MaxUint32 clamp to MaxUint32.
func ClampUintToUint32(x uint) uint32 {
	if uint64(x) > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(x)
}

// ClampUintToUint16 converts a uint to uint16, clamping to valid range.
// Values > MaxUint16 clamp to MaxUint16.
func ClampUintToUint16(x uint) uint16 {
	if x > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(x)
}

// ClampUintToUint8 converts a uint to uint8, clamping to valid range.
// Values > MaxUint8 clamp to MaxUint8.
func ClampUintToUint8(x uint) uint8 {
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}
