# safeconv

[![Go Reference](https://pkg.go.dev/badge/github.com/heisenbergs-uncertainty/safeconv.svg)](https://pkg.go.dev/github.com/heisenbergs-uncertainty/safeconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/heisenbergs-uncertainty/safeconv)](https://goreportcard.com/report/github.com/heisenbergs-uncertainty/safeconv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Safe type conversions for Go that detect overflow, underflow, and precision loss at runtime.**

Unlike Go's built-in type conversions which silently truncate or wrap, `safeconv` functions return errors when a conversion would lose data.

```go
// Go's built-in conversion silently overflows
x := int64(math.MaxInt64)
y := int32(x)  // y = -1 (silent wraparound!)

// safeconv catches the overflow
y, err := safeconv.Int64ToInt32(x)
if errors.Is(err, safeconv.ErrOverflow) {
    // Handle the overflow safely
}
```

## Installation

```bash
go get github.com/heisenbergs-uncertainty/safeconv
```

Requires Go 1.21 or later.

## Why safeconv?

Go's type conversions are fast but dangerous. They silently produce incorrect results:

| Scenario | Go Built-in | Result | safeconv |
|----------|-------------|--------|----------|
| `int64(-1)` → `uint32` | `uint32(-1)` | `4294967295` | `ErrUnderflow` |
| `int64(MaxInt64)` → `int32` | `int32(x)` | `-1` | `ErrOverflow` |
| `float64(NaN)` → `int64` | `int64(x)` | `-9223372036854775808` | `ErrNaN` |
| `float64(1e100)` → `int64` | `int64(x)` | `-9223372036854775808` | `ErrOverflow` |

These silent failures cause security vulnerabilities ([gosec G115](https://github.com/securego/gosec/blob/master/rules/integer.go), [G109](https://github.com/securego/gosec/blob/master/rules/integer.go)), data corruption, and hard-to-debug production issues.

## Quick Start

```go
import "github.com/heisenbergs-uncertainty/safeconv"

// Integer conversion with error handling
func processUserID(id int64) (uint32, error) {
    return safeconv.Int64ToUint32(id)
}

// Float to integer with explicit rounding behavior
func calculateIndex(ratio float64) (int, error) {
    return safeconv.Float64RoundToInt(ratio * 100)
}

// Must variant for constants and initialization
var maxConnections = safeconv.MustInt64ToUint32(10000)
```

## API Overview

safeconv provides **229 conversion functions** covering all Go numeric types:

### Integer Conversions

#### Signed ↔ Unsigned

Convert between signed and unsigned integers with underflow/overflow detection:

```go
// Signed to Unsigned (returns error if negative)
val, err := safeconv.Int64ToUint32(x)
val, err := safeconv.Int32ToUint16(x)
val, err := safeconv.IntToUint(x)

// Unsigned to Signed (returns error if too large)
val, err := safeconv.Uint64ToInt32(x)
val, err := safeconv.Uint32ToInt16(x)
```

#### Narrowing Conversions

Convert to smaller integer types with range checking:

```go
// Signed narrowing
val, err := safeconv.Int64ToInt32(x)
val, err := safeconv.Int32ToInt16(x)
val, err := safeconv.Int16ToInt8(x)

// Unsigned narrowing
val, err := safeconv.Uint64ToUint32(x)
val, err := safeconv.Uint32ToUint16(x)
val, err := safeconv.Uint16ToUint8(x)
```

#### Widening Conversions (Always Safe)

Some conversions can never fail and return values directly (no error):

```go
val := safeconv.Int8ToInt64(x)   // Always succeeds
val := safeconv.Uint16ToUint64(x) // Always succeeds
val := safeconv.Int32ToInt64(x)   // Always succeeds
```

### Float Conversions

#### Float to Integer

Float-to-integer conversions require explicitly choosing between **truncation** or **rounding**:

```go
// Truncate toward zero (like Go's built-in, but safe)
val, err := safeconv.Float64TruncToInt64(3.7)   // → 3
val, err := safeconv.Float64TruncToInt64(-3.7)  // → -3

// Round to nearest (half away from zero)
val, err := safeconv.Float64RoundToInt64(3.5)   // → 4
val, err := safeconv.Float64RoundToInt64(-3.5)  // → -4
val, err := safeconv.Float64RoundToInt64(3.4)   // → 3
```

Both variants handle special float values:
- `NaN` → `ErrNaN`
- `±Inf` → `ErrInfinity`
- Out of range → `ErrOverflow` / `ErrUnderflow`

#### Float Narrowing

```go
// Float64 to Float32 (errors on overflow, precision loss is accepted)
val, err := safeconv.Float64ToFloat32(x)
```

Note: Precision loss is inherent to floating-point and is **not** treated as an error. Only overflow (value too large for float32) returns an error. NaN and ±Inf propagate correctly.

### Must Variants

Every fallible function has a `Must*` variant that panics on error. Use these for:
- Package-level constants
- Test setup
- Cases where failure is a programming error

```go
// Package-level initialization
var (
    MaxConnections = safeconv.MustInt64ToUint32(10000)
    BufferSize     = safeconv.MustFloat64TruncToInt(1024.0)
)

// Test assertions
func TestSomething(t *testing.T) {
    result := safeconv.MustUint64ToInt32(expectedValue)
    // ...
}
```

## Error Handling

### Sentinel Errors

Check error types with `errors.Is`:

```go
val, err := safeconv.Int64ToUint32(x)
if err != nil {
    switch {
    case errors.Is(err, safeconv.ErrOverflow):
        // Value too large for target type
    case errors.Is(err, safeconv.ErrUnderflow):
        // Value too small (e.g., negative to unsigned)
    case errors.Is(err, safeconv.ErrNaN):
        // Float value is NaN
    case errors.Is(err, safeconv.ErrInfinity):
        // Float value is ±Inf
    }
}
```

### Detailed Error Information

Use `errors.As` to get full conversion context:

```go
var convErr *safeconv.ConversionError
if errors.As(err, &convErr) {
    log.Printf("Cannot convert %s(%v) to %s: %v",
        convErr.From,   // "int64"
        convErr.Value,  // -5
        convErr.To,     // "uint32"
        convErr.Err,    // ErrUnderflow
    )
}
```

### Error Types

| Error | Meaning |
|-------|---------|
| `ErrOverflow` | Value exceeds the maximum of the target type |
| `ErrUnderflow` | Value is below the minimum of the target type (includes negative → unsigned) |
| `ErrNaN` | Float value is NaN (not a number) |
| `ErrInfinity` | Float value is positive or negative infinity |

## Complete Function Reference

### Integer: Signed to Unsigned

| From | To | Function | Can Fail? |
|------|-----|----------|-----------|
| `int64` | `uint64` | `Int64ToUint64` | Yes (negative) |
| `int64` | `uint32` | `Int64ToUint32` | Yes |
| `int64` | `uint16` | `Int64ToUint16` | Yes |
| `int64` | `uint8` | `Int64ToUint8` | Yes |
| `int64` | `uint` | `Int64ToUint` | Yes |
| `int32` | `uint64` | `Int32ToUint64` | Yes (negative) |
| `int32` | `uint32` | `Int32ToUint32` | Yes (negative) |
| `int32` | `uint16` | `Int32ToUint16` | Yes |
| `int32` | `uint8` | `Int32ToUint8` | Yes |
| `int32` | `uint` | `Int32ToUint` | Yes (negative) |
| `int16` | `uint64` | `Int16ToUint64` | Yes (negative) |
| `int16` | `uint32` | `Int16ToUint32` | Yes (negative) |
| `int16` | `uint16` | `Int16ToUint16` | Yes (negative) |
| `int16` | `uint8` | `Int16ToUint8` | Yes |
| `int16` | `uint` | `Int16ToUint` | Yes (negative) |
| `int8` | `uint64` | `Int8ToUint64` | Yes (negative) |
| `int8` | `uint32` | `Int8ToUint32` | Yes (negative) |
| `int8` | `uint16` | `Int8ToUint16` | Yes (negative) |
| `int8` | `uint8` | `Int8ToUint8` | Yes (negative) |
| `int8` | `uint` | `Int8ToUint` | Yes (negative) |
| `int` | `uint64` | `IntToUint64` | Yes (negative) |
| `int` | `uint32` | `IntToUint32` | Yes |
| `int` | `uint16` | `IntToUint16` | Yes |
| `int` | `uint8` | `IntToUint8` | Yes |
| `int` | `uint` | `IntToUint` | Yes (negative) |

### Integer: Unsigned to Signed

| From | To | Function | Can Fail? |
|------|-----|----------|-----------|
| `uint64` | `int64` | `Uint64ToInt64` | Yes (overflow) |
| `uint64` | `int32` | `Uint64ToInt32` | Yes |
| `uint64` | `int16` | `Uint64ToInt16` | Yes |
| `uint64` | `int8` | `Uint64ToInt8` | Yes |
| `uint64` | `int` | `Uint64ToInt` | Yes |
| `uint32` | `int64` | `Uint32ToInt64` | **No** |
| `uint32` | `int32` | `Uint32ToInt32` | Yes |
| `uint32` | `int16` | `Uint32ToInt16` | Yes |
| `uint32` | `int8` | `Uint32ToInt8` | Yes |
| `uint32` | `int` | `Uint32ToInt` | **No** |
| `uint16` | `int64` | `Uint16ToInt64` | **No** |
| `uint16` | `int32` | `Uint16ToInt32` | **No** |
| `uint16` | `int16` | `Uint16ToInt16` | Yes |
| `uint16` | `int8` | `Uint16ToInt8` | Yes |
| `uint16` | `int` | `Uint16ToInt` | **No** |
| `uint8` | `int64` | `Uint8ToInt64` | **No** |
| `uint8` | `int32` | `Uint8ToInt32` | **No** |
| `uint8` | `int16` | `Uint8ToInt16` | **No** |
| `uint8` | `int8` | `Uint8ToInt8` | Yes |
| `uint8` | `int` | `Uint8ToInt` | **No** |
| `uint` | `int64` | `UintToInt64` | Yes (on 64-bit) |
| `uint` | `int32` | `UintToInt32` | Yes |
| `uint` | `int16` | `UintToInt16` | Yes |
| `uint` | `int8` | `UintToInt8` | Yes |
| `uint` | `int` | `UintToInt` | Yes |

### Integer: Signed Narrowing

| From | To | Function | Can Fail? |
|------|-----|----------|-----------|
| `int64` | `int32` | `Int64ToInt32` | Yes |
| `int64` | `int16` | `Int64ToInt16` | Yes |
| `int64` | `int8` | `Int64ToInt8` | Yes |
| `int64` | `int` | `Int64ToInt` | Yes (on 32-bit) |
| `int32` | `int64` | `Int32ToInt64` | **No** |
| `int32` | `int16` | `Int32ToInt16` | Yes |
| `int32` | `int8` | `Int32ToInt8` | Yes |
| `int16` | `int64` | `Int16ToInt64` | **No** |
| `int16` | `int32` | `Int16ToInt32` | **No** |
| `int16` | `int8` | `Int16ToInt8` | Yes |
| `int8` | `int64` | `Int8ToInt64` | **No** |
| `int8` | `int32` | `Int8ToInt32` | **No** |
| `int8` | `int16` | `Int8ToInt16` | **No** |
| `int` | `int64` | `IntToInt64` | **No** |
| `int` | `int32` | `IntToInt32` | Yes (on 64-bit) |
| `int` | `int16` | `IntToInt16` | Yes |
| `int` | `int8` | `IntToInt8` | Yes |

### Integer: Unsigned Narrowing

| From | To | Function | Can Fail? |
|------|-----|----------|-----------|
| `uint64` | `uint32` | `Uint64ToUint32` | Yes |
| `uint64` | `uint16` | `Uint64ToUint16` | Yes |
| `uint64` | `uint8` | `Uint64ToUint8` | Yes |
| `uint64` | `uint` | `Uint64ToUint` | Yes (on 32-bit) |
| `uint32` | `uint64` | `Uint32ToUint64` | **No** |
| `uint32` | `uint16` | `Uint32ToUint16` | Yes |
| `uint32` | `uint8` | `Uint32ToUint8` | Yes |
| `uint16` | `uint64` | `Uint16ToUint64` | **No** |
| `uint16` | `uint32` | `Uint16ToUint32` | **No** |
| `uint16` | `uint8` | `Uint16ToUint8` | Yes |
| `uint8` | `uint64` | `Uint8ToUint64` | **No** |
| `uint8` | `uint32` | `Uint8ToUint32` | **No** |
| `uint8` | `uint16` | `Uint8ToUint16` | **No** |
| `uint` | `uint64` | `UintToUint64` | **No** |
| `uint` | `uint32` | `UintToUint32` | Yes (on 64-bit) |
| `uint` | `uint16` | `UintToUint16` | Yes |
| `uint` | `uint8` | `UintToUint8` | Yes |

### Float to Integer (Truncate)

All `Float*TruncTo*` functions truncate toward zero:

| From | To | Function |
|------|----|----------|
| `float64` | `int64` | `Float64TruncToInt64` |
| `float64` | `int32` | `Float64TruncToInt32` |
| `float64` | `int16` | `Float64TruncToInt16` |
| `float64` | `int8` | `Float64TruncToInt8` |
| `float64` | `int` | `Float64TruncToInt` |
| `float64` | `uint64` | `Float64TruncToUint64` |
| `float64` | `uint32` | `Float64TruncToUint32` |
| `float64` | `uint16` | `Float64TruncToUint16` |
| `float64` | `uint8` | `Float64TruncToUint8` |
| `float64` | `uint` | `Float64TruncToUint` |
| `float32` | `int64` | `Float32TruncToInt64` |
| `float32` | `int32` | `Float32TruncToInt32` |
| `float32` | `int16` | `Float32TruncToInt16` |
| `float32` | `int8` | `Float32TruncToInt8` |
| `float32` | `int` | `Float32TruncToInt` |
| `float32` | `uint64` | `Float32TruncToUint64` |
| `float32` | `uint32` | `Float32TruncToUint32` |
| `float32` | `uint16` | `Float32TruncToUint16` |
| `float32` | `uint8` | `Float32TruncToUint8` |
| `float32` | `uint` | `Float32TruncToUint` |

### Float to Integer (Round)

All `Float*RoundTo*` functions round to nearest, with ties going away from zero:

| From | To | Function |
|------|----|----------|
| `float64` | `int64` | `Float64RoundToInt64` |
| `float64` | `int32` | `Float64RoundToInt32` |
| `float64` | `int16` | `Float64RoundToInt16` |
| `float64` | `int8` | `Float64RoundToInt8` |
| `float64` | `int` | `Float64RoundToInt` |
| `float64` | `uint64` | `Float64RoundToUint64` |
| `float64` | `uint32` | `Float64RoundToUint32` |
| `float64` | `uint16` | `Float64RoundToUint16` |
| `float64` | `uint8` | `Float64RoundToUint8` |
| `float64` | `uint` | `Float64RoundToUint` |
| `float32` | `int64` | `Float32RoundToInt64` |
| `float32` | `int32` | `Float32RoundToInt32` |
| `float32` | `int16` | `Float32RoundToInt16` |
| `float32` | `int8` | `Float32RoundToInt8` |
| `float32` | `int` | `Float32RoundToInt` |
| `float32` | `uint64` | `Float32RoundToUint64` |
| `float32` | `uint32` | `Float32RoundToUint32` |
| `float32` | `uint16` | `Float32RoundToUint16` |
| `float32` | `uint8` | `Float32RoundToUint8` |
| `float32` | `uint` | `Float32RoundToUint` |

### Float Narrowing

| From | To | Function | Notes |
|------|----|----------|-------|
| `float64` | `float32` | `Float64ToFloat32` | Errors on overflow; NaN/Inf propagate |

## Design Philosophy

### Explicit Over Implicit

- **No silent truncation**: Every potentially lossy conversion returns an error
- **No magic**: Function names explicitly state source and target types
- **No surprises**: Widening conversions that can't fail return values directly (no error)

### Truncate vs Round

Float-to-integer conversions force you to choose:
- `Trunc` - truncates toward zero (matches Go's built-in behavior)
- `Round` - rounds to nearest, ties away from zero (matches `math.Round`)

This prevents bugs from accidentally choosing the wrong behavior.

### Error Design

Errors support both quick checks and detailed debugging:

```go
// Quick sentinel check
if errors.Is(err, safeconv.ErrOverflow) { ... }

// Detailed debugging
var e *safeconv.ConversionError
if errors.As(err, &e) {
    log.Printf("%s(%v) → %s failed", e.From, e.Value, e.To)
}
```

## Performance

safeconv functions have minimal overhead:
- Single comparison for most conversions
- No allocations on the success path
- Error objects are allocated only on failure

For performance-critical hot paths where you've validated inputs, use Go's built-in conversions. Use safeconv at system boundaries and when processing external data.

## Testing

The library has comprehensive test coverage (97.8%) with tests for:
- Zero values
- Boundary conditions (min/max values)
- Overflow/underflow edge cases
- Special float values (NaN, ±Inf, ±0)
- Platform-specific `int`/`uint` sizes

Run tests:
```bash
go test -v ./...
go test -cover ./...
go test -race ./...
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## License

MIT License - see [LICENSE](LICENSE) for details.

## Acknowledgments

Inspired by:
- [gosec](https://github.com/securego/gosec) rules G115 and G109 for integer overflow detection
- The Go standard library's `strconv` package API design
- Rust's `TryFrom` and `TryInto` traits for safe conversions
