package safeconv

import (
	"errors"
	"math"
	"testing"
)

// =============================================================================
// Float64 to Int* rounding conversion tests
// =============================================================================

func TestFloat64RoundToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"small positive fraction", 0.4, 0, nil},
		{"small positive half", 0.5, 1, nil},
		{"small negative fraction", -0.4, 0, nil},
		{"small negative half", -0.5, -1, nil},
		{"max int64 representable", 9223372036854775000.0, 9223372036854774784, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", 1e19, 0, ErrOverflow},
		{"underflow", -1e19, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToInt64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToInt64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int32
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int32", float64(math.MaxInt32), math.MaxInt32, nil},
		{"min int32", float64(math.MinInt32), math.MinInt32, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", float64(math.MaxInt32) + 1, 0, ErrOverflow},
		{"underflow", float64(math.MinInt32) - 1, 0, ErrUnderflow},
		// Edge case: rounding causes overflow
		{"rounding causes overflow", float64(math.MaxInt32) + 0.5, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToInt32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int16", float64(math.MaxInt16), math.MaxInt16, nil},
		{"min int16", float64(math.MinInt16), math.MinInt16, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", float64(math.MaxInt16) + 1, 0, ErrOverflow},
		{"underflow", float64(math.MinInt16) - 1, 0, ErrUnderflow},
		{"rounding causes overflow", float64(math.MaxInt16) + 0.5, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToInt16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToInt16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"max int8", float64(math.MaxInt8), math.MaxInt8, nil},
		{"min int8", float64(math.MinInt8), math.MinInt8, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"overflow", float64(math.MaxInt8) + 1, 0, ErrOverflow},
		{"underflow", float64(math.MinInt8) - 1, 0, ErrUnderflow},
		{"rounding causes overflow", float64(math.MaxInt8) + 0.5, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToInt8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToInt8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
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
			got, err := Float64RoundToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToInt(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToInt(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Float64 to Uint* rounding conversion tests
// =============================================================================

func TestFloat64RoundToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"small positive fraction", 0.4, 0, nil},
		{"small positive half", 0.5, 1, nil},
		{"max uint64 representable", 1.8446744073709550e19, 18446744073709549568, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"negative fraction rounds to negative", -0.6, 0, ErrUnderflow},
		{"negative half rounds to negative", -0.5, 0, ErrUnderflow},
		{"small negative rounds to zero", -0.4, 0, nil}, // rounds to 0, which is valid
		{"overflow", 2e19, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToUint64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToUint64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint32
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint32", float64(math.MaxUint32), math.MaxUint32, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float64(math.MaxUint32) + 1, 0, ErrOverflow},
		{"rounding causes overflow", float64(math.MaxUint32) + 0.5, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToUint32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToUint32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint16", float64(math.MaxUint16), math.MaxUint16, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float64(math.MaxUint16) + 1, 0, ErrOverflow},
		{"rounding causes overflow", float64(math.MaxUint16) + 0.5, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToUint16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToUint16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"max uint8", float64(math.MaxUint8), math.MaxUint8, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float64(math.MaxUint8) + 1, 0, ErrOverflow},
		{"rounding causes overflow", float64(math.MaxUint8) + 0.5, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToUint8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToUint8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64RoundToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", math.NaN(), 0, ErrNaN},
		{"positive infinity", math.Inf(1), 0, ErrInfinity},
		{"negative infinity", math.Inf(-1), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", 2e19, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64RoundToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64RoundToUint(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float64RoundToUint(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Float32 to Int* rounding conversion tests
// =============================================================================

func TestFloat32RoundToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
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
			got, err := Float32RoundToInt64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToInt64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToInt64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int32
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
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
			got, err := Float32RoundToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToInt32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
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
			got, err := Float32RoundToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToInt16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToInt16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
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
			got, err := Float32RoundToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToInt8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToInt8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"negative round toward zero", -3.4, -3, nil},
		{"negative half rounds away from zero", -3.5, -4, nil},
		{"negative round away from zero", -3.6, -4, nil},
		{"positive whole", 100.0, 100, nil},
		{"negative whole", -100.0, -100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32RoundToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToInt(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToInt(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Float32 to Uint* rounding conversion tests
// =============================================================================

func TestFloat32RoundToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint64
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"large positive", 1e15, 999999986991104, nil}, // float32 precision limit
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32RoundToUint64(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToUint64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToUint64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint32
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
		{"overflow", float32(math.MaxUint32) + 10000, 0, ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32RoundToUint32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToUint32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToUint32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint16
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
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
			got, err := Float32RoundToUint16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToUint16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToUint16(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint8
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
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
			got, err := Float32RoundToUint8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToUint8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToUint8(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat32RoundToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    uint
		wantErr error
	}{
		{"zero", 0.0, 0, nil},
		{"positive round down", 3.4, 3, nil},
		{"positive half rounds up", 3.5, 4, nil},
		{"positive round up", 3.6, 4, nil},
		{"positive whole", 100.0, 100, nil},
		{"NaN", float32(math.NaN()), 0, ErrNaN},
		{"positive infinity", float32(math.Inf(1)), 0, ErrInfinity},
		{"negative infinity", float32(math.Inf(-1)), 0, ErrInfinity},
		{"negative value", -1.0, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float32RoundToUint(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float32RoundToUint(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Float32RoundToUint(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// ConversionError tests for float rounding
// =============================================================================

func TestFloatRoundConversionErrors(t *testing.T) {
	t.Run("Float64RoundToInt64 NaN error details", func(t *testing.T) {
		_, err := Float64RoundToInt64(math.NaN())
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

	t.Run("Float64RoundToInt64 infinity error details", func(t *testing.T) {
		_, err := Float64RoundToInt64(math.Inf(1))
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

	t.Run("Float64RoundToInt32 overflow error details", func(t *testing.T) {
		_, err := Float64RoundToInt32(float64(math.MaxInt32) + 1)
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

	t.Run("Float64RoundToInt16 underflow error details", func(t *testing.T) {
		_, err := Float64RoundToInt16(float64(math.MinInt16) - 1)
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

	t.Run("Float64RoundToUint64 negative error details", func(t *testing.T) {
		_, err := Float64RoundToUint64(-1.0)
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

	t.Run("Float32RoundToInt64 NaN error details", func(t *testing.T) {
		_, err := Float32RoundToInt64(float32(math.NaN()))
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

	t.Run("Float32RoundToUint32 negative error details", func(t *testing.T) {
		_, err := Float32RoundToUint32(-1.0)
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

	t.Run("Float64RoundToUint overflow error details", func(t *testing.T) {
		_, err := Float64RoundToUint(2e19)
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

	t.Run("Float64RoundToInt8 overflow error details", func(t *testing.T) {
		_, err := Float64RoundToInt8(float64(math.MaxInt8) + 1)
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

	t.Run("Float32RoundToInt8 underflow error details", func(t *testing.T) {
		_, err := Float32RoundToInt8(float32(math.MinInt8) - 1)
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

// =============================================================================
// Rounding vs Truncation comparison tests
// =============================================================================

func TestRoundingVsTruncationBehavior(t *testing.T) {
	// These tests verify the key difference between rounding and truncation
	tests := []struct {
		name         string
		input        float64
		wantRound    int64
		wantTrunc    int64
		roundDiffers bool
	}{
		{"3.4 rounds down, truncates down", 3.4, 3, 3, false},
		{"3.5 rounds up, truncates down", 3.5, 4, 3, true},
		{"3.6 rounds up, truncates down", 3.6, 4, 3, true},
		{"-3.4 rounds toward zero, truncates toward zero", -3.4, -3, -3, false},
		{"-3.5 rounds away from zero, truncates toward zero", -3.5, -4, -3, true},
		{"-3.6 rounds away from zero, truncates toward zero", -3.6, -4, -3, true},
		{"2.0 same result", 2.0, 2, 2, false},
		{"-2.0 same result", -2.0, -2, -2, false},
		{"0.5 rounds up, truncates to 0", 0.5, 1, 0, true},
		{"-0.5 rounds to -1, truncates to 0", -0.5, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRound, err := Float64RoundToInt64(tt.input)
			if err != nil {
				t.Errorf("Float64RoundToInt64(%v) unexpected error: %v", tt.input, err)
				return
			}
			if gotRound != tt.wantRound {
				t.Errorf("Float64RoundToInt64(%v) = %d, want %d", tt.input, gotRound, tt.wantRound)
			}

			gotTrunc, err := Float64TruncToInt64(tt.input)
			if err != nil {
				t.Errorf("Float64TruncToInt64(%v) unexpected error: %v", tt.input, err)
				return
			}
			if gotTrunc != tt.wantTrunc {
				t.Errorf("Float64TruncToInt64(%v) = %d, want %d", tt.input, gotTrunc, tt.wantTrunc)
			}

			if tt.roundDiffers && gotRound == gotTrunc {
				t.Errorf("Expected round and trunc to differ for %v, but both returned %d", tt.input, gotRound)
			}
			if !tt.roundDiffers && gotRound != gotTrunc {
				t.Errorf("Expected round and trunc to be equal for %v, got round=%d trunc=%d", tt.input, gotRound, gotTrunc)
			}
		})
	}
}
