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
		isNaN   bool
		isInf   int // 0 = not inf, 1 = +inf, -1 = -inf
	}{
		{"zero", 0, 0, nil, false, 0},
		{"positive", 3.14, 3.14, nil, false, 0},
		{"negative", -3.14, -3.14, nil, false, 0},
		{"max float32", math.MaxFloat32, math.MaxFloat32, nil, false, 0},
		{"min float32", -math.MaxFloat32, -math.MaxFloat32, nil, false, 0},
		{"small positive", 1e-10, 1e-10, nil, false, 0},
		{"small negative", -1e-10, -1e-10, nil, false, 0},
		{"nan", math.NaN(), 0, nil, true, 0},
		{"positive inf", math.Inf(1), 0, nil, false, 1},
		{"negative inf", math.Inf(-1), 0, nil, false, -1},
		{"overflow positive", math.MaxFloat64, 0, ErrOverflow, false, 0},
		{"overflow negative", -math.MaxFloat64, 0, ErrOverflow, false, 0},
		{"overflow just above max float32", math.MaxFloat32 * 2, 0, ErrOverflow, false, 0},
		{"overflow just below min float32", -math.MaxFloat32 * 2, 0, ErrOverflow, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64ToFloat32(tt.input)

			// Check error
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Float64ToFloat32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}

			// Skip value comparison on error cases
			if err != nil {
				return
			}

			// Handle NaN comparison specially (NaN != NaN)
			if tt.isNaN {
				if !math.IsNaN(float64(got)) {
					t.Errorf("Float64ToFloat32(%v) = %v, want NaN", tt.input, got)
				}
				return
			}

			// Handle Inf comparison specially
			if tt.isInf != 0 {
				if !math.IsInf(float64(got), tt.isInf) {
					t.Errorf("Float64ToFloat32(%v) = %v, want Inf with sign %d", tt.input, got, tt.isInf)
				}
				return
			}

			// Normal value comparison
			if got != tt.want {
				t.Errorf("Float64ToFloat32(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFloat64ToFloat32_PrecisionLoss(t *testing.T) {
	// Precision loss is accepted and not an error
	// float64 has ~15-17 decimal digits of precision
	// float32 has ~6-9 decimal digits of precision
	tests := []struct {
		name  string
		input float64
	}{
		{"many decimal places", 3.141592653589793},
		{"large with precision", 123456.789012345},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Float64ToFloat32(tt.input)
			if err != nil {
				t.Errorf("Float64ToFloat32(%v) unexpected error: %v", tt.input, err)
				return
			}
			// Just verify we got a valid float32, precision loss is expected
			if math.IsNaN(float64(got)) || math.IsInf(float64(got), 0) {
				t.Errorf("Float64ToFloat32(%v) = %v, expected a finite value", tt.input, got)
			}
		})
	}
}

func TestFloat64ToFloat32_ConversionError(t *testing.T) {
	t.Run("overflow error details", func(t *testing.T) {
		_, err := Float64ToFloat32(math.MaxFloat64)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "float32" {
			t.Errorf("To = %q, want %q", convErr.To, "float32")
		}
		if !errors.Is(convErr, ErrOverflow) {
			t.Errorf("expected ErrOverflow, got %v", convErr.Err)
		}
	})

	t.Run("negative overflow error details", func(t *testing.T) {
		_, err := Float64ToFloat32(-math.MaxFloat64)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "float64" {
			t.Errorf("From = %q, want %q", convErr.From, "float64")
		}
		if convErr.To != "float32" {
			t.Errorf("To = %q, want %q", convErr.To, "float32")
		}
		if !errors.Is(convErr, ErrOverflow) {
			t.Errorf("expected ErrOverflow, got %v", convErr.Err)
		}
	})
}

func TestFloat64ToFloat32_EdgeCases(t *testing.T) {
	t.Run("positive zero", func(t *testing.T) {
		got, err := Float64ToFloat32(0.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != 0 {
			t.Errorf("Float64ToFloat32(0.0) = %v, want 0", got)
		}
	})

	t.Run("negative zero", func(t *testing.T) {
		got, err := Float64ToFloat32(math.Copysign(0, -1))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// Check that it's negative zero
		if got != 0 || !math.Signbit(float64(got)) {
			t.Errorf("Float64ToFloat32(-0.0) = %v, want -0", got)
		}
	})

	t.Run("smallest positive float32", func(t *testing.T) {
		got, err := Float64ToFloat32(math.SmallestNonzeroFloat32)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != math.SmallestNonzeroFloat32 {
			t.Errorf("Float64ToFloat32(SmallestNonzeroFloat32) = %v, want %v", got, math.SmallestNonzeroFloat32)
		}
	})

	t.Run("value at boundary", func(t *testing.T) {
		// MaxFloat32 should succeed
		got, err := Float64ToFloat32(math.MaxFloat32)
		if err != nil {
			t.Errorf("unexpected error at MaxFloat32: %v", err)
		}
		if got != math.MaxFloat32 {
			t.Errorf("Float64ToFloat32(MaxFloat32) = %v, want %v", got, math.MaxFloat32)
		}

		// Just above MaxFloat32 should fail
		_, err = Float64ToFloat32(math.MaxFloat32 * 1.0001)
		if !errors.Is(err, ErrOverflow) {
			t.Errorf("Float64ToFloat32(MaxFloat32*1.0001) error = %v, want ErrOverflow", err)
		}
	})
}
