package safeconv

import (
	"errors"
	"math"
	"testing"
)

// =============================================================================
// Tests for Uint64 to Int* conversions
// =============================================================================

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
				t.Errorf("Uint64ToInt64(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int32 + 1", math.MaxInt32 + 1, 0, ErrOverflow},
		{"max uint32", math.MaxUint32, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToInt32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16 + 1", math.MaxInt16 + 1, 0, ErrOverflow},
		{"max uint16", math.MaxUint16, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToInt16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"max uint8", math.MaxUint8, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int", uint64(math.MaxInt), math.MaxInt, nil},
		{"max int + 1", uint64(math.MaxInt) + 1, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToInt(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Tests for Uint32 to Int* conversions
// =============================================================================

func TestUint32ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"max uint32", math.MaxUint32, math.MaxUint32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("Uint32ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int32 + 1", math.MaxInt32 + 1, 0, ErrOverflow},
		{"max uint32", math.MaxUint32, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint32ToInt32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16 + 1", math.MaxInt16 + 1, 0, ErrOverflow},
		{"max uint16", math.MaxUint16, 0, ErrOverflow},
		{"max uint32", math.MaxUint32, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint32ToInt16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"max uint8", math.MaxUint8, 0, ErrOverflow},
		{"max uint32", math.MaxUint32, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint32ToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"max uint32", math.MaxUint32, math.MaxUint32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32ToInt(tt.input)
			if got != tt.want {
				t.Errorf("Uint32ToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Tests for Uint16 to Int* conversions
// =============================================================================

func TestUint16ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("Uint16ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToInt32(tt.input)
			if got != tt.want {
				t.Errorf("Uint16ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16 + 1", math.MaxInt16 + 1, 0, ErrOverflow},
		{"max uint16", math.MaxUint16, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint16ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint16ToInt16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint16ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"max uint8", math.MaxUint8, 0, ErrOverflow},
		{"max uint16", math.MaxUint16, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint16ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint16ToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint16ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToInt(tt.input)
			if got != tt.want {
				t.Errorf("Uint16ToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Tests for Uint8 to Int* conversions
// =============================================================================

func TestUint8ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt32(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt16(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"max uint8", math.MaxUint8, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint8ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint8ToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint8ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToInt(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Tests for Uint to Int* conversions
// =============================================================================

func TestUintToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    int64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int64", uint(math.MaxInt64), math.MaxInt64, nil},
		// On 64-bit systems, uint can exceed MaxInt64
		{"max int64 + 1", uint(math.MaxInt64) + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToInt64(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToInt64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int32 + 1", math.MaxInt32 + 1, 0, ErrOverflow},
		{"max uint32", math.MaxUint32, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToInt32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToInt32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16 + 1", math.MaxInt16 + 1, 0, ErrOverflow},
		{"max uint16", math.MaxUint16, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToInt16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToInt16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"max uint8", math.MaxUint8, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToInt8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int", uint(math.MaxInt), math.MaxInt, nil},
		{"max int + 1", uint(math.MaxInt) + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToInt(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Tests for ConversionError details
// =============================================================================

func TestUintToIntConversionErrors(t *testing.T) {
	t.Run("Uint64ToInt64 overflow error details", func(t *testing.T) {
		input := uint64(math.MaxInt64 + 1)
		_, err := Uint64ToInt64(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "int64" {
			t.Errorf("To = %q, want %q", convErr.To, "int64")
		}
		if convErr.Value != input {
			t.Errorf("Value = %v, want %v", convErr.Value, input)
		}
		if !errors.Is(convErr.Err, ErrOverflow) {
			t.Errorf("Err = %v, want %v", convErr.Err, ErrOverflow)
		}
	})

	t.Run("Uint64ToInt32 overflow error details", func(t *testing.T) {
		input := uint64(math.MaxInt32 + 1)
		_, err := Uint64ToInt32(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "int32" {
			t.Errorf("To = %q, want %q", convErr.To, "int32")
		}
	})

	t.Run("Uint64ToInt16 overflow error details", func(t *testing.T) {
		input := uint64(math.MaxInt16 + 1)
		_, err := Uint64ToInt16(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
	})

	t.Run("Uint64ToInt8 overflow error details", func(t *testing.T) {
		input := uint64(math.MaxInt8 + 1)
		_, err := Uint64ToInt8(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("Uint64ToInt overflow error details", func(t *testing.T) {
		input := uint64(math.MaxInt) + 1
		_, err := Uint64ToInt(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "int" {
			t.Errorf("To = %q, want %q", convErr.To, "int")
		}
	})

	t.Run("Uint32ToInt32 overflow error details", func(t *testing.T) {
		input := uint32(math.MaxInt32 + 1)
		_, err := Uint32ToInt32(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint32" {
			t.Errorf("From = %q, want %q", convErr.From, "uint32")
		}
		if convErr.To != "int32" {
			t.Errorf("To = %q, want %q", convErr.To, "int32")
		}
	})

	t.Run("Uint32ToInt16 overflow error details", func(t *testing.T) {
		input := uint32(math.MaxInt16 + 1)
		_, err := Uint32ToInt16(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint32" {
			t.Errorf("From = %q, want %q", convErr.From, "uint32")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
	})

	t.Run("Uint32ToInt8 overflow error details", func(t *testing.T) {
		input := uint32(math.MaxInt8 + 1)
		_, err := Uint32ToInt8(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint32" {
			t.Errorf("From = %q, want %q", convErr.From, "uint32")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("Uint16ToInt16 overflow error details", func(t *testing.T) {
		input := uint16(math.MaxInt16 + 1)
		_, err := Uint16ToInt16(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint16" {
			t.Errorf("From = %q, want %q", convErr.From, "uint16")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
	})

	t.Run("Uint16ToInt8 overflow error details", func(t *testing.T) {
		input := uint16(math.MaxInt8 + 1)
		_, err := Uint16ToInt8(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint16" {
			t.Errorf("From = %q, want %q", convErr.From, "uint16")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("Uint8ToInt8 overflow error details", func(t *testing.T) {
		input := uint8(math.MaxInt8 + 1)
		_, err := Uint8ToInt8(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint8" {
			t.Errorf("From = %q, want %q", convErr.From, "uint8")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("UintToInt64 overflow error details", func(t *testing.T) {
		input := uint(math.MaxInt64) + 1
		_, err := UintToInt64(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "int64" {
			t.Errorf("To = %q, want %q", convErr.To, "int64")
		}
	})

	t.Run("UintToInt32 overflow error details", func(t *testing.T) {
		input := uint(math.MaxInt32 + 1)
		_, err := UintToInt32(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "int32" {
			t.Errorf("To = %q, want %q", convErr.To, "int32")
		}
	})

	t.Run("UintToInt16 overflow error details", func(t *testing.T) {
		input := uint(math.MaxInt16 + 1)
		_, err := UintToInt16(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
	})

	t.Run("UintToInt8 overflow error details", func(t *testing.T) {
		input := uint(math.MaxInt8 + 1)
		_, err := UintToInt8(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("UintToInt overflow error details", func(t *testing.T) {
		input := uint(math.MaxInt) + 1
		_, err := UintToInt(input)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "int" {
			t.Errorf("To = %q, want %q", convErr.To, "int")
		}
	})
}

// =============================================================================
// Tests for always-succeed functions (no error return)
// =============================================================================

func TestAlwaysSucceedFunctions(t *testing.T) {
	// Verify that the always-succeed functions work correctly at boundary values

	t.Run("Uint32ToInt64 with max uint32", func(t *testing.T) {
		got := Uint32ToInt64(math.MaxUint32)
		if got != math.MaxUint32 {
			t.Errorf("Uint32ToInt64(MaxUint32) = %d, want %d", got, int64(math.MaxUint32))
		}
	})

	t.Run("Uint32ToInt with max uint32", func(t *testing.T) {
		got := Uint32ToInt(math.MaxUint32)
		if got != math.MaxUint32 {
			t.Errorf("Uint32ToInt(MaxUint32) = %d, want %d", got, int(math.MaxUint32))
		}
	})

	t.Run("Uint16ToInt64 with max uint16", func(t *testing.T) {
		got := Uint16ToInt64(math.MaxUint16)
		if got != math.MaxUint16 {
			t.Errorf("Uint16ToInt64(MaxUint16) = %d, want %d", got, int64(math.MaxUint16))
		}
	})

	t.Run("Uint16ToInt32 with max uint16", func(t *testing.T) {
		got := Uint16ToInt32(math.MaxUint16)
		if got != math.MaxUint16 {
			t.Errorf("Uint16ToInt32(MaxUint16) = %d, want %d", got, int32(math.MaxUint16))
		}
	})

	t.Run("Uint16ToInt with max uint16", func(t *testing.T) {
		got := Uint16ToInt(math.MaxUint16)
		if got != math.MaxUint16 {
			t.Errorf("Uint16ToInt(MaxUint16) = %d, want %d", got, int(math.MaxUint16))
		}
	})

	t.Run("Uint8ToInt64 with max uint8", func(t *testing.T) {
		got := Uint8ToInt64(math.MaxUint8)
		if got != math.MaxUint8 {
			t.Errorf("Uint8ToInt64(MaxUint8) = %d, want %d", got, int64(math.MaxUint8))
		}
	})

	t.Run("Uint8ToInt32 with max uint8", func(t *testing.T) {
		got := Uint8ToInt32(math.MaxUint8)
		if got != math.MaxUint8 {
			t.Errorf("Uint8ToInt32(MaxUint8) = %d, want %d", got, int32(math.MaxUint8))
		}
	})

	t.Run("Uint8ToInt16 with max uint8", func(t *testing.T) {
		got := Uint8ToInt16(math.MaxUint8)
		if got != math.MaxUint8 {
			t.Errorf("Uint8ToInt16(MaxUint8) = %d, want %d", got, int16(math.MaxUint8))
		}
	})

	t.Run("Uint8ToInt with max uint8", func(t *testing.T) {
		got := Uint8ToInt(math.MaxUint8)
		if got != math.MaxUint8 {
			t.Errorf("Uint8ToInt(MaxUint8) = %d, want %d", got, int(math.MaxUint8))
		}
	})
}
