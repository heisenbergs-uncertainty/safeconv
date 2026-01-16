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

// Int32ToUint64 converts an int32 to uint64.
// Returns ErrUnderflow if the value is negative.
func Int32ToUint64(x int32) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Int32ToUint32 converts an int32 to uint32.
// Returns ErrUnderflow if the value is negative.
func Int32ToUint32(x int32) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	return uint32(x), nil
}

// Int32ToUint16 converts an int32 to uint16.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func Int32ToUint16(x int32) (uint16, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint16 {
		return 0, &ConversionError{From: "int32", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Int32ToUint8 converts an int32 to uint8.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func Int32ToUint8(x int32) (uint8, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "int32", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Int32ToUint converts an int32 to uint.
// Returns ErrUnderflow if the value is negative.
func Int32ToUint(x int32) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}

// Int16ToUint64 converts an int16 to uint64.
// Returns ErrUnderflow if the value is negative.
func Int16ToUint64(x int16) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Int16ToUint32 converts an int16 to uint32.
// Returns ErrUnderflow if the value is negative.
func Int16ToUint32(x int16) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	return uint32(x), nil
}

// Int16ToUint16 converts an int16 to uint16.
// Returns ErrUnderflow if the value is negative.
func Int16ToUint16(x int16) (uint16, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	return uint16(x), nil
}

// Int16ToUint8 converts an int16 to uint8.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func Int16ToUint8(x int16) (uint8, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "int16", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Int16ToUint converts an int16 to uint.
// Returns ErrUnderflow if the value is negative.
func Int16ToUint(x int16) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}

// Int8ToUint64 converts an int8 to uint64.
// Returns ErrUnderflow if the value is negative.
func Int8ToUint64(x int8) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Int8ToUint32 converts an int8 to uint32.
// Returns ErrUnderflow if the value is negative.
func Int8ToUint32(x int8) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	return uint32(x), nil
}

// Int8ToUint16 converts an int8 to uint16.
// Returns ErrUnderflow if the value is negative.
func Int8ToUint16(x int8) (uint16, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	return uint16(x), nil
}

// Int8ToUint8 converts an int8 to uint8.
// Returns ErrUnderflow if the value is negative.
func Int8ToUint8(x int8) (uint8, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	return uint8(x), nil
}

// Int8ToUint converts an int8 to uint.
// Returns ErrUnderflow if the value is negative.
func Int8ToUint(x int8) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}

// IntToUint64 converts an int to uint64.
// Returns ErrUnderflow if the value is negative.
func IntToUint64(x int) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// IntToUint32 converts an int to uint32.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func IntToUint32(x int) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if uint64(x) > math.MaxUint32 {
		return 0, &ConversionError{From: "int", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// IntToUint16 converts an int to uint16.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func IntToUint16(x int) (uint16, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint16 {
		return 0, &ConversionError{From: "int", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// IntToUint8 converts an int to uint8.
// Returns ErrUnderflow if negative, ErrOverflow if too large.
func IntToUint8(x int) (uint8, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if x > math.MaxUint8 {
		return 0, &ConversionError{From: "int", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// IntToUint converts an int to uint.
// Returns ErrUnderflow if the value is negative.
func IntToUint(x int) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}
