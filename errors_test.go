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

func TestConversionError_IsOverflow(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"overflow error", ErrOverflow, true},
		{"underflow error", ErrUnderflow, false},
		{"nan error", ErrNaN, false},
		{"infinity error", ErrInfinity, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convErr := &ConversionError{
				From:  "int64",
				To:    "uint32",
				Value: int64(0),
				Err:   tt.err,
			}
			if got := convErr.IsOverflow(); got != tt.want {
				t.Errorf("IsOverflow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversionError_IsUnderflow(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"overflow error", ErrOverflow, false},
		{"underflow error", ErrUnderflow, true},
		{"nan error", ErrNaN, false},
		{"infinity error", ErrInfinity, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convErr := &ConversionError{
				From:  "int64",
				To:    "uint32",
				Value: int64(0),
				Err:   tt.err,
			}
			if got := convErr.IsUnderflow(); got != tt.want {
				t.Errorf("IsUnderflow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversionError_IsNaN(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"overflow error", ErrOverflow, false},
		{"underflow error", ErrUnderflow, false},
		{"nan error", ErrNaN, true},
		{"infinity error", ErrInfinity, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convErr := &ConversionError{
				From:  "float64",
				To:    "int64",
				Value: float64(0),
				Err:   tt.err,
			}
			if got := convErr.IsNaN(); got != tt.want {
				t.Errorf("IsNaN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversionError_IsInfinity(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"overflow error", ErrOverflow, false},
		{"underflow error", ErrUnderflow, false},
		{"nan error", ErrNaN, false},
		{"infinity error", ErrInfinity, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convErr := &ConversionError{
				From:  "float64",
				To:    "int64",
				Value: float64(0),
				Err:   tt.err,
			}
			if got := convErr.IsInfinity(); got != tt.want {
				t.Errorf("IsInfinity() = %v, want %v", got, tt.want)
			}
		})
	}
}
