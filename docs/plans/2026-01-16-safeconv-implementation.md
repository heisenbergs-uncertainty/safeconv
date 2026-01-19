# safeconv Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Implement a Go library for safe type conversions that detect overflow, underflow, and precision loss at runtime.

**Architecture:** Single flat package with explicit function names for each type conversion. Generic helpers internally reduce duplication. Hybrid error types (sentinels + detailed struct) enable both easy checking and rich debugging.

**Tech Stack:** Go 1.21+, standard library only (no external dependencies)

---

## Task 1: Update Module Name

**Files:**
- Modify: `go.mod`

**Step 1: Update module path**

Change module name from `github.com/matthew-reed-holden/safe` to `github.com/matthew-reed-holden/safeconv`:

```go
module github.com/matthew-reed-holden/safeconv

go 1.24.11
```

**Step 2: Commit**

```bash
git add go.mod
git commit -m "chore: rename module to safeconv"
```

---

## Task 2: Error Types

**Files:**
- Create: `errors.go`
- Create: `errors_test.go`

**Step 1: Write failing test for ConversionError**

```go
package safeconv

import (
	"errors"
	"testing"
)

func TestConversionError_Error(t *testing.T) {
	err := &ConversionError{
		From:  "int64",
		To:    "uint32",
		Value: int64(-5),
		Err:   ErrUnderflow,
	}

	got := err.Error()
	want := "cannot convert int64(-5) to uint32: value underflows target type"

	if got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestConversionError_Unwrap(t *testing.T) {
	err := &ConversionError{
		From:  "int64",
		To:    "uint32",
		Value: int64(-5),
		Err:   ErrUnderflow,
	}

	if !errors.Is(err, ErrUnderflow) {
		t.Error("errors.Is(err, ErrUnderflow) = false, want true")
	}
}

func TestConversionError_As(t *testing.T) {
	err := &ConversionError{
		From:  "int64",
		To:    "uint32",
		Value: int64(-5),
		Err:   ErrUnderflow,
	}

	var convErr *ConversionError
	if !errors.As(err, &convErr) {
		t.Error("errors.As failed")
	}

	if convErr.From != "int64" {
		t.Errorf("From = %q, want %q", convErr.From, "int64")
	}
	if convErr.To != "uint32" {
		t.Errorf("To = %q, want %q", convErr.To, "uint32")
	}
	if convErr.Value != int64(-5) {
		t.Errorf("Value = %v, want %v", convErr.Value, int64(-5))
	}
}
```

**Step 2: Run test to verify it fails**

Run: `go test -v -run TestConversionError`
Expected: FAIL (types not defined)

**Step 3: Write implementation**

```go
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
```

**Step 4: Run test to verify it passes**

Run: `go test -v -run TestConversionError`
Expected: PASS

**Step 5: Commit**

```bash
git add errors.go errors_test.go
git commit -m "feat: add error types and sentinel errors"
```

---

## Task 3: Type Constraints

**Files:**
- Create: `constraints.go`

**Step 1: Write type constraints**

No test needed - these are compile-time constructs.

```go
package safeconv

// SignedInt is a constraint for signed integer types.
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInt is a constraint for unsigned integer types.
type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer is a constraint for all integer types.
type Integer interface {
	SignedInt | UnsignedInt
}

// Float is a constraint for floating-point types.
type Float interface {
	~float32 | ~float64
}
```

**Step 2: Commit**

```bash
git add constraints.go
git commit -m "feat: add type constraints for generics"
```

---

## Task 4: Internal Helpers

**Files:**
- Create: `convert.go`
- Create: `convert_test.go`

**Step 1: Write failing test for type name helper**

```go
package safeconv

import "testing"

func TestTypeName(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"int", typeName[int](), "int"},
		{"int8", typeName[int8](), "int8"},
		{"int16", typeName[int16](), "int16"},
		{"int32", typeName[int32](), "int32"},
		{"int64", typeName[int64](), "int64"},
		{"uint", typeName[uint](), "uint"},
		{"uint8", typeName[uint8](), "uint8"},
		{"uint16", typeName[uint16](), "uint16"},
		{"uint32", typeName[uint32](), "uint32"},
		{"uint64", typeName[uint64](), "uint64"},
		{"float32", typeName[float32](), "float32"},
		{"float64", typeName[float64](), "float64"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("typeName[%s]() = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}
```

**Step 2: Run test to verify it fails**

Run: `go test -v -run TestTypeName`
Expected: FAIL (typeName not defined)

**Step 3: Write implementation**

```go
package safeconv

import "reflect"

// typeName returns the name of a type for error messages.
func typeName[T any]() string {
	var zero T
	return reflect.TypeOf(zero).Name()
}
```

**Step 4: Run test to verify it passes**

Run: `go test -v -run TestTypeName`
Expected: PASS

**Step 5: Commit**

```bash
git add convert.go convert_test.go
git commit -m "feat: add internal type name helper"
```

---

## Task 5: Signed to Unsigned Conversions (Int64 → Uint*)

**Files:**
- Create: `int_to_uint.go`
- Create: `int_to_uint_test.go`

**Step 1: Write failing tests for Int64ToUint64**

```go
package safeconv

import (
	"errors"
	"math"
	"testing"
)

func TestInt64ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int64", math.MaxInt64, math.MaxInt64, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToUint64(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt64ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint32", math.MaxUint32, math.MaxUint32, nil},
		{"max uint32 + 1", math.MaxUint32 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToUint32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt64ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint16", math.MaxUint16, math.MaxUint16, nil},
		{"max uint16 + 1", math.MaxUint16 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt64ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt64ToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    uint
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToUint(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `go test -v -run "TestInt64ToUint"`
Expected: FAIL (functions not defined)

**Step 3: Write implementation**

```go
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
// Returns ErrUnderflow if negative, ErrOverflow if too large (32-bit systems).
func Int64ToUint(x int64) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int64", To: "uint", Value: x, Err: ErrUnderflow}
	}
	if uint64(x) > uint64(^uint(0)) {
		return 0, &ConversionError{From: "int64", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(x), nil
}
```

**Step 4: Run tests to verify they pass**

Run: `go test -v -run "TestInt64ToUint"`
Expected: PASS

**Step 5: Commit**

```bash
git add int_to_uint.go int_to_uint_test.go
git commit -m "feat: add Int64 to unsigned conversions"
```

---

## Task 6: Remaining Signed to Unsigned Conversions

**Files:**
- Modify: `int_to_uint.go`
- Modify: `int_to_uint_test.go`

**Step 1: Add tests for Int32, Int16, Int8, Int → Uint* conversions**

Add to `int_to_uint_test.go`:

```go
func TestInt32ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt32ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt32ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint16", math.MaxUint16, math.MaxUint16, nil},
		{"overflow", math.MaxUint16 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt32ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"overflow", math.MaxUint8 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt32ToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    uint
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt16ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt16ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt16ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt16ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"overflow", math.MaxUint8 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt16ToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    uint
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt8ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int8
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt8ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int8
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt8ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   int8
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt8ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   int8
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInt8ToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   int8
		want    uint
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIntToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIntToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIntToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint16", math.MaxUint16, math.MaxUint16, nil},
		{"overflow", math.MaxUint16 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIntToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"overflow", math.MaxUint8 + 1, 0, ErrOverflow},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIntToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    uint
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `go test -v -run "TestInt(32|16|8|ToUint)"`
Expected: FAIL

**Step 3: Add implementations to int_to_uint.go**

```go
// Int32ToUint64 converts an int32 to uint64.
func Int32ToUint64(x int32) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Int32ToUint32 converts an int32 to uint32.
func Int32ToUint32(x int32) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	return uint32(x), nil
}

// Int32ToUint16 converts an int32 to uint16.
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
func Int32ToUint(x int32) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int32", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}

// Int16ToUint64 converts an int16 to uint64.
func Int16ToUint64(x int16) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Int16ToUint32 converts an int16 to uint32.
func Int16ToUint32(x int16) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	return uint32(x), nil
}

// Int16ToUint16 converts an int16 to uint16.
func Int16ToUint16(x int16) (uint16, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	return uint16(x), nil
}

// Int16ToUint8 converts an int16 to uint8.
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
func Int16ToUint(x int16) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int16", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}

// Int8ToUint64 converts an int8 to uint64.
func Int8ToUint64(x int8) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Int8ToUint32 converts an int8 to uint32.
func Int8ToUint32(x int8) (uint32, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	return uint32(x), nil
}

// Int8ToUint16 converts an int8 to uint16.
func Int8ToUint16(x int8) (uint16, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	return uint16(x), nil
}

// Int8ToUint8 converts an int8 to uint8.
func Int8ToUint8(x int8) (uint8, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	return uint8(x), nil
}

// Int8ToUint converts an int8 to uint.
func Int8ToUint(x int8) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int8", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}

// IntToUint64 converts an int to uint64.
func IntToUint64(x int) (uint64, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// IntToUint32 converts an int to uint32.
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
func IntToUint(x int) (uint, error) {
	if x < 0 {
		return 0, &ConversionError{From: "int", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}
```

**Step 4: Run all tests**

Run: `go test -v ./...`
Expected: PASS

**Step 5: Commit**

```bash
git add int_to_uint.go int_to_uint_test.go
git commit -m "feat: add all signed to unsigned integer conversions"
```

---

## Task 7: Unsigned to Signed Conversions

**Files:**
- Create: `uint_to_int.go`
- Create: `uint_to_int_test.go`

**Step 1: Write failing tests**

```go
package safeconv

import (
	"errors"
	"math"
	"testing"
)

func TestUint64ToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    int64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int64", math.MaxInt64, math.MaxInt64, nil},
		{"max int64 + 1", math.MaxInt64 + 1, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint64ToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"overflow", math.MaxInt32 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint64ToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    int16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"overflow", math.MaxInt16 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint64ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"overflow", math.MaxInt8 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint64ToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    int
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint32ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint32ToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"overflow", math.MaxInt32 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint32ToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		want    int16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"overflow", math.MaxInt16 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint32ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"overflow", math.MaxInt8 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint32ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32ToInt(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint16ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint16ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  int32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToInt32(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint16ToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		want    int16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"overflow", math.MaxInt16 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint16ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint16ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"overflow", math.MaxInt8 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint16ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint16ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToInt(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint8ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint8ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  int32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt32(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint8ToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  int16
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt16(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint8ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint8
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"overflow", math.MaxInt8 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint8ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUint8ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUintToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    int64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUintToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"overflow", math.MaxInt32 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUintToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    int16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"overflow", math.MaxInt16 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUintToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"overflow", math.MaxInt8 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestUintToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    int
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `go test -v -run "TestUint.*ToInt"`
Expected: FAIL

**Step 3: Write implementation**

```go
package safeconv

import "math"

// Uint64ToInt64 converts a uint64 to int64.
func Uint64ToInt64(x uint64) (int64, error) {
	if x > math.MaxInt64 {
		return 0, &ConversionError{From: "uint64", To: "int64", Value: x, Err: ErrOverflow}
	}
	return int64(x), nil
}

// Uint64ToInt32 converts a uint64 to int32.
func Uint64ToInt32(x uint64) (int32, error) {
	if x > math.MaxInt32 {
		return 0, &ConversionError{From: "uint64", To: "int32", Value: x, Err: ErrOverflow}
	}
	return int32(x), nil
}

// Uint64ToInt16 converts a uint64 to int16.
func Uint64ToInt16(x uint64) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint64", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// Uint64ToInt8 converts a uint64 to int8.
func Uint64ToInt8(x uint64) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint64", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint64ToInt converts a uint64 to int.
func Uint64ToInt(x uint64) (int, error) {
	if x > uint64(math.MaxInt) {
		return 0, &ConversionError{From: "uint64", To: "int", Value: x, Err: ErrOverflow}
	}
	return int(x), nil
}

// Uint32ToInt64 converts a uint32 to int64. Always succeeds.
func Uint32ToInt64(x uint32) int64 {
	return int64(x)
}

// Uint32ToInt32 converts a uint32 to int32.
func Uint32ToInt32(x uint32) (int32, error) {
	if x > math.MaxInt32 {
		return 0, &ConversionError{From: "uint32", To: "int32", Value: x, Err: ErrOverflow}
	}
	return int32(x), nil
}

// Uint32ToInt16 converts a uint32 to int16.
func Uint32ToInt16(x uint32) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint32", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// Uint32ToInt8 converts a uint32 to int8.
func Uint32ToInt8(x uint32) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint32", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint32ToInt converts a uint32 to int. Always succeeds on 64-bit systems.
func Uint32ToInt(x uint32) int {
	return int(x)
}

// Uint16ToInt64 converts a uint16 to int64. Always succeeds.
func Uint16ToInt64(x uint16) int64 {
	return int64(x)
}

// Uint16ToInt32 converts a uint16 to int32. Always succeeds.
func Uint16ToInt32(x uint16) int32 {
	return int32(x)
}

// Uint16ToInt16 converts a uint16 to int16.
func Uint16ToInt16(x uint16) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint16", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// Uint16ToInt8 converts a uint16 to int8.
func Uint16ToInt8(x uint16) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint16", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint16ToInt converts a uint16 to int. Always succeeds.
func Uint16ToInt(x uint16) int {
	return int(x)
}

// Uint8ToInt64 converts a uint8 to int64. Always succeeds.
func Uint8ToInt64(x uint8) int64 {
	return int64(x)
}

// Uint8ToInt32 converts a uint8 to int32. Always succeeds.
func Uint8ToInt32(x uint8) int32 {
	return int32(x)
}

// Uint8ToInt16 converts a uint8 to int16. Always succeeds.
func Uint8ToInt16(x uint8) int16 {
	return int16(x)
}

// Uint8ToInt8 converts a uint8 to int8.
func Uint8ToInt8(x uint8) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint8", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// Uint8ToInt converts a uint8 to int. Always succeeds.
func Uint8ToInt(x uint8) int {
	return int(x)
}

// UintToInt64 converts a uint to int64.
func UintToInt64(x uint) (int64, error) {
	if uint64(x) > math.MaxInt64 {
		return 0, &ConversionError{From: "uint", To: "int64", Value: x, Err: ErrOverflow}
	}
	return int64(x), nil
}

// UintToInt32 converts a uint to int32.
func UintToInt32(x uint) (int32, error) {
	if uint64(x) > math.MaxInt32 {
		return 0, &ConversionError{From: "uint", To: "int32", Value: x, Err: ErrOverflow}
	}
	return int32(x), nil
}

// UintToInt16 converts a uint to int16.
func UintToInt16(x uint) (int16, error) {
	if x > math.MaxInt16 {
		return 0, &ConversionError{From: "uint", To: "int16", Value: x, Err: ErrOverflow}
	}
	return int16(x), nil
}

// UintToInt8 converts a uint to int8.
func UintToInt8(x uint) (int8, error) {
	if x > math.MaxInt8 {
		return 0, &ConversionError{From: "uint", To: "int8", Value: x, Err: ErrOverflow}
	}
	return int8(x), nil
}

// UintToInt converts a uint to int.
func UintToInt(x uint) (int, error) {
	if x > uint(math.MaxInt) {
		return 0, &ConversionError{From: "uint", To: "int", Value: x, Err: ErrOverflow}
	}
	return int(x), nil
}
```

**Step 4: Run tests**

Run: `go test -v ./...`
Expected: PASS

**Step 5: Commit**

```bash
git add uint_to_int.go uint_to_int_test.go
git commit -m "feat: add unsigned to signed integer conversions"
```

---

## Task 8: Integer Narrowing Conversions

**Files:**
- Create: `int_narrow.go`
- Create: `int_narrow_test.go`
- Create: `uint_narrow.go`
- Create: `uint_narrow_test.go`

This task follows the same TDD pattern. Implementation includes:
- Int64 → Int32, Int16, Int8, Int
- Int32 → Int16, Int8
- Int16 → Int8
- Int → Int64, Int32, Int16, Int8
- Uint64 → Uint32, Uint16, Uint8, Uint
- Uint32 → Uint16, Uint8
- Uint16 → Uint8
- Uint → Uint64, Uint32, Uint16, Uint8

(Detailed test and implementation code follows same pattern as Tasks 5-7)

**Commit:**
```bash
git add int_narrow.go int_narrow_test.go uint_narrow.go uint_narrow_test.go
git commit -m "feat: add integer narrowing conversions"
```

---

## Task 9: Float to Integer Truncation

**Files:**
- Create: `float_trunc.go`
- Create: `float_trunc_test.go`

**Step 1: Write failing tests**

```go
package safeconv

import (
	"errors"
	"math"
	"testing"
)

func TestFloat64TruncToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 3.7, 3, nil},
		{"negative", -3.7, -3, nil},
		{"positive round down", 3.2, 3, nil},
		{"large positive", 1e10, 10000000000, nil},
		{"nan", math.NaN(), 0, ErrNaN},
		{"positive inf", math.Inf(1), 0, ErrInfinity},
		{"negative inf", math.Inf(-1), 0, ErrInfinity},
		{"overflow", float64(math.MaxInt64) * 2, 0, ErrOverflow},
		{"underflow", float64(math.MinInt64) * 2, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 3.7, 3, nil},
		{"negative", -3.7, -3, nil},
		{"max int32", float64(math.MaxInt32), math.MaxInt32, nil},
		{"min int32", float64(math.MinInt32), math.MinInt32, nil},
		{"overflow", float64(math.MaxInt32) + 1, 0, ErrOverflow},
		{"underflow", float64(math.MinInt32) - 1, 0, ErrUnderflow},
		{"nan", math.NaN(), 0, ErrNaN},
		{"inf", math.Inf(1), 0, ErrInfinity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 3.7, 3, nil},
		{"negative", -1, 0, ErrUnderflow},
		{"nan", math.NaN(), 0, ErrNaN},
		{"inf", math.Inf(1), 0, ErrInfinity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 3.7, 3, nil},
		{"negative", -3.7, -3, nil},
		{"nan", float32(math.NaN()), 0, ErrNaN},
		{"inf", float32(math.Inf(1)), 0, ErrInfinity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `go test -v -run "TestFloat.*Trunc"`
Expected: FAIL

**Step 3: Write implementation**

```go
package safeconv

import "math"

// Float64TruncToInt64 converts a float64 to int64 by truncating toward zero.
func Float64TruncToInt64(x float64) (int64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt64) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt64) {
		return 0, &ConversionError{From: "float64", To: "int64", Value: x, Err: ErrUnderflow}
	}
	return int64(x), nil
}

// Float64TruncToInt32 converts a float64 to int32 by truncating toward zero.
func Float64TruncToInt32(x float64) (int32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt32) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt32) {
		return 0, &ConversionError{From: "float64", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(x), nil
}

// Float64TruncToInt16 converts a float64 to int16 by truncating toward zero.
func Float64TruncToInt16(x float64) (int16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt16) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt16) {
		return 0, &ConversionError{From: "float64", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(x), nil
}

// Float64TruncToInt8 converts a float64 to int8 by truncating toward zero.
func Float64TruncToInt8(x float64) (int8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt8) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt8) {
		return 0, &ConversionError{From: "float64", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}

// Float64TruncToInt converts a float64 to int by truncating toward zero.
func Float64TruncToInt(x float64) (int, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrInfinity}
	}
	if x > float64(math.MaxInt) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrOverflow}
	}
	if x < float64(math.MinInt) {
		return 0, &ConversionError{From: "float64", To: "int", Value: x, Err: ErrUnderflow}
	}
	return int(x), nil
}

// Float64TruncToUint64 converts a float64 to uint64 by truncating toward zero.
func Float64TruncToUint64(x float64) (uint64, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint64) {
		return 0, &ConversionError{From: "float64", To: "uint64", Value: x, Err: ErrOverflow}
	}
	return uint64(x), nil
}

// Float64TruncToUint32 converts a float64 to uint32 by truncating toward zero.
func Float64TruncToUint32(x float64) (uint32, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint32) {
		return 0, &ConversionError{From: "float64", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// Float64TruncToUint16 converts a float64 to uint16 by truncating toward zero.
func Float64TruncToUint16(x float64) (uint16, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint16) {
		return 0, &ConversionError{From: "float64", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Float64TruncToUint8 converts a float64 to uint8 by truncating toward zero.
func Float64TruncToUint8(x float64) (uint8, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if x > float64(math.MaxUint8) {
		return 0, &ConversionError{From: "float64", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Float64TruncToUint converts a float64 to uint by truncating toward zero.
func Float64TruncToUint(x float64) (uint, error) {
	if math.IsNaN(x) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(x, 0) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrUnderflow}
	}
	if x > float64(^uint(0)) {
		return 0, &ConversionError{From: "float64", To: "uint", Value: x, Err: ErrOverflow}
	}
	return uint(x), nil
}

// Float32TruncToInt64 converts a float32 to int64 by truncating toward zero.
func Float32TruncToInt64(x float32) (int64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int64", Value: x, Err: ErrInfinity}
	}
	return int64(x), nil
}

// Float32TruncToInt32 converts a float32 to int32 by truncating toward zero.
func Float32TruncToInt32(x float32) (int32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrInfinity}
	}
	if x > float32(math.MaxInt32) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrOverflow}
	}
	if x < float32(math.MinInt32) {
		return 0, &ConversionError{From: "float32", To: "int32", Value: x, Err: ErrUnderflow}
	}
	return int32(x), nil
}

// Float32TruncToInt16 converts a float32 to int16 by truncating toward zero.
func Float32TruncToInt16(x float32) (int16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrInfinity}
	}
	if x > float32(math.MaxInt16) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrOverflow}
	}
	if x < float32(math.MinInt16) {
		return 0, &ConversionError{From: "float32", To: "int16", Value: x, Err: ErrUnderflow}
	}
	return int16(x), nil
}

// Float32TruncToInt8 converts a float32 to int8 by truncating toward zero.
func Float32TruncToInt8(x float32) (int8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrInfinity}
	}
	if x > float32(math.MaxInt8) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrOverflow}
	}
	if x < float32(math.MinInt8) {
		return 0, &ConversionError{From: "float32", To: "int8", Value: x, Err: ErrUnderflow}
	}
	return int8(x), nil
}

// Float32TruncToInt converts a float32 to int by truncating toward zero.
func Float32TruncToInt(x float32) (int, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "int", Value: x, Err: ErrInfinity}
	}
	return int(x), nil
}

// Float32TruncToUint64 converts a float32 to uint64 by truncating toward zero.
func Float32TruncToUint64(x float32) (uint64, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint64", Value: x, Err: ErrUnderflow}
	}
	return uint64(x), nil
}

// Float32TruncToUint32 converts a float32 to uint32 by truncating toward zero.
func Float32TruncToUint32(x float32) (uint32, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrUnderflow}
	}
	if x > float32(math.MaxUint32) {
		return 0, &ConversionError{From: "float32", To: "uint32", Value: x, Err: ErrOverflow}
	}
	return uint32(x), nil
}

// Float32TruncToUint16 converts a float32 to uint16 by truncating toward zero.
func Float32TruncToUint16(x float32) (uint16, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrUnderflow}
	}
	if x > float32(math.MaxUint16) {
		return 0, &ConversionError{From: "float32", To: "uint16", Value: x, Err: ErrOverflow}
	}
	return uint16(x), nil
}

// Float32TruncToUint8 converts a float32 to uint8 by truncating toward zero.
func Float32TruncToUint8(x float32) (uint8, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrUnderflow}
	}
	if x > float32(math.MaxUint8) {
		return 0, &ConversionError{From: "float32", To: "uint8", Value: x, Err: ErrOverflow}
	}
	return uint8(x), nil
}

// Float32TruncToUint converts a float32 to uint by truncating toward zero.
func Float32TruncToUint(x float32) (uint, error) {
	if math.IsNaN(float64(x)) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrNaN}
	}
	if math.IsInf(float64(x), 0) {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrInfinity}
	}
	if x < 0 {
		return 0, &ConversionError{From: "float32", To: "uint", Value: x, Err: ErrUnderflow}
	}
	return uint(x), nil
}
```

**Step 4: Run tests**

Run: `go test -v ./...`
Expected: PASS

**Step 5: Commit**

```bash
git add float_trunc.go float_trunc_test.go
git commit -m "feat: add float truncation to integer conversions"
```

---

## Task 10: Float to Integer Rounding

**Files:**
- Create: `float_round.go`
- Create: `float_round_test.go`

Same pattern as Task 9, but using `math.Round()` before conversion.

**Commit:**
```bash
git add float_round.go float_round_test.go
git commit -m "feat: add float rounding to integer conversions"
```

---

## Task 11: Float Narrowing (Float64 → Float32)

**Files:**
- Create: `float_narrow.go`
- Create: `float_narrow_test.go`

**Step 1: Write failing tests**

```go
package safeconv

import (
	"errors"
	"math"
	"testing"
)

func TestFloat64ToFloat32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    float32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 3.14, 3.14, nil},
		{"negative", -3.14, -3.14, nil},
		{"nan", math.NaN(), float32(math.NaN()), nil},
		{"positive inf", math.Inf(1), float32(math.Inf(1)), nil},
		{"negative inf", math.Inf(-1), float32(math.Inf(-1)), nil},
		{"overflow positive", math.MaxFloat64, 0, ErrOverflow},
		{"overflow negative", -math.MaxFloat64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64ToFloat32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if math.IsNaN(float64(tt.want)) {
					if !math.IsNaN(float64(got)) {
						t.Errorf("got %v, want NaN", got)
					}
				} else if got != tt.want {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `go test -v -run TestFloat64ToFloat32`
Expected: FAIL

**Step 3: Write implementation**

```go
package safeconv

import "math"

// Float64ToFloat32 converts a float64 to float32.
// Returns ErrOverflow if the value is too large for float32.
// NaN and Inf values are propagated (not treated as errors).
func Float64ToFloat32(x float64) (float32, error) {
	// NaN and Inf propagate
	if math.IsNaN(x) {
		return float32(math.NaN()), nil
	}
	if math.IsInf(x, 1) {
		return float32(math.Inf(1)), nil
	}
	if math.IsInf(x, -1) {
		return float32(math.Inf(-1)), nil
	}

	// Check for overflow
	if x > math.MaxFloat32 || x < -math.MaxFloat32 {
		return 0, &ConversionError{From: "float64", To: "float32", Value: x, Err: ErrOverflow}
	}

	return float32(x), nil
}
```

**Step 4: Run tests**

Run: `go test -v -run TestFloat64ToFloat32`
Expected: PASS

**Step 5: Commit**

```bash
git add float_narrow.go float_narrow_test.go
git commit -m "feat: add Float64 to Float32 conversion"
```

---

## Task 12: Must Variants

**Files:**
- Create: `must.go`
- Create: `must_test.go`

**Step 1: Write failing tests**

```go
package safeconv

import "testing"

func TestMustInt64ToUint32_Success(t *testing.T) {
	got := MustInt64ToUint32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToUint32(-1)
}

func TestMustFloat64TruncToInt64_Success(t *testing.T) {
	got := MustFloat64TruncToInt64(3.7)
	if got != 3 {
		t.Errorf("got %d, want 3", got)
	}
}

func TestMustFloat64TruncToInt64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToInt64(math.NaN())
}
```

**Step 2: Run tests to verify they fail**

Run: `go test -v -run TestMust`
Expected: FAIL

**Step 3: Write implementation**

```go
package safeconv

// must is an internal helper that panics on error.
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// MustInt64ToUint64 converts int64 to uint64, panicking on error.
func MustInt64ToUint64(x int64) uint64 { return must(Int64ToUint64(x)) }

// MustInt64ToUint32 converts int64 to uint32, panicking on error.
func MustInt64ToUint32(x int64) uint32 { return must(Int64ToUint32(x)) }

// MustInt64ToUint16 converts int64 to uint16, panicking on error.
func MustInt64ToUint16(x int64) uint16 { return must(Int64ToUint16(x)) }

// MustInt64ToUint8 converts int64 to uint8, panicking on error.
func MustInt64ToUint8(x int64) uint8 { return must(Int64ToUint8(x)) }

// MustInt64ToUint converts int64 to uint, panicking on error.
func MustInt64ToUint(x int64) uint { return must(Int64ToUint(x)) }

// ... (continue pattern for all conversion functions that return error)

// MustFloat64TruncToInt64 truncates float64 to int64, panicking on error.
func MustFloat64TruncToInt64(x float64) int64 { return must(Float64TruncToInt64(x)) }

// MustFloat64TruncToInt32 truncates float64 to int32, panicking on error.
func MustFloat64TruncToInt32(x float64) int32 { return must(Float64TruncToInt32(x)) }

// ... (continue for all float truncation functions)

// MustFloat64RoundToInt64 rounds float64 to int64, panicking on error.
func MustFloat64RoundToInt64(x float64) int64 { return must(Float64RoundToInt64(x)) }

// ... (continue for all float rounding functions)

// MustFloat64ToFloat32 converts float64 to float32, panicking on error.
func MustFloat64ToFloat32(x float64) float32 { return must(Float64ToFloat32(x)) }
```

**Step 4: Run tests**

Run: `go test -v ./...`
Expected: PASS

**Step 5: Commit**

```bash
git add must.go must_test.go
git commit -m "feat: add Must variants for all conversions"
```

---

## Task 13: Package Documentation

**Files:**
- Create: `doc.go`

**Step 1: Write package documentation**

```go
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
```

**Step 2: Commit**

```bash
git add doc.go
git commit -m "docs: add package documentation"
```

---

## Task 14: Final Verification

**Step 1: Run all tests**

Run: `go test -v ./...`
Expected: All tests pass

**Step 2: Run go vet**

Run: `go vet ./...`
Expected: No issues

**Step 3: Check test coverage**

Run: `go test -cover ./...`
Expected: >80% coverage

**Step 4: Verify godoc renders correctly**

Run: `go doc -all .`
Expected: Clean documentation output

---

## Summary

| Task | Description | Files |
|------|-------------|-------|
| 1 | Update module name | go.mod |
| 2 | Error types | errors.go, errors_test.go |
| 3 | Type constraints | constraints.go |
| 4 | Internal helpers | convert.go, convert_test.go |
| 5 | Int64 → Uint* | int_to_uint.go, int_to_uint_test.go |
| 6 | Remaining signed → unsigned | int_to_uint.go, int_to_uint_test.go |
| 7 | Unsigned → signed | uint_to_int.go, uint_to_int_test.go |
| 8 | Integer narrowing | int_narrow.go, uint_narrow.go, tests |
| 9 | Float truncation | float_trunc.go, float_trunc_test.go |
| 10 | Float rounding | float_round.go, float_round_test.go |
| 11 | Float narrowing | float_narrow.go, float_narrow_test.go |
| 12 | Must variants | must.go, must_test.go |
| 13 | Package docs | doc.go |
| 14 | Final verification | - |
