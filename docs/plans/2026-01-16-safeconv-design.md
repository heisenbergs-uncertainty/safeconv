# safeconv Design Document

**Date:** 2026-01-16
**Module:** `github.com/heisenbergs-uncertainty/safeconv`
**Status:** Approved

## Overview

`safeconv` is a Go library for safe type conversions that detect overflow, underflow, and precision loss at runtime. Unlike Go's built-in type conversions which silently truncate or wrap, `safeconv` functions return errors when a conversion would lose data.

**Inspiration:** gosec rules G115 (integer overflow) and G109 (integer underflow).

## Scope

### v1.0 (Current)
- Integer conversions (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
- Float conversions (float32, float64)

### Future Versions
- Duration/time conversions (e.g., `DurationToMillis`)
- Numeric widening with consistent API
- Byte/rune conversions

### Out of Scope
`safeconv` is NOT a replacement for parsing or formatting. We do not duplicate functionality from `strconv` or `time`. The focus is strictly on safe type conversions.

## Design Decisions

### 1. Error Handling: Return Errors
Functions return `(value, error)` - idiomatic Go that forces explicit handling.

```go
val, err := safeconv.Int64ToUint32(x)
```

### 2. Package Structure: Single Flat Package
Following the `strconv` pattern, all functions live at the top level.

```go
import "github.com/heisenbergs-uncertainty/safeconv"

val, err := safeconv.Int64ToUint32(x)
```

### 3. Error Types: Hybrid Sentinel + Detailed
Sentinel errors for easy checking, wrapped in a detailed type for debugging.

```go
// Sentinel errors
var (
    ErrOverflow  = errors.New("value overflows target type")
    ErrUnderflow = errors.New("value underflows target type")
    ErrNaN       = errors.New("cannot convert NaN")
    ErrInfinity  = errors.New("cannot convert infinity")
)

// Detailed error type
type ConversionError struct {
    From  string // source type: "int64", "float64", etc.
    To    string // target type: "uint32", "int8", etc.
    Value any    // the problematic value
    Err   error  // underlying sentinel
}

func (e *ConversionError) Error() string
func (e *ConversionError) Unwrap() error
```

Usage:
```go
val, err := safeconv.Int64ToUint32(-5)
if errors.Is(err, safeconv.ErrUnderflow) {
    // handle negative value
}

var convErr *safeconv.ConversionError
if errors.As(err, &convErr) {
    log.Printf("failed to convert %v (%s) to %s", convErr.Value, convErr.From, convErr.To)
}
```

### 4. API Style: Explicit Function Names
One function per type pair for clarity. Generics used internally to reduce duplication.

```go
// Public API - explicit
func Int64ToUint32(x int64) (uint32, error)
func Float64TruncToInt64(x float64) (int64, error)

// Internal - generic helpers (unexported)
func signedToUnsigned[From SignedInt, To UnsignedInt](x From) (To, error)
```

### 5. Float-to-Integer: Separate Truncate/Round Functions
Callers must explicitly choose truncation or rounding behavior.

```go
safeconv.Float64TruncToInt64(3.7)  // → 3 (truncates toward zero)
safeconv.Float64RoundToInt64(3.7)  // → 4 (rounds to nearest)
```

### 6. Must Variants
Panic-on-error variants for initialization and tests.

```go
var limit = safeconv.MustInt64ToUint32(1000)  // panics if conversion fails
```

### 7. Special Float Values
- **Float-to-integer:** `NaN` returns `ErrNaN`, `±Inf` returns `ErrInfinity`
- **Float-to-float:** `NaN` and `±Inf` propagate (well-defined behavior)

### 8. Float64-to-Float32
Errors on overflow only. Precision loss is accepted as inherent to floating-point.

```go
safeconv.Float64ToFloat32(math.MaxFloat64)  // error (overflows)
safeconv.Float64ToFloat32(1.0000000001)     // success (precision loss accepted)
```

## API Surface

### Integer Conversions

**Signed to Unsigned:**
```go
func Int8ToUint8(x int8) (uint8, error)
func Int8ToUint16(x int8) (uint16, error)
func Int8ToUint32(x int8) (uint32, error)
func Int8ToUint64(x int8) (uint64, error)
func Int8ToUint(x int8) (uint, error)
// ... same pattern for Int16, Int32, Int64, Int
```

**Unsigned to Signed:**
```go
func Uint8ToInt8(x uint8) (int8, error)
func Uint16ToInt8(x uint16) (int8, error)
// ... etc
```

**Narrowing:**
```go
func Int64ToInt32(x int64) (int32, error)
func Uint64ToUint32(x uint64) (uint32, error)
// ... etc
```

### Float Conversions

**Float to Integer (truncate):**
```go
func Float32TruncToInt8(x float32) (int8, error)
func Float32TruncToInt16(x float32) (int16, error)
// ... all integer types for Float32 and Float64
```

**Float to Integer (round):**
```go
func Float32RoundToInt8(x float32) (int8, error)
func Float32RoundToInt16(x float32) (int16, error)
// ... all integer types for Float32 and Float64
```

**Float narrowing:**
```go
func Float64ToFloat32(x float64) (float32, error)
```

### Must Variants
Every function has a corresponding `Must` variant:
```go
func MustInt64ToUint32(x int64) uint32
func MustFloat64TruncToInt64(x float64) int64
// ... etc
```

## File Organization

```
safeconv/
├── go.mod
├── doc.go             # Package documentation
├── errors.go          # ConversionError, sentinel errors
├── constraints.go     # SignedInt, UnsignedInt, Float interfaces
├── convert.go         # Internal generic helpers (unexported)
├── int_to_uint.go     # Int* → Uint* conversions
├── uint_to_int.go     # Uint* → Int* conversions
├── int_narrow.go      # Int64 → Int32, Int16, etc.
├── uint_narrow.go     # Uint64 → Uint32, Uint16, etc.
├── float_trunc.go     # Float*TruncTo* conversions
├── float_round.go     # Float*RoundTo* conversions
├── float_narrow.go    # Float64 → Float32
├── must.go            # All Must* variants
└── *_test.go          # Corresponding test files
```

## Testing Strategy

Table-driven tests covering:
- Zero values
- Positive in-range values
- Boundary maximum (exactly at target type's max)
- Boundary overflow (max + 1)
- Negative values (for signed-to-unsigned)
- NaN/Inf (for float conversions)

Must variants tested for panic behavior.

## Implementation Notes

### Internal Generic Helpers

```go
type SignedInt interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
    ~float32 | ~float64
}
```

Public functions wrap internal generic helpers to maintain explicit API while reducing code duplication.
