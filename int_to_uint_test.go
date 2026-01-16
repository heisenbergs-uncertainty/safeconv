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

// Tests for Int32 to Uint* conversions

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
		{"negative one", -1, 0, ErrUnderflow},
		{"min int32", math.MinInt32, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int32ToUint64(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"negative one", -1, 0, ErrUnderflow},
		{"min int32", math.MinInt32, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int32ToUint32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max uint16 + 1", math.MaxUint16 + 1, 0, ErrOverflow},
		{"max int32", math.MaxInt32, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int32", math.MinInt32, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int32ToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
		{"max int32", math.MaxInt32, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int32", math.MinInt32, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int32ToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int32", math.MinInt32, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int32ToUint(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// Tests for Int16 to Uint* conversions

func TestInt16ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int16", math.MinInt16, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int16ToUint64(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int16", math.MinInt16, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int16ToUint32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"negative one", -1, 0, ErrUnderflow},
		{"min int16", math.MinInt16, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int16ToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
		{"max int16", math.MaxInt16, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int16", math.MinInt16, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int16ToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int16", math.MinInt16, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int16ToUint(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// Tests for Int8 to Uint* conversions

func TestInt8ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int8
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int8", math.MinInt8, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int8ToUint64(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int8", math.MinInt8, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int8ToUint32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int8", math.MinInt8, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int8ToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"negative one", -1, 0, ErrUnderflow},
		{"min int8", math.MinInt8, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int8ToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int8", math.MinInt8, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int8ToUint(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// Tests for Int to Uint* conversions

func TestIntToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"max int", math.MaxInt, math.MaxInt, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToUint64(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToUint64(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max uint32", math.MaxUint32, math.MaxUint32, nil},
		{"max uint32 + 1", math.MaxUint32 + 1, 0, ErrOverflow},
		{"max int", math.MaxInt, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToUint32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToUint32(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max uint16 + 1", math.MaxUint16 + 1, 0, ErrOverflow},
		{"max int", math.MaxInt, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToUint16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToUint16(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max uint8 + 1", math.MaxUint8 + 1, 0, ErrOverflow},
		{"max int", math.MaxInt, 0, ErrOverflow},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToUint8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToUint8(%d) = %d, want %d", tt.input, got, tt.want)
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
		{"max int", math.MaxInt, math.MaxInt, nil},
		{"negative one", -1, 0, ErrUnderflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToUint(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// TestNewConversionErrors verifies that errors from new functions are ConversionErrors with correct fields.
func TestNewConversionErrors(t *testing.T) {
	t.Run("Int32ToUint64 underflow error details", func(t *testing.T) {
		_, err := Int32ToUint64(-1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int32" {
			t.Errorf("From = %q, want %q", convErr.From, "int32")
		}
		if convErr.To != "uint64" {
			t.Errorf("To = %q, want %q", convErr.To, "uint64")
		}
	})

	t.Run("Int32ToUint16 overflow error details", func(t *testing.T) {
		_, err := Int32ToUint16(math.MaxUint16 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int32" {
			t.Errorf("From = %q, want %q", convErr.From, "int32")
		}
		if convErr.To != "uint16" {
			t.Errorf("To = %q, want %q", convErr.To, "uint16")
		}
	})

	t.Run("Int16ToUint8 overflow error details", func(t *testing.T) {
		_, err := Int16ToUint8(math.MaxUint8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int16" {
			t.Errorf("From = %q, want %q", convErr.From, "int16")
		}
		if convErr.To != "uint8" {
			t.Errorf("To = %q, want %q", convErr.To, "uint8")
		}
	})

	t.Run("Int8ToUint64 underflow error details", func(t *testing.T) {
		_, err := Int8ToUint64(-1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int8" {
			t.Errorf("From = %q, want %q", convErr.From, "int8")
		}
		if convErr.To != "uint64" {
			t.Errorf("To = %q, want %q", convErr.To, "uint64")
		}
	})

	t.Run("IntToUint32 overflow error details", func(t *testing.T) {
		_, err := IntToUint32(math.MaxUint32 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int" {
			t.Errorf("From = %q, want %q", convErr.From, "int")
		}
		if convErr.To != "uint32" {
			t.Errorf("To = %q, want %q", convErr.To, "uint32")
		}
	})

	t.Run("IntToUint underflow error details", func(t *testing.T) {
		_, err := IntToUint(-100)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int" {
			t.Errorf("From = %q, want %q", convErr.From, "int")
		}
		if convErr.To != "uint" {
			t.Errorf("To = %q, want %q", convErr.To, "uint")
		}
	})
}
