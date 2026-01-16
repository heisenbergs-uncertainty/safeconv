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
		{"max int64", math.MaxInt64, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
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
		{"max int64", math.MaxInt64, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
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
		{"max int64", math.MaxInt64, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
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
		{"max int64", math.MaxInt64, math.MaxInt64, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
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

// TestInt64ToUintConversionError verifies that errors are ConversionErrors with correct fields.
func TestInt64ToUintConversionError(t *testing.T) {
	t.Run("Int64ToUint64 underflow error details", func(t *testing.T) {
		_, err := Int64ToUint64(-1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "uint64" {
			t.Errorf("To = %q, want %q", convErr.To, "uint64")
		}
		if convErr.Value != int64(-1) {
			t.Errorf("Value = %v, want %v", convErr.Value, int64(-1))
		}
	})

	t.Run("Int64ToUint32 overflow error details", func(t *testing.T) {
		_, err := Int64ToUint32(math.MaxUint32 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "uint32" {
			t.Errorf("To = %q, want %q", convErr.To, "uint32")
		}
		if convErr.Value != int64(math.MaxUint32+1) {
			t.Errorf("Value = %v, want %v", convErr.Value, int64(math.MaxUint32+1))
		}
	})

	t.Run("Int64ToUint16 overflow error details", func(t *testing.T) {
		_, err := Int64ToUint16(math.MaxUint16 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "uint16" {
			t.Errorf("To = %q, want %q", convErr.To, "uint16")
		}
	})

	t.Run("Int64ToUint8 overflow error details", func(t *testing.T) {
		_, err := Int64ToUint8(math.MaxUint8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "uint8" {
			t.Errorf("To = %q, want %q", convErr.To, "uint8")
		}
	})

	t.Run("Int64ToUint underflow error details", func(t *testing.T) {
		_, err := Int64ToUint(-100)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "uint" {
			t.Errorf("To = %q, want %q", convErr.To, "uint")
		}
	})
}
