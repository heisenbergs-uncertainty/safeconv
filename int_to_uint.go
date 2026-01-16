package safeconv

import "math"

// Int64ToUint64 converts an int64 to uint64.
// Returns ErrUnderflow if the value is negative.
func Int64ToUint64(x int64) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int64", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Int64ToUint32 converts an int64 to uint32.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func Int64ToUint32(x int64) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int64", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint32 {
		return 0, &ConversionError{From: "int64", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// Int64ToUint16 converts an int64 to uint16.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func Int64ToUint16(x int64) (uint16, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int64", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint16 {
		return 0, &ConversionError{From: "int64", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Int64ToUint8 converts an int64 to uint8.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func Int64ToUint8(x int64) (uint8, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int64", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "int64", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Int64ToUint converts an int64 to uint.
// Returns ErrUnderflow if negative, ErrOverflow if too large for the platform's uint.
func Int64ToUint(x int64) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int64", To: "uint", Value: x, Err: ErrUnderflow}
	}
	if uint64(x) > uint64(^uint(0)) {
		return 0, &ConversionError{From: "int64", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(x), nil
}
