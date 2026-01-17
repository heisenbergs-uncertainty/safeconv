package safeconv

import (
	"math"
	"testing"
)

// =============================================================================
// Int to Uint Must* tests
// =============================================================================

func TestMustInt64ToUint64_Success(t *testing.T) {
	got := MustInt64ToUint64(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToUint64(-1) // Should panic
}

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
	MustInt64ToUint32(-1) // Should panic
}

func TestMustInt64ToUint16_Success(t *testing.T) {
	got := MustInt64ToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToUint16(-1) // Should panic
}

func TestMustInt64ToUint8_Success(t *testing.T) {
	got := MustInt64ToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToUint8(-1) // Should panic
}

func TestMustInt64ToUint_Success(t *testing.T) {
	got := MustInt64ToUint(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToUint(-1) // Should panic
}

func TestMustInt32ToUint64_Success(t *testing.T) {
	got := MustInt32ToUint64(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt32ToUint64(-1) // Should panic
}

func TestMustInt32ToUint32_Success(t *testing.T) {
	got := MustInt32ToUint32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt32ToUint32(-1) // Should panic
}

func TestMustInt32ToUint16_Success(t *testing.T) {
	got := MustInt32ToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt32ToUint16(-1) // Should panic
}

func TestMustInt32ToUint8_Success(t *testing.T) {
	got := MustInt32ToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt32ToUint8(-1) // Should panic
}

func TestMustInt32ToUint_Success(t *testing.T) {
	got := MustInt32ToUint(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt32ToUint(-1) // Should panic
}

func TestMustInt16ToUint64_Success(t *testing.T) {
	got := MustInt16ToUint64(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt16ToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt16ToUint64(-1) // Should panic
}

func TestMustInt16ToUint32_Success(t *testing.T) {
	got := MustInt16ToUint32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt16ToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt16ToUint32(-1) // Should panic
}

func TestMustInt16ToUint16_Success(t *testing.T) {
	got := MustInt16ToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt16ToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt16ToUint16(-1) // Should panic
}

func TestMustInt16ToUint8_Success(t *testing.T) {
	got := MustInt16ToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt16ToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt16ToUint8(-1) // Should panic
}

func TestMustInt16ToUint_Success(t *testing.T) {
	got := MustInt16ToUint(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt16ToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt16ToUint(-1) // Should panic
}

func TestMustInt8ToUint64_Success(t *testing.T) {
	got := MustInt8ToUint64(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt8ToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt8ToUint64(-1) // Should panic
}

func TestMustInt8ToUint32_Success(t *testing.T) {
	got := MustInt8ToUint32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt8ToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt8ToUint32(-1) // Should panic
}

func TestMustInt8ToUint16_Success(t *testing.T) {
	got := MustInt8ToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt8ToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt8ToUint16(-1) // Should panic
}

func TestMustInt8ToUint8_Success(t *testing.T) {
	got := MustInt8ToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt8ToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt8ToUint8(-1) // Should panic
}

func TestMustInt8ToUint_Success(t *testing.T) {
	got := MustInt8ToUint(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt8ToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt8ToUint(-1) // Should panic
}

func TestMustIntToUint64_Success(t *testing.T) {
	got := MustIntToUint64(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToUint64(-1) // Should panic
}

func TestMustIntToUint32_Success(t *testing.T) {
	got := MustIntToUint32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToUint32(-1) // Should panic
}

func TestMustIntToUint16_Success(t *testing.T) {
	got := MustIntToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToUint16(-1) // Should panic
}

func TestMustIntToUint8_Success(t *testing.T) {
	got := MustIntToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToUint8(-1) // Should panic
}

func TestMustIntToUint_Success(t *testing.T) {
	got := MustIntToUint(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToUint(-1) // Should panic
}

// =============================================================================
// Uint to Int Must* tests
// =============================================================================

func TestMustUint64ToInt64_Success(t *testing.T) {
	got := MustUint64ToInt64(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToInt64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToInt64(math.MaxUint64) // Should panic (exceeds MaxInt64)
}

func TestMustUint64ToInt32_Success(t *testing.T) {
	got := MustUint64ToInt32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToInt32(math.MaxInt32 + 1) // Should panic
}

func TestMustUint64ToInt16_Success(t *testing.T) {
	got := MustUint64ToInt16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToInt16(math.MaxInt16 + 1) // Should panic
}

func TestMustUint64ToInt8_Success(t *testing.T) {
	got := MustUint64ToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustUint64ToInt_Success(t *testing.T) {
	got := MustUint64ToInt(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToInt_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToInt(math.MaxUint64) // Should panic
}

func TestMustUint32ToInt32_Success(t *testing.T) {
	got := MustUint32ToInt32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint32ToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint32ToInt32(math.MaxInt32 + 1) // Should panic
}

func TestMustUint32ToInt16_Success(t *testing.T) {
	got := MustUint32ToInt16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint32ToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint32ToInt16(math.MaxInt16 + 1) // Should panic
}

func TestMustUint32ToInt8_Success(t *testing.T) {
	got := MustUint32ToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint32ToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint32ToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustUint16ToInt16_Success(t *testing.T) {
	got := MustUint16ToInt16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint16ToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint16ToInt16(math.MaxInt16 + 1) // Should panic
}

func TestMustUint16ToInt8_Success(t *testing.T) {
	got := MustUint16ToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint16ToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint16ToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustUint8ToInt8_Success(t *testing.T) {
	got := MustUint8ToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint8ToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint8ToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustUintToInt64_Success(t *testing.T) {
	got := MustUintToInt64(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToInt32_Success(t *testing.T) {
	got := MustUintToInt32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUintToInt32(math.MaxInt32 + 1) // Should panic
}

func TestMustUintToInt16_Success(t *testing.T) {
	got := MustUintToInt16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUintToInt16(math.MaxInt16 + 1) // Should panic
}

func TestMustUintToInt8_Success(t *testing.T) {
	got := MustUintToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUintToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustUintToInt_Success(t *testing.T) {
	got := MustUintToInt(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToInt_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUintToInt(uint(math.MaxInt) + 1) // Should panic
}

// =============================================================================
// Int narrowing Must* tests
// =============================================================================

func TestMustInt64ToInt32_Success(t *testing.T) {
	got := MustInt64ToInt32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToInt32(math.MaxInt32 + 1) // Should panic
}

func TestMustInt64ToInt16_Success(t *testing.T) {
	got := MustInt64ToInt16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToInt16(math.MaxInt16 + 1) // Should panic
}

func TestMustInt64ToInt8_Success(t *testing.T) {
	got := MustInt64ToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt64ToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt64ToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustInt64ToInt_Success(t *testing.T) {
	got := MustInt64ToInt(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToInt16_Success(t *testing.T) {
	got := MustInt32ToInt16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt32ToInt16(math.MaxInt16 + 1) // Should panic
}

func TestMustInt32ToInt8_Success(t *testing.T) {
	got := MustInt32ToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt32ToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt32ToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustInt16ToInt8_Success(t *testing.T) {
	got := MustInt16ToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustInt16ToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustInt16ToInt8(math.MaxInt8 + 1) // Should panic
}

func TestMustIntToInt32_Success(t *testing.T) {
	got := MustIntToInt32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToInt32(math.MaxInt32 + 1) // Should panic
}

func TestMustIntToInt16_Success(t *testing.T) {
	got := MustIntToInt16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToInt16(math.MaxInt16 + 1) // Should panic
}

func TestMustIntToInt8_Success(t *testing.T) {
	got := MustIntToInt8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustIntToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustIntToInt8(math.MaxInt8 + 1) // Should panic
}

// =============================================================================
// Uint narrowing Must* tests
// =============================================================================

func TestMustUint64ToUint32_Success(t *testing.T) {
	got := MustUint64ToUint32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToUint32(math.MaxUint32 + 1) // Should panic
}

func TestMustUint64ToUint16_Success(t *testing.T) {
	got := MustUint64ToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToUint16(math.MaxUint16 + 1) // Should panic
}

func TestMustUint64ToUint8_Success(t *testing.T) {
	got := MustUint64ToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint64ToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint64ToUint8(math.MaxUint8 + 1) // Should panic
}

func TestMustUint64ToUint_Success(t *testing.T) {
	got := MustUint64ToUint(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint32ToUint16_Success(t *testing.T) {
	got := MustUint32ToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint32ToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint32ToUint16(math.MaxUint16 + 1) // Should panic
}

func TestMustUint32ToUint8_Success(t *testing.T) {
	got := MustUint32ToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint32ToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint32ToUint8(math.MaxUint8 + 1) // Should panic
}

func TestMustUint16ToUint8_Success(t *testing.T) {
	got := MustUint16ToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUint16ToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUint16ToUint8(math.MaxUint8 + 1) // Should panic
}

func TestMustUintToUint32_Success(t *testing.T) {
	got := MustUintToUint32(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToUint16_Success(t *testing.T) {
	got := MustUintToUint16(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUintToUint16(math.MaxUint16 + 1) // Should panic
}

func TestMustUintToUint8_Success(t *testing.T) {
	got := MustUintToUint8(100)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustUintToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustUintToUint8(math.MaxUint8 + 1) // Should panic
}

// =============================================================================
// Float truncation Must* tests
// =============================================================================

func TestMustFloat64TruncToInt64_Success(t *testing.T) {
	got := MustFloat64TruncToInt64(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToInt64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToInt64(math.NaN()) // Should panic
}

func TestMustFloat64TruncToInt32_Success(t *testing.T) {
	got := MustFloat64TruncToInt32(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToInt32(math.Inf(1)) // Should panic
}

func TestMustFloat64TruncToInt16_Success(t *testing.T) {
	got := MustFloat64TruncToInt16(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToInt16(float64(math.MaxInt16 + 1)) // Should panic
}

func TestMustFloat64TruncToInt8_Success(t *testing.T) {
	got := MustFloat64TruncToInt8(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToInt8(float64(math.MaxInt8 + 1)) // Should panic
}

func TestMustFloat64TruncToInt_Success(t *testing.T) {
	got := MustFloat64TruncToInt(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToInt_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToInt(math.NaN()) // Should panic
}

func TestMustFloat64TruncToUint64_Success(t *testing.T) {
	got := MustFloat64TruncToUint64(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToUint64(-1.0) // Should panic
}

func TestMustFloat64TruncToUint32_Success(t *testing.T) {
	got := MustFloat64TruncToUint32(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToUint32(-1.0) // Should panic
}

func TestMustFloat64TruncToUint16_Success(t *testing.T) {
	got := MustFloat64TruncToUint16(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToUint16(-1.0) // Should panic
}

func TestMustFloat64TruncToUint8_Success(t *testing.T) {
	got := MustFloat64TruncToUint8(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToUint8(-1.0) // Should panic
}

func TestMustFloat64TruncToUint_Success(t *testing.T) {
	got := MustFloat64TruncToUint(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64TruncToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64TruncToUint(-1.0) // Should panic
}

// Float32 truncation tests

func TestMustFloat32TruncToInt64_Success(t *testing.T) {
	got := MustFloat32TruncToInt64(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToInt64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToInt64(float32(math.NaN())) // Should panic
}

func TestMustFloat32TruncToInt32_Success(t *testing.T) {
	got := MustFloat32TruncToInt32(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToInt32(float32(math.Inf(1))) // Should panic
}

func TestMustFloat32TruncToInt16_Success(t *testing.T) {
	got := MustFloat32TruncToInt16(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToInt16(float32(math.MaxInt16 + 1)) // Should panic
}

func TestMustFloat32TruncToInt8_Success(t *testing.T) {
	got := MustFloat32TruncToInt8(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToInt8(float32(math.MaxInt8 + 1)) // Should panic
}

func TestMustFloat32TruncToInt_Success(t *testing.T) {
	got := MustFloat32TruncToInt(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToInt_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToInt(float32(math.NaN())) // Should panic
}

func TestMustFloat32TruncToUint64_Success(t *testing.T) {
	got := MustFloat32TruncToUint64(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToUint64(-1.0) // Should panic
}

func TestMustFloat32TruncToUint32_Success(t *testing.T) {
	got := MustFloat32TruncToUint32(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToUint32(-1.0) // Should panic
}

func TestMustFloat32TruncToUint16_Success(t *testing.T) {
	got := MustFloat32TruncToUint16(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToUint16(-1.0) // Should panic
}

func TestMustFloat32TruncToUint8_Success(t *testing.T) {
	got := MustFloat32TruncToUint8(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToUint8(-1.0) // Should panic
}

func TestMustFloat32TruncToUint_Success(t *testing.T) {
	got := MustFloat32TruncToUint(100.9)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32TruncToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32TruncToUint(-1.0) // Should panic
}

// =============================================================================
// Float rounding Must* tests
// =============================================================================

func TestMustFloat64RoundToInt64_Success(t *testing.T) {
	got := MustFloat64RoundToInt64(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToInt64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToInt64(math.NaN()) // Should panic
}

func TestMustFloat64RoundToInt32_Success(t *testing.T) {
	got := MustFloat64RoundToInt32(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToInt32(math.Inf(1)) // Should panic
}

func TestMustFloat64RoundToInt16_Success(t *testing.T) {
	got := MustFloat64RoundToInt16(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToInt16(float64(math.MaxInt16 + 1)) // Should panic
}

func TestMustFloat64RoundToInt8_Success(t *testing.T) {
	got := MustFloat64RoundToInt8(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToInt8(float64(math.MaxInt8 + 1)) // Should panic
}

func TestMustFloat64RoundToInt_Success(t *testing.T) {
	got := MustFloat64RoundToInt(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToInt_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToInt(math.NaN()) // Should panic
}

func TestMustFloat64RoundToUint64_Success(t *testing.T) {
	got := MustFloat64RoundToUint64(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToUint64(-1.0) // Should panic
}

func TestMustFloat64RoundToUint32_Success(t *testing.T) {
	got := MustFloat64RoundToUint32(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToUint32(-1.0) // Should panic
}

func TestMustFloat64RoundToUint16_Success(t *testing.T) {
	got := MustFloat64RoundToUint16(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToUint16(-1.0) // Should panic
}

func TestMustFloat64RoundToUint8_Success(t *testing.T) {
	got := MustFloat64RoundToUint8(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToUint8(-1.0) // Should panic
}

func TestMustFloat64RoundToUint_Success(t *testing.T) {
	got := MustFloat64RoundToUint(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat64RoundToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64RoundToUint(-1.0) // Should panic
}

// Float32 rounding tests

func TestMustFloat32RoundToInt64_Success(t *testing.T) {
	got := MustFloat32RoundToInt64(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToInt64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToInt64(float32(math.NaN())) // Should panic
}

func TestMustFloat32RoundToInt32_Success(t *testing.T) {
	got := MustFloat32RoundToInt32(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToInt32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToInt32(float32(math.Inf(1))) // Should panic
}

func TestMustFloat32RoundToInt16_Success(t *testing.T) {
	got := MustFloat32RoundToInt16(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToInt16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToInt16(float32(math.MaxInt16 + 1)) // Should panic
}

func TestMustFloat32RoundToInt8_Success(t *testing.T) {
	got := MustFloat32RoundToInt8(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToInt8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToInt8(float32(math.MaxInt8 + 1)) // Should panic
}

func TestMustFloat32RoundToInt_Success(t *testing.T) {
	got := MustFloat32RoundToInt(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToInt_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToInt(float32(math.NaN())) // Should panic
}

func TestMustFloat32RoundToUint64_Success(t *testing.T) {
	got := MustFloat32RoundToUint64(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToUint64_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToUint64(-1.0) // Should panic
}

func TestMustFloat32RoundToUint32_Success(t *testing.T) {
	got := MustFloat32RoundToUint32(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToUint32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToUint32(-1.0) // Should panic
}

func TestMustFloat32RoundToUint16_Success(t *testing.T) {
	got := MustFloat32RoundToUint16(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToUint16_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToUint16(-1.0) // Should panic
}

func TestMustFloat32RoundToUint8_Success(t *testing.T) {
	got := MustFloat32RoundToUint8(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToUint8_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToUint8(-1.0) // Should panic
}

func TestMustFloat32RoundToUint_Success(t *testing.T) {
	got := MustFloat32RoundToUint(100.4)
	if got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

func TestMustFloat32RoundToUint_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat32RoundToUint(-1.0) // Should panic
}

// =============================================================================
// Float narrowing Must* tests
// =============================================================================

func TestMustFloat64ToFloat32_Success(t *testing.T) {
	got := MustFloat64ToFloat32(100.5)
	if got != 100.5 {
		t.Errorf("got %f, want 100.5", got)
	}
}

func TestMustFloat64ToFloat32_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	MustFloat64ToFloat32(math.MaxFloat64) // Should panic (overflow)
}
