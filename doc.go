// Package safeconv provides safe type conversions that detect overflow,
// underflow, and precision loss at runtime.
//
// Unlike Go's built-in type conversions which silently truncate or wrap,
// safeconv functions return errors when a conversion would lose data.
//
// # Integer Conversions
//
// Convert between signed and unsigned integers of any size:
//
//	val, err := safeconv.Int64ToUint32(x)
//	if errors.Is(err, safeconv.ErrOverflow) {
//	    // handle overflow
//	}
//	if errors.Is(err, safeconv.ErrUnderflow) {
//	    // handle negative value
//	}
//
// # Float Conversions
//
// Float-to-integer conversions require explicit truncation or rounding:
//
//	val, err := safeconv.Float64TruncToInt64(x)  // truncates toward zero
//	val, err := safeconv.Float64RoundToInt64(x)  // rounds to nearest
//
// Special float values (NaN, Inf) return specific errors:
//
//	if errors.Is(err, safeconv.ErrNaN) { ... }
//	if errors.Is(err, safeconv.ErrInfinity) { ... }
//
// # Must Variants
//
// For cases where failure is a programming error (e.g., constants),
// Must variants panic instead of returning errors:
//
//	var limit = safeconv.MustInt64ToUint32(10000)
//
// # Error Handling
//
// All errors can be checked with errors.Is for the underlying cause,
// or errors.As to get full conversion context:
//
//	var convErr *safeconv.ConversionError
//	if errors.As(err, &convErr) {
//	    log.Printf("failed: %s(%v) -> %s", convErr.From, convErr.Value, convErr.To)
//	}
package safeconv
