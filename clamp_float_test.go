package safeconv

import (
	"errors"
	"math"
	"testing"
)

func TestClampFloat64TruncToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100.7, 100, nil},
		{"negative", -100.7, -100, nil},
		{"overflow clamps to max", 1e20, math.MaxInt64, nil},
		{"underflow clamps to min", -1e20, math.MinInt64, nil},
		{"NaN errors", math.NaN(), 0, ErrNaN},
		{"positive infinity errors", math.Inf(1), 0, ErrInfinity},
		{"negative infinity errors", math.Inf(-1), 0, ErrInfinity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClampFloat64TruncToInt64(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ClampFloat64TruncToInt64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ClampFloat64TruncToInt64(%v) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("ClampFloat64TruncToInt64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampFloat64TruncToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100.7, 100, nil},
		{"negative", -100.7, -100, nil},
		{"max int32", float64(math.MaxInt32), math.MaxInt32, nil},
		{"min int32", float64(math.MinInt32), math.MinInt32, nil},
		{"overflow clamps to max", float64(math.MaxInt32) + 1000, math.MaxInt32, nil},
		{"underflow clamps to min", float64(math.MinInt32) - 1000, math.MinInt32, nil},
		{"NaN errors", math.NaN(), 0, ErrNaN},
		{"infinity errors", math.Inf(1), 0, ErrInfinity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClampFloat64TruncToInt32(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ClampFloat64TruncToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ClampFloat64TruncToInt32(%v) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("ClampFloat64TruncToInt32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampFloat64TruncToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100.7, 100, nil},
		{"negative clamps to zero", -100.7, 0, nil},
		{"max uint32", float64(math.MaxUint32), math.MaxUint32, nil},
		{"overflow clamps to max", float64(math.MaxUint32) + 1000, math.MaxUint32, nil},
		{"NaN errors", math.NaN(), 0, ErrNaN},
		{"infinity errors", math.Inf(1), 0, ErrInfinity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClampFloat64TruncToUint32(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ClampFloat64TruncToUint32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ClampFloat64TruncToUint32(%v) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("ClampFloat64TruncToUint32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampFloat64RoundToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    int64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"rounds up", 100.5, 101, nil},
		{"rounds down", 100.4, 100, nil},
		{"negative rounds", -100.5, -101, nil},
		{"overflow clamps to max", 1e20, math.MaxInt64, nil},
		{"underflow clamps to min", -1e20, math.MinInt64, nil},
		{"NaN errors", math.NaN(), 0, ErrNaN},
		{"infinity errors", math.Inf(1), 0, ErrInfinity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClampFloat64RoundToInt64(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ClampFloat64RoundToInt64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ClampFloat64RoundToInt64(%v) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("ClampFloat64RoundToInt64(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampFloat64RoundToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		want    uint32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"rounds up", 100.5, 101, nil},
		{"rounds down", 100.4, 100, nil},
		{"negative clamps to zero", -100.5, 0, nil},
		{"overflow clamps to max", 1e20, math.MaxUint32, nil},
		{"NaN errors", math.NaN(), 0, ErrNaN},
		{"infinity errors", math.Inf(1), 0, ErrInfinity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClampFloat64RoundToUint32(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ClampFloat64RoundToUint32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ClampFloat64RoundToUint32(%v) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("ClampFloat64RoundToUint32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampFloat32TruncToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100.7, 100, nil},
		{"negative", -100.7, -100, nil},
		{"NaN errors", float32(math.NaN()), 0, ErrNaN},
		{"infinity errors", float32(math.Inf(1)), 0, ErrInfinity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClampFloat32TruncToInt32(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ClampFloat32TruncToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ClampFloat32TruncToInt32(%v) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("ClampFloat32TruncToInt32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampFloat32RoundToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   float32
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"rounds up", 100.5, 101, nil},
		{"rounds down", 100.4, 100, nil},
		{"NaN errors", float32(math.NaN()), 0, ErrNaN},
		{"infinity errors", float32(math.Inf(1)), 0, ErrInfinity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClampFloat32RoundToInt32(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ClampFloat32RoundToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ClampFloat32RoundToInt32(%v) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("ClampFloat32RoundToInt32(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
