package safeconv

import "math"

// ClampInt64ToUint64 converts an int64 to uint64, clamping to valid range.
// Negative values clamp to 0.
func ClampInt64ToUint64(x int64) uint64 {
	if x < 0 {
		return 0
	}
	return uint64(x)
}

// ClampInt64ToUint32 converts an int64 to uint32, clamping to valid range.
// Negative values clamp to 0, values > MaxUint32 clamp to MaxUint32.
func ClampInt64ToUint32(x int64) uint32 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(x)
}

// ClampInt64ToUint16 converts an int64 to uint16, clamping to valid range.
// Negative values clamp to 0, values > MaxUint16 clamp to MaxUint16.
func ClampInt64ToUint16(x int64) uint16 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(x)
}

// ClampInt64ToUint8 converts an int64 to uint8, clamping to valid range.
// Negative values clamp to 0, values > MaxUint8 clamp to MaxUint8.
func ClampInt64ToUint8(x int64) uint8 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}

// ClampInt64ToUint converts an int64 to uint, clamping to valid range.
// Negative values clamp to 0, values > max uint clamp to max uint.
func ClampInt64ToUint(x int64) uint {
	if x < 0 {
		return 0
	}
	if uint64(x) > uint64(^uint(0)) {
		return ^uint(0)
	}
	return uint(x)
}

// ClampInt32ToUint64 converts an int32 to uint64, clamping to valid range.
// Negative values clamp to 0.
func ClampInt32ToUint64(x int32) uint64 {
	if x < 0 {
		return 0
	}
	return uint64(x)
}

// ClampInt32ToUint32 converts an int32 to uint32, clamping to valid range.
// Negative values clamp to 0.
func ClampInt32ToUint32(x int32) uint32 {
	if x < 0 {
		return 0
	}
	return uint32(x)
}

// ClampInt32ToUint16 converts an int32 to uint16, clamping to valid range.
// Negative values clamp to 0, values > MaxUint16 clamp to MaxUint16.
func ClampInt32ToUint16(x int32) uint16 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(x)
}

// ClampInt32ToUint8 converts an int32 to uint8, clamping to valid range.
// Negative values clamp to 0, values > MaxUint8 clamp to MaxUint8.
func ClampInt32ToUint8(x int32) uint8 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}

// ClampInt32ToUint converts an int32 to uint, clamping to valid range.
// Negative values clamp to 0.
func ClampInt32ToUint(x int32) uint {
	if x < 0 {
		return 0
	}
	return uint(x)
}

// ClampInt16ToUint64 converts an int16 to uint64, clamping to valid range.
// Negative values clamp to 0.
func ClampInt16ToUint64(x int16) uint64 {
	if x < 0 {
		return 0
	}
	return uint64(x)
}

// ClampInt16ToUint32 converts an int16 to uint32, clamping to valid range.
// Negative values clamp to 0.
func ClampInt16ToUint32(x int16) uint32 {
	if x < 0 {
		return 0
	}
	return uint32(x)
}

// ClampInt16ToUint16 converts an int16 to uint16, clamping to valid range.
// Negative values clamp to 0.
func ClampInt16ToUint16(x int16) uint16 {
	if x < 0 {
		return 0
	}
	return uint16(x)
}

// ClampInt16ToUint8 converts an int16 to uint8, clamping to valid range.
// Negative values clamp to 0, values > MaxUint8 clamp to MaxUint8.
func ClampInt16ToUint8(x int16) uint8 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}

// ClampInt16ToUint converts an int16 to uint, clamping to valid range.
// Negative values clamp to 0.
func ClampInt16ToUint(x int16) uint {
	if x < 0 {
		return 0
	}
	return uint(x)
}

// ClampInt8ToUint64 converts an int8 to uint64, clamping to valid range.
// Negative values clamp to 0.
func ClampInt8ToUint64(x int8) uint64 {
	if x < 0 {
		return 0
	}
	return uint64(x)
}

// ClampInt8ToUint32 converts an int8 to uint32, clamping to valid range.
// Negative values clamp to 0.
func ClampInt8ToUint32(x int8) uint32 {
	if x < 0 {
		return 0
	}
	return uint32(x)
}

// ClampInt8ToUint16 converts an int8 to uint16, clamping to valid range.
// Negative values clamp to 0.
func ClampInt8ToUint16(x int8) uint16 {
	if x < 0 {
		return 0
	}
	return uint16(x)
}

// ClampInt8ToUint8 converts an int8 to uint8, clamping to valid range.
// Negative values clamp to 0.
func ClampInt8ToUint8(x int8) uint8 {
	if x < 0 {
		return 0
	}
	return uint8(x)
}

// ClampInt8ToUint converts an int8 to uint, clamping to valid range.
// Negative values clamp to 0.
func ClampInt8ToUint(x int8) uint {
	if x < 0 {
		return 0
	}
	return uint(x)
}

// ClampIntToUint64 converts an int to uint64, clamping to valid range.
// Negative values clamp to 0.
func ClampIntToUint64(x int) uint64 {
	if x < 0 {
		return 0
	}
	return uint64(x)
}

// ClampIntToUint32 converts an int to uint32, clamping to valid range.
// Negative values clamp to 0, values > MaxUint32 clamp to MaxUint32.
func ClampIntToUint32(x int) uint32 {
	if x < 0 {
		return 0
	}
	if uint64(x) > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(x)
}

// ClampIntToUint16 converts an int to uint16, clamping to valid range.
// Negative values clamp to 0, values > MaxUint16 clamp to MaxUint16.
func ClampIntToUint16(x int) uint16 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(x)
}

// ClampIntToUint8 converts an int to uint8, clamping to valid range.
// Negative values clamp to 0, values > MaxUint8 clamp to MaxUint8.
func ClampIntToUint8(x int) uint8 {
	if x < 0 {
		return 0
	}
	if x > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(x)
}

// ClampIntToUint converts an int to uint, clamping to valid range.
// Negative values clamp to 0.
func ClampIntToUint(x int) uint {
	if x < 0 {
		return 0
	}
	return uint(x)
}
