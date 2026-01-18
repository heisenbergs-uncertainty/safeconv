# safeconv Roadmap

**Date:** 2026-01-18
**Status:** Approved

## Overview

This roadmap outlines the planned evolution of safeconv from v0.1.0 to v1.0.0 and beyond. The focus is on three pillars:

1. **API Refinement** - Improve error handling and add Clamp variants
2. **Type Coverage** - Extend to rune with Unicode validation
3. **Ecosystem** - Safe types for JSON marshaling with extensible foundation

## Design Principles

- **Explicit over implicit** - Follow strconv pattern with explicit function names
- **Zero dependencies** - Core library uses stdlib only; ecosystem integrations live in separate repos
- **Additive changes** - v0.x allows API iteration; v1.0.0 commits to stability

## Version Milestones

### v0.2.0 - API Refinement

**Focus:** Improve error API and add Clamp variants

**Error Convenience Methods:**
```go
func (e *ConversionError) IsOverflow() bool
func (e *ConversionError) IsUnderflow() bool
func (e *ConversionError) IsNaN() bool
func (e *ConversionError) IsInfinity() bool
```

**Clamp Variants:**
- Add `Clamp*` functions for integer conversions where clamping makes sense
- Clamp numeric overflow/underflow to min/max of target type
- NaN and Infinity still return errors (not silently converted)
- Example: `ClampInt64ToUint32(-5)` → `0`, `ClampInt64ToUint32(1e10)` → `math.MaxUint32`

**Scope:**
- Signed to unsigned clamping
- Unsigned to signed clamping
- Integer narrowing clamping
- Must variants for all new Clamp functions

---

### v0.3.0 - Type Coverage

**Focus:** Rune support with Unicode validation

**Rune Conversions:**
```go
func Int64ToRune(x int64) (rune, error)
func Int32ToRune(x int32) (rune, error)
func IntToRune(x int) (rune, error)
// ... and reverse: RuneToInt64, RuneToInt32, etc.
```

**Unicode Validation:**
- Valid range: 0x0 to 0x10FFFF
- Reject surrogate code points: 0xD800 to 0xDFFF
- New sentinel error: `ErrInvalidUnicode`

**Byte Handling:**
- No explicit `*Byte` functions (byte = uint8 is well-known)
- Document alias relationship in README

**Not Included:**
- `time.Duration` conversions (removed from scope; revisit if users request)

---

### v0.4.0 - Safe Types (JSON)

**Focus:** Extensible safe types with JSON marshaling support

**Hybrid Generic Approach:**
```go
// Generic core type
type Safe[T Numeric] struct {
    value T
}

// Explicit aliases for clean API
type Uint32 = Safe[uint32]
type Int64 = Safe[int64]
type Float64 = Safe[float64]
// ... etc for all numeric types
```

**Interfaces Implemented:**
- `json.Unmarshaler` - Safe unmarshaling with overflow detection
- `json.Marshaler` - Standard marshaling
- `fmt.Stringer` - String representation

**Methods:**
```go
func (s Safe[T]) Value() T
func (s *Safe[T]) Set(v T)
func (s *Safe[T]) From(v any) error  // Dynamic conversion
```

**Design Goals:**
- Foundation for downstream packages (e.g., `safeconv-pgtype`)
- Zero external dependencies
- Composable with existing code

---

### v0.5.0 - Research & Exploration

**Focus:** Explore additional integrations based on user feedback

**Research Items:**
- `database/sql.Scanner` support for safe types
- Struct tag validation approach (e.g., `safeconv:"check"` tags)
- Evaluate feedback from v0.2–v0.4 releases

**Outcome:** Decide what to include in v1.0.0 stable release

---

### v1.0.0 - Stable Release

**Focus:** API stability commitment

**Requirements:**
- Finalize public API based on v0.x learnings
- Comprehensive documentation and examples
- Migration guide from v0.x if breaking changes needed
- All planned features implemented and tested

**Stability Guarantee:**
- No breaking changes in v1.x releases
- New features added in minor versions
- Bug fixes in patch versions

---

## Future (Post v1.0.0)

**Potential additions based on user demand:**
- `time.Duration` conversions (if overflow detection proves valuable)
- Separate ecosystem packages:
  - `safeconv-pgtype` - PostgreSQL type integration
  - `safeconv-gin` - Gin framework integration
  - Others based on community interest
- Outcomes from v0.5.0 research

---

## Out of Scope

The following are explicitly not planned:

- **Generic public API** - `Convert[From, To](x)` stays internal only
- **TryFrom/TryInto interfaces** - Not idiomatic Go
- **Parsing/formatting** - Use `strconv` for that
- **External dependencies in core** - Ecosystem integrations live in separate repos

---

## Decision Log

| Decision | Rationale |
|----------|-----------|
| Drop Duration | Overflow at 292+ years is validation, not conversion |
| Drop TryFrom/TryInto | Doesn't fit Go idioms |
| Skip byte aliases | `byte` = `uint8` is fundamental Go knowledge |
| Add rune with Unicode validation | Semantic value beyond int32 range check |
| Hybrid Safe[T] approach | Clean API (`Uint32`) + single implementation + extensibility |
| Zero deps in core | Ecosystem packages can add deps separately |
| Struct tags as research | Explore but don't commit yet |
