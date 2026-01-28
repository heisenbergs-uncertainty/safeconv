package safeconv

import (
	"errors"
	"fmt"
)

// Sentinel errors for type conversion failures.
var (
	// ErrOverflow indicates the value is too large for the target type.
	ErrOverflow = errors.New("value overflows target type")

	// ErrUnderflow indicates the value is too small for the target type
	// (e.g., negative value to unsigned type).
	ErrUnderflow = errors.New("value underflows target type")

	// ErrNaN indicates the value is NaN and cannot be converted.
	ErrNaN = errors.New("cannot convert NaN")

	// ErrInfinity indicates the value is infinite and cannot be converted.
	ErrInfinity = errors.New("cannot convert infinity")

	// ErrInvalidUnicode indicates the value is not a valid Unicode code point.
	// Valid code points are 0x0 to 0x10FFFF, excluding surrogates (0xD800-0xDFFF).
	ErrInvalidUnicode = errors.New("value is not a valid Unicode code point")
)

// ConversionError provides detailed information about a failed conversion.
type ConversionError struct {
	From  string // Source type name (e.g., "int64")
	To    string // Target type name (e.g., "uint32")
	Value any    // The value that failed to convert
	Err   error  // Underlying sentinel error
}

// Error returns a human-readable error message.
func (e *ConversionError) Error() string {
	return fmt.Sprintf("cannot convert %s(%v) to %s: %s", e.From, e.Value, e.To, e.Err)
}

// Unwrap returns the underlying sentinel error for use with errors.Is.
func (e *ConversionError) Unwrap() error {
	return e.Err
}

// IsOverflow returns true if the error is due to overflow.
func (e *ConversionError) IsOverflow() bool {
	return errors.Is(e.Err, ErrOverflow)
}

// IsUnderflow returns true if the error is due to underflow.
func (e *ConversionError) IsUnderflow() bool {
	return errors.Is(e.Err, ErrUnderflow)
}

// IsNaN returns true if the error is due to NaN.
func (e *ConversionError) IsNaN() bool {
	return errors.Is(e.Err, ErrNaN)
}

// IsInfinity returns true if the error is due to infinity.
func (e *ConversionError) IsInfinity() bool {
	return errors.Is(e.Err, ErrInfinity)
}

// IsInvalidUnicode returns true if the error is due to an invalid Unicode code point.
func (e *ConversionError) IsInvalidUnicode() bool {
	return errors.Is(e.Err, ErrInvalidUnicode)
}
