package safeconv

import (
	"errors"
	"math"
	"testing"
)

// =============================================================================
// Float64 to Int* truncation conversion tests
// =============================================================================

func TestFloat64TruncToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"small positive fraction", 0.9, 0, nil},
		{"small negative fraction", -0.9, 0, nil},
		{"max int64 representable", 9223372036854775000.0, 9223372036854774784, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", 1e19, 0, ErrOverflow},
		{"underflow", -1e19, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToInt64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToInt64(%v) = %d, want %d", tt.input, got, tt.want)
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
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int32", float64(math.MaxInt32), math.MaxInt32, nil},
		{"min int32", float64(math.MinInt32), math.MinInt32, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", float64(math.MaxInt32) + 1, 0, ErrOverflow},
		{"underflow", float64(math.MinInt32) - 1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToInt32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int16", float64(math.MaxInt16), math.MaxInt16, nil},
		{"min int16", float64(math.MinInt16), math.MinInt16, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", float64(math.MaxInt16) + 1, 0, ErrOverflow},
		{"underflow", float64(math.MinInt16) - 1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToInt16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToInt16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int8", float64(math.MaxInt8), math.MaxInt8, nil},
		{"min int8", float64(math.MinInt8), math.MinInt8, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", float64(math.MaxInt8) + 1, 0, ErrOverflow},
		{"underflow", float64(math.MinInt8) - 1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToInt8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToInt8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", 1e19, 0, ErrOverflow},
		{"underflow", -1e19, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToInt(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToInt(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Float64 to Uint* truncation conversion tests
// =============================================================================

func TestFloat64TruncToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"small positive fraction", 0.9, 0, nil},
		{"max uint64 representable", 1.8446744073709550e19, 18446744073709549568, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"negative fraction", -0.5, 0, ErrUnderflow},
		{"overflow", 2e19, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToUint64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToUint64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint32
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint32", float64(math.MaxUint32), math.MaxUint32, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float64(math.MaxUint32) + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToUint32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToUint32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint16", float64(math.MaxUint16), math.MaxUint16, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float64(math.MaxUint16) + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToUint16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToUint16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint8", float64(math.MaxUint8), math.MaxUint8, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float64(math.MaxUint8) + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToUint8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToUint8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64TruncToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", 2e19, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64TruncToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64TruncToUint(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64TruncToUint(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Float32 to Int* truncation conversion tests
// =============================================================================

func TestFloat32TruncToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"large positive", 1e15, 999999986991104, nil},   // float32 precision limit
		{"large negative", -1e15, -999999986991104, nil}, // float32 precision limit
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToInt64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToInt64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int32
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"overflow", float32(math.MaxInt32) + 1000, 0, ErrOverflow},
		{"underflow", float32(math.MinInt32) - 1000, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToInt32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int16", float32(math.MaxInt16), math.MaxInt16, nil},
		{"min int16", float32(math.MinInt16), math.MinInt16, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"overflow", float32(math.MaxInt16) + 1, 0, ErrOverflow},
		{"underflow", float32(math.MinInt16) - 1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToInt16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToInt16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int8", float32(math.MaxInt8), math.MaxInt8, nil},
		{"min int8", float32(math.MinInt8), math.MinInt8, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"overflow", float32(math.MaxInt8) + 1, 0, ErrOverflow},
		{"underflow", float32(math.MinInt8) - 1, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToInt8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToInt8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"negative truncate toward zero", -3.7, -3, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToInt(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToInt(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Float32 to Uint* truncation conversion tests
// =============================================================================

func TestFloat32TruncToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"large positive", 1e15, 999999986991104, nil}, // float32 precision limit
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToUint64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToUint64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint32
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float32(math.MaxUint32) + 10000, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToUint32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToUint32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint16", float32(math.MaxUint16), math.MaxUint16, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float32(math.MaxUint16) + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToUint16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToUint16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint8", float32(math.MaxUint8), math.MaxUint8, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float32(math.MaxUint8) + 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToUint8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToUint8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32TruncToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive truncate down", 3.7, 3, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32TruncToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32TruncToUint(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32TruncToUint(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// ConversionError tests for float truncation
// =============================================================================

func TestFloatTruncConversionErrors(t *testing.T) {
	t.Run("Float64TruncToInt64 NaN error details", func(t *testing.T) {
		_, err := Float64TruncToInt64(math.NaN())
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "int64" {
			t.Errorf("To = %q, want %q", convErr.To, "int64")
		}
		if !errors.Is(convErr, ErrNaN) {
			t.Errorf("expected ErrNaN, got %v", convErr.Err)
		}
	})

	t.Run("Float64TruncToInt64 infinity error details", func(t *testing.T) {
		_, err := Float64TruncToInt64(math.Inf(1))
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "int64" {
			t.Errorf("To = %q, want %q", convErr.To, "int64")
		}
		if !errors.Is(convErr, ErrInfinity) {
			t.Errorf("expected ErrInfinity, got %v", convErr.Err)
		}
	})

	t.Run("Float64TruncToInt32 overflow error details", func(t *testing.T) {
		_, err := Float64TruncToInt32(float64(math.MaxInt32) + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "int32" {
			t.Errorf("To = %q, want %q", convErr.To, "int32")
		}
		if !errors.Is(convErr, ErrOverflow) {
			t.Errorf("expected ErrOverflow, got %v", convErr.Err)
		}
	})

	t.Run("Float64TruncToInt16 underflow error details", func(t *testing.T) {
		_, err := Float64TruncToInt16(float64(math.MinInt16) - 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
		if !errors.Is(convErr, ErrUnderflow) {
			t.Errorf("expected ErrUnderflow, got %v", convErr.Err)
		}
	})

	t.Run("Float64TruncToUint64 negative error details", func(t *testing.T) {
		_, err := Float64TruncToUint64(-1.0)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "uint64" {
			t.Errorf("To = %q, want %q", convErr.To, "uint64")
		}
		if !errors.Is(convErr, ErrUnderflow) {
			t.Errorf("expected ErrUnderflow, got %v", convErr.Err)
		}
	})

	t.Run("Float32TruncToInt64 NaN error details", func(t *testing.T) {
		_, err := Float32TruncToInt64(float32(math.NaN()))
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float32" {
			t.Errorf("From = %q, want %q", convErr.From, "float32")
		}
		if convErr.To != "int64" {
			t.Errorf("To = %q, want %q", convErr.To, "int64")
		}
		if !errors.Is(convErr, ErrNaN) {
			t.Errorf("expected ErrNaN, got %v", convErr.Err)
		}
	})

	t.Run("Float32TruncToUint32 negative error details", func(t *testing.T) {
		_, err := Float32TruncToUint32(-1.0)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float32" {
			t.Errorf("From = %q, want %q", convErr.From, "float32")
		}
		if convErr.To != "uint32" {
			t.Errorf("To = %q, want %q", convErr.To, "uint32")
		}
		if !errors.Is(convErr, ErrUnderflow) {
			t.Errorf("expected ErrUnderflow, got %v", convErr.Err)
		}
	})

	t.Run("Float64TruncToUint overflow error details", func(t *testing.T) {
		_, err := Float64TruncToUint(2e19)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "uint" {
			t.Errorf("To = %q, want %q", convErr.To, "uint")
		}
		if !errors.Is(convErr, ErrOverflow) {
			t.Errorf("expected ErrOverflow, got %v", convErr.Err)
		}
	})

	t.Run("Float64TruncToInt8 overflow error details", func(t *testing.T) {
		_, err := Float64TruncToInt8(float64(math.MaxInt8) + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
		if !errors.Is(convErr, ErrOverflow) {
			t.Errorf("expected ErrOverflow, got %v", convErr.Err)
		}
	})

	t.Run("Float32TruncToInt8 underflow error details", func(t *testing.T) {
		_, err := Float32TruncToInt8(float32(math.MinInt8) - 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float32" {
			t.Errorf("From = %q, want %q", convErr.From, "float32")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
		if !errors.Is(convErr, ErrUnderflow) {
			t.Errorf("expected ErrUnderflow, got %v", convErr.Err)
		}
	})
}
