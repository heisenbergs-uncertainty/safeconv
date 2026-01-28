package safeconv

import "unicode/utf8"

// Unicode constants for validation
const (
	maxUnicode      = 0x10FFFF // Maximum valid Unicode code point
	surrogateMin    = 0xD800   // Start of surrogate range
	surrogateMax    = 0xDFFF   // End of surrogate range
	replacementChar = '\uFFFD' // Unicode replacement character
)

// isValidUnicode checks if a value is a valid Unicode code point.
// Valid code points are 0x0 to 0x10FFFF, excluding surrogates (0xD800-0xDFFF).
func isValidUnicode(v int64) bool {
	return v >= 0 && v <= maxUnicode && (v < surrogateMin || v > surrogateMax)
}

// =============================================================================
// Signed Integer to Rune conversions
// =============================================================================

// Int64ToRune converts an int64 to a rune (Unicode code point).
// Returns ErrUnderflow if negative.
// Returns ErrInvalidUnicode if the value is not a valid Unicode code point
// (outside 0x0-0x10FFFF or in surrogate range 0xD800-0xDFFF).
func Int64ToRune(x int64) (rune, error) {
	if x < 0 {
		return replacementChar, &ConversionError{From: "int64", To: "rune", Value: x, Err: ErrUnderflow}
	}
	if !isValidUnicode(x) {
		return replacementChar, &ConversionError{From: "int64", To: "rune", Value: x, Err: ErrInvalidUnicode}
	}
	return rune(x), nil
}

// Int32ToRune converts an int32 to a rune (Unicode code point).
// Returns ErrUnderflow if negative.
// Returns ErrInvalidUnicode if the value is not a valid Unicode code point.
func Int32ToRune(x int32) (rune, error) {
	if x < 0 {
		return replacementChar, &ConversionError{From: "int32", To: "rune", Value: x, Err: ErrUnderflow}
	}
	if !isValidUnicode(int64(x)) {
		return replacementChar, &ConversionError{From: "int32", To: "rune", Value: x, Err: ErrInvalidUnicode}
	}
	return x, nil // rune is an alias for int32
}

// IntToRune converts an int to a rune (Unicode code point).
// Returns ErrUnderflow if negative.
// Returns ErrInvalidUnicode if the value is not a valid Unicode code point.
func IntToRune(x int) (rune, error) {
	if x < 0 {
		return replacementChar, &ConversionError{From: "int", To: "rune", Value: x, Err: ErrUnderflow}
	}
	if !isValidUnicode(int64(x)) {
		return replacementChar, &ConversionError{From: "int", To: "rune", Value: x, Err: ErrInvalidUnicode}
	}
	return rune(x), nil
}

// Int16ToRune converts an int16 to a rune (Unicode code point).
// Returns ErrUnderflow if negative.
// Note: All non-negative int16 values are valid Unicode code points (max 32767).
func Int16ToRune(x int16) (rune, error) {
	if x < 0 {
		return replacementChar, &ConversionError{From: "int16", To: "rune", Value: x, Err: ErrUnderflow}
	}
	return rune(x), nil
}

// Int8ToRune converts an int8 to a rune (Unicode code point).
// Returns ErrUnderflow if negative.
// Note: All non-negative int8 values are valid Unicode code points (max 127).
func Int8ToRune(x int8) (rune, error) {
	if x < 0 {
		return replacementChar, &ConversionError{From: "int8", To: "rune", Value: x, Err: ErrUnderflow}
	}
	return rune(x), nil
}

// =============================================================================
// Unsigned Integer to Rune conversions
// =============================================================================

// Uint64ToRune converts a uint64 to a rune (Unicode code point).
// Returns ErrInvalidUnicode if the value is not a valid Unicode code point.
func Uint64ToRune(x uint64) (rune, error) {
	if x > maxUnicode || (x >= surrogateMin && x <= surrogateMax) {
		return replacementChar, &ConversionError{From: "uint64", To: "rune", Value: x, Err: ErrInvalidUnicode}
	}
	return rune(x), nil
}

// Uint32ToRune converts a uint32 to a rune (Unicode code point).
// Returns ErrInvalidUnicode if the value is not a valid Unicode code point.
func Uint32ToRune(x uint32) (rune, error) {
	if x > maxUnicode || (x >= surrogateMin && x <= surrogateMax) {
		return replacementChar, &ConversionError{From: "uint32", To: "rune", Value: x, Err: ErrInvalidUnicode}
	}
	return rune(x), nil
}

// UintToRune converts a uint to a rune (Unicode code point).
// Returns ErrInvalidUnicode if the value is not a valid Unicode code point.
func UintToRune(x uint) (rune, error) {
	if x > maxUnicode || (x >= surrogateMin && x <= surrogateMax) {
		return replacementChar, &ConversionError{From: "uint", To: "rune", Value: x, Err: ErrInvalidUnicode}
	}
	return rune(x), nil
}

// Uint16ToRune converts a uint16 to a rune (Unicode code point).
// Returns ErrInvalidUnicode if the value is in the surrogate range (0xD800-0xDFFF).
// Note: uint16 max (65535) is within valid Unicode range, but surrogates must be rejected.
func Uint16ToRune(x uint16) (rune, error) {
	if x >= surrogateMin && x <= surrogateMax {
		return replacementChar, &ConversionError{From: "uint16", To: "rune", Value: x, Err: ErrInvalidUnicode}
	}
	return rune(x), nil
}

// Uint8ToRune converts a uint8 to a rune (Unicode code point). Always succeeds.
// All uint8 values (0-255) are valid Unicode code points.
func Uint8ToRune(x uint8) rune {
	return rune(x)
}

// =============================================================================
// Rune to Integer conversions
// =============================================================================

// RuneToInt64 converts a rune to int64. Always succeeds.
// Note: This does not validate that the rune is a valid Unicode code point.
// Use unicode.ValidRune() if validation is needed.
func RuneToInt64(r rune) int64 {
	return int64(r)
}

// RuneToInt32 converts a rune to int32. Always succeeds.
// rune is an alias for int32, so this is a no-op.
func RuneToInt32(r rune) int32 {
	return r
}

// RuneToInt converts a rune to int. Always succeeds.
func RuneToInt(r rune) int {
	return int(r)
}

// RuneToUint64 converts a rune to uint64.
// Returns ErrUnderflow if the rune value is negative (invalid rune).
func RuneToUint64(r rune) (uint64, error) {
	if r < 0 {
		return 0, &ConversionError{From: "rune", To: "uint64", Value: r, Err: ErrUnderflow}
	}
	return uint64(r), nil
}

// RuneToUint32 converts a rune to uint32.
// Returns ErrUnderflow if the rune value is negative (invalid rune).
func RuneToUint32(r rune) (uint32, error) {
	if r < 0 {
		return 0, &ConversionError{From: "rune", To: "uint32", Value: r, Err: ErrUnderflow}
	}
	return uint32(r), nil
}

// RuneToUint converts a rune to uint.
// Returns ErrUnderflow if the rune value is negative (invalid rune).
func RuneToUint(r rune) (uint, error) {
	if r < 0 {
		return 0, &ConversionError{From: "rune", To: "uint", Value: r, Err: ErrUnderflow}
	}
	return uint(r), nil
}

// RuneToUint16 converts a rune to uint16.
// Returns ErrUnderflow if negative.
// Returns ErrOverflow if the rune exceeds math.MaxUint16 (65535).
func RuneToUint16(r rune) (uint16, error) {
	if r < 0 {
		return 0, &ConversionError{From: "rune", To: "uint16", Value: r, Err: ErrUnderflow}
	}
	if r > 0xFFFF {
		return 0, &ConversionError{From: "rune", To: "uint16", Value: r, Err: ErrOverflow}
	}
	return uint16(r), nil
}

// RuneToUint8 converts a rune to uint8.
// Returns ErrUnderflow if negative.
// Returns ErrOverflow if the rune exceeds math.MaxUint8 (255).
func RuneToUint8(r rune) (uint8, error) {
	if r < 0 {
		return 0, &ConversionError{From: "rune", To: "uint8", Value: r, Err: ErrUnderflow}
	}
	if r > 0xFF {
		return 0, &ConversionError{From: "rune", To: "uint8", Value: r, Err: ErrOverflow}
	}
	return uint8(r), nil
}

// =============================================================================
// Must variants for Rune conversions
// =============================================================================

// MustInt64ToRune is like Int64ToRune but panics on error.
func MustInt64ToRune(x int64) rune {
	r, err := Int64ToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustInt32ToRune is like Int32ToRune but panics on error.
func MustInt32ToRune(x int32) rune {
	r, err := Int32ToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustIntToRune is like IntToRune but panics on error.
func MustIntToRune(x int) rune {
	r, err := IntToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustInt16ToRune is like Int16ToRune but panics on error.
func MustInt16ToRune(x int16) rune {
	r, err := Int16ToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustInt8ToRune is like Int8ToRune but panics on error.
func MustInt8ToRune(x int8) rune {
	r, err := Int8ToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustUint64ToRune is like Uint64ToRune but panics on error.
func MustUint64ToRune(x uint64) rune {
	r, err := Uint64ToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustUint32ToRune is like Uint32ToRune but panics on error.
func MustUint32ToRune(x uint32) rune {
	r, err := Uint32ToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustUintToRune is like UintToRune but panics on error.
func MustUintToRune(x uint) rune {
	r, err := UintToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustUint16ToRune is like Uint16ToRune but panics on error.
func MustUint16ToRune(x uint16) rune {
	r, err := Uint16ToRune(x)
	if err != nil {
		panic(err)
	}
	return r
}

// MustRuneToUint64 is like RuneToUint64 but panics on error.
func MustRuneToUint64(r rune) uint64 {
	v, err := RuneToUint64(r)
	if err != nil {
		panic(err)
	}
	return v
}

// MustRuneToUint32 is like RuneToUint32 but panics on error.
func MustRuneToUint32(r rune) uint32 {
	v, err := RuneToUint32(r)
	if err != nil {
		panic(err)
	}
	return v
}

// MustRuneToUint is like RuneToUint but panics on error.
func MustRuneToUint(r rune) uint {
	v, err := RuneToUint(r)
	if err != nil {
		panic(err)
	}
	return v
}

// MustRuneToUint16 is like RuneToUint16 but panics on error.
func MustRuneToUint16(r rune) uint16 {
	v, err := RuneToUint16(r)
	if err != nil {
		panic(err)
	}
	return v
}

// MustRuneToUint8 is like RuneToUint8 but panics on error.
func MustRuneToUint8(r rune) uint8 {
	v, err := RuneToUint8(r)
	if err != nil {
		panic(err)
	}
	return v
}

// =============================================================================
// Validation helper
// =============================================================================

// IsValidRune reports whether the rune is a valid Unicode code point.
// This is equivalent to utf8.ValidRune but provided here for convenience.
func IsValidRune(r rune) bool {
	return utf8.ValidRune(r)
}
