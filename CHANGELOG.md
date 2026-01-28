# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2026-01-28

### Added

- **Rune conversions** with Unicode validation:
  - Integer to rune: `Int64ToRune`, `Int32ToRune`, `IntToRune`, `Int16ToRune`, `Int8ToRune`
  - Unsigned to rune: `Uint64ToRune`, `Uint32ToRune`, `UintToRune`, `Uint16ToRune`, `Uint8ToRune`
  - Rune to integer: `RuneToInt64`, `RuneToInt32`, `RuneToInt`
  - Rune to unsigned: `RuneToUint64`, `RuneToUint32`, `RuneToUint`, `RuneToUint16`, `RuneToUint8`
  - Must variants for all fallible rune functions
  - Validates Unicode range (0x0 to 0x10FFFF, excluding surrogates 0xD800-0xDFFF)

- **New error type**: `ErrInvalidUnicode` for invalid Unicode code points
- **New convenience method**: `IsInvalidUnicode()` on `ConversionError`
- **Helper function**: `IsValidRune()` for quick Unicode validation

### Notes

- `byte` is an alias for `uint8` - use `*Uint8*` functions for byte conversions
- `rune` is an alias for `int32` but rune functions add Unicode validation

## [0.2.0] - 2026-01-26

### Added

- **Error convenience methods** on `ConversionError`:
  - `IsOverflow()` - returns true if error is due to overflow
  - `IsUnderflow()` - returns true if error is due to underflow
  - `IsNaN()` - returns true if error is due to NaN
  - `IsInfinity()` - returns true if error is due to infinity

- **Clamp variants** for all conversion types:
  - Signed to unsigned: `ClampInt64ToUint32`, etc. (25 functions)
  - Unsigned to signed: `ClampUint64ToInt32`, etc. (17 functions)
  - Signed narrowing: `ClampInt64ToInt32`, etc. (10 functions)
  - Unsigned narrowing: `ClampUint64ToUint32`, etc. (10 functions)
  - Float truncation: `ClampFloat64TruncToInt64`, etc. (20 functions)
  - Float rounding: `ClampFloat64RoundToInt64`, etc. (20 functions)
  - Clamp functions return clamped values for overflow/underflow
  - Float clamp functions still return errors for NaN and Infinity

## [0.1.1] - 2026-01-22

### Fixed

- Correct module path in release tags so `go get github.com/matthew-reed-holden/safeconv` resolves without mismatched module errors.

## [0.1.0] - 2026-01-17

### Added

- **Integer conversions** with overflow/underflow detection:
  - Signed to unsigned: `Int64ToUint32`, `Int32ToUint16`, etc. (25 functions)
  - Unsigned to signed: `Uint64ToInt32`, `Uint32ToInt16`, etc. (25 functions)
  - Signed narrowing: `Int64ToInt32`, `Int32ToInt16`, etc.
  - Unsigned narrowing: `Uint64ToUint32`, `Uint32ToUint16`, etc.
  - Widening conversions that always succeed return values directly (no error)

- **Float-to-integer conversions** with explicit rounding behavior:
  - Truncation: `Float64TruncToInt64`, `Float32TruncToInt32`, etc. (20 functions)
  - Rounding: `Float64RoundToInt64`, `Float32RoundToInt32`, etc. (20 functions)
  - Proper handling of NaN (`ErrNaN`) and infinity (`ErrInfinity`)

- **Float narrowing**: `Float64ToFloat32` with overflow detection

- **Must variants** for all fallible functions (91 functions)
  - Panic on error for use in initialization and tests

- **Error types**:
  - Sentinel errors: `ErrOverflow`, `ErrUnderflow`, `ErrNaN`, `ErrInfinity`
  - Detailed `ConversionError` type with `From`, `To`, `Value`, and `Err` fields
  - Compatible with `errors.Is()` and `errors.As()`

- Comprehensive test coverage (97.8%)
- Package documentation with examples

[Unreleased]: https://github.com/matthew-reed-holden/safeconv/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/matthew-reed-holden/safeconv/releases/tag/v0.3.0
[0.2.0]: https://github.com/matthew-reed-holden/safeconv/releases/tag/v0.2.0
[0.1.1]: https://github.com/matthew-reed-holden/safeconv/releases/tag/v0.1.1
[0.1.0]: https://github.com/matthew-reed-holden/safeconv/releases/tag/v0.1.0
