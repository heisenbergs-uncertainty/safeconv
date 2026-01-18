# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/heisenbergs-uncertainty/safeconv/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/heisenbergs-uncertainty/safeconv/releases/tag/v0.1.0
