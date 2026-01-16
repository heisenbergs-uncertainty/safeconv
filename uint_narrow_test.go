package safeconv

import (
	"errors"
	"math"
	"testing"
)

// =============================================================================
// Uint64 to Uint* narrowing conversion tests
// =============================================================================

func TestUint64ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint32", math.MaxUint32, math.MaxUint32, nil},
		{"max uint32 + 1", math.MaxUint32 + 1, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToUint32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint64ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint16", math.MaxUint16, math.MaxUint16, nil},
		{"max uint16 + 1", math.MaxUint16 + 1, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint64ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
		{"max uint64", math.MaxUint64, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint64ToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    uint
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint32", math.MaxUint32, math.MaxUint32, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint64ToUint(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Uint32 to Uint* conversion tests
// =============================================================================

func TestUint32ToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32ToUint64(tt.input)
			if got != tt.want {
				t.Errorf("Uint32ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint32ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint16", math.MaxUint16, math.MaxUint16, nil},
		{"max uint16 + 1", math.MaxUint16 + 1, 0, ErrOverflow},
		{"max uint32", math.MaxUint32, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint32ToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint32ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
		{"max uint32", math.MaxUint32, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint32ToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Uint16 to Uint* conversion tests
// =============================================================================

func TestUint16ToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToUint64(tt.input)
			if got != tt.want {
				t.Errorf("Uint16ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint16ToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToUint32(tt.input)
			if got != tt.want {
				t.Errorf("Uint16ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint16ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
		{"max uint16", math.MaxUint16, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint16ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Uint16ToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint16ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Uint8 to Uint* widening conversion tests (always succeed)
// =============================================================================

func TestUint8ToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToUint64(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint8ToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToUint32(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint8ToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToUint16(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Uint to Uint* conversion tests
// =============================================================================

func TestUintToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UintToUint64(tt.input)
			if got != tt.want {
				t.Errorf("UintToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUintToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint32", math.MaxUint32, math.MaxUint32, nil},
		{"max uint32 + 1", math.MaxUint32 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToUint32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUintToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint16", math.MaxUint16, math.MaxUint16, nil},
		{"max uint16 + 1", math.MaxUint16 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestUintToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max uint8", math.MaxUint8, math.MaxUint8, nil},
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UintToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UintToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// ConversionError tests for uint narrowing
// =============================================================================

func TestUintNarrowConversionErrors(t *testing.T) {
	t.Run("Uint64ToUint32 overflow error details", func(t *testing.T) {
		_, err := Uint64ToUint32(math.MaxUint32 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "uint32" {
			t.Errorf("To = %q, want %q", convErr.To, "uint32")
		}
		if convErr.Value != uint64(math.MaxUint32+1) {
			t.Errorf("Value = %v, want %v", convErr.Value, uint64(math.MaxUint32+1))
		}
		if !errors.Is(convErr, ErrOverflow) {
			t.Errorf("expected ErrOverflow, got %v", convErr.Err)
		}
	})

	t.Run("Uint64ToUint16 overflow error details", func(t *testing.T) {
		_, err := Uint64ToUint16(math.MaxUint16 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "uint16" {
			t.Errorf("To = %q, want %q", convErr.To, "uint16")
		}
	})

	t.Run("Uint64ToUint8 overflow error details", func(t *testing.T) {
		_, err := Uint64ToUint8(math.MaxUint8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint64" {
			t.Errorf("From = %q, want %q", convErr.From, "uint64")
		}
		if convErr.To != "uint8" {
			t.Errorf("To = %q, want %q", convErr.To, "uint8")
		}
	})

	t.Run("Uint32ToUint16 overflow error details", func(t *testing.T) {
		_, err := Uint32ToUint16(math.MaxUint16 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint32" {
			t.Errorf("From = %q, want %q", convErr.From, "uint32")
		}
		if convErr.To != "uint16" {
			t.Errorf("To = %q, want %q", convErr.To, "uint16")
		}
	})

	t.Run("Uint32ToUint8 overflow error details", func(t *testing.T) {
		_, err := Uint32ToUint8(math.MaxUint8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint32" {
			t.Errorf("From = %q, want %q", convErr.From, "uint32")
		}
		if convErr.To != "uint8" {
			t.Errorf("To = %q, want %q", convErr.To, "uint8")
		}
	})

	t.Run("Uint16ToUint8 overflow error details", func(t *testing.T) {
		_, err := Uint16ToUint8(math.MaxUint8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint16" {
			t.Errorf("From = %q, want %q", convErr.From, "uint16")
		}
		if convErr.To != "uint8" {
			t.Errorf("To = %q, want %q", convErr.To, "uint8")
		}
	})

	t.Run("UintToUint32 overflow error details", func(t *testing.T) {
		_, err := UintToUint32(math.MaxUint32 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "uint32" {
			t.Errorf("To = %q, want %q", convErr.To, "uint32")
		}
	})

	t.Run("UintToUint16 overflow error details", func(t *testing.T) {
		_, err := UintToUint16(math.MaxUint16 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "uint16" {
			t.Errorf("To = %q, want %q", convErr.To, "uint16")
		}
	})

	t.Run("UintToUint8 overflow error details", func(t *testing.T) {
		_, err := UintToUint8(math.MaxUint8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "uint" {
			t.Errorf("From = %q, want %q", convErr.From, "uint")
		}
		if convErr.To != "uint8" {
			t.Errorf("To = %q, want %q", convErr.To, "uint8")
		}
	})
}
