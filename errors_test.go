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
