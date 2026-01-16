package safeconv

import (
	"errors"
	"math"
	"testing"
)

// =============================================================================
// Int64 to Int* narrowing conversion tests
// =============================================================================

func TestInt64ToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"min int32", math.MinInt32, math.MinInt32, nil},
		{"max int32 + 1", math.MaxInt32 + 1, 0, ErrOverflow},
		{"min int32 - 1", math.MinInt32 - 1, 0, ErrUnderflow},
		{"max int64", math.MaxInt64, 0, ErrOverflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToInt32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt64ToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    int16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"min int16", math.MinInt16, math.MinInt16, nil},
		{"max int16 + 1", math.MaxInt16 + 1, 0, ErrOverflow},
		{"min int16 - 1", math.MinInt16 - 1, 0, ErrUnderflow},
		{"max int64", math.MaxInt64, 0, ErrOverflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToInt16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt64ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"min int8", math.MinInt8, math.MinInt8, nil},
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"min int8 - 1", math.MinInt8 - 1, 0, ErrUnderflow},
		{"max int64", math.MaxInt64, 0, ErrOverflow},
		{"min int64", math.MinInt64, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt64ToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    int
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int", math.MaxInt, math.MaxInt, nil},
		{"min int", math.MinInt, math.MinInt, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToInt(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int64ToInt(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Int32 to Int* conversion tests
// =============================================================================

func TestInt32ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"min int32", math.MinInt32, math.MinInt32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("Int32ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt32ToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    int16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"min int16", math.MinInt16, math.MinInt16, nil},
		{"max int16 + 1", math.MaxInt16 + 1, 0, ErrOverflow},
		{"min int16 - 1", math.MinInt16 - 1, 0, ErrUnderflow},
		{"max int32", math.MaxInt32, 0, ErrOverflow},
		{"min int32", math.MinInt32, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int32ToInt16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt32ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"min int8", math.MinInt8, math.MinInt8, nil},
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"min int8 - 1", math.MinInt8 - 1, 0, ErrUnderflow},
		{"max int32", math.MaxInt32, 0, ErrOverflow},
		{"min int32", math.MinInt32, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int32ToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Int16 to Int* conversion tests
// =============================================================================

func TestInt16ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"min int16", math.MinInt16, math.MinInt16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("Int16ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt16ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  int32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"min int16", math.MinInt16, math.MinInt16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16ToInt32(tt.input)
			if got != tt.want {
				t.Errorf("Int16ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt16ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"min int8", math.MinInt8, math.MinInt8, nil},
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"min int8 - 1", math.MinInt8 - 1, 0, ErrUnderflow},
		{"max int16", math.MaxInt16, 0, ErrOverflow},
		{"min int16", math.MinInt16, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Int16ToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Int8 to Int* widening conversion tests (always succeed)
// =============================================================================

func TestInt8ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"min int8", math.MinInt8, math.MinInt8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8ToInt64(tt.input)
			if got != tt.want {
				t.Errorf("Int8ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt8ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  int32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"min int8", math.MinInt8, math.MinInt8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8ToInt32(tt.input)
			if got != tt.want {
				t.Errorf("Int8ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt8ToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  int16
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"min int8", math.MinInt8, math.MinInt8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8ToInt16(tt.input)
			if got != tt.want {
				t.Errorf("Int8ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Int to Int* conversion tests
// =============================================================================

func TestIntToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
		{"max int", math.MaxInt, math.MaxInt},
		{"min int", math.MinInt, math.MinInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToInt64(tt.input)
			if got != tt.want {
				t.Errorf("IntToInt64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIntToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int32
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int32", math.MaxInt32, math.MaxInt32, nil},
		{"min int32", math.MinInt32, math.MinInt32, nil},
		{"max int32 + 1", math.MaxInt32 + 1, 0, ErrOverflow},
		{"min int32 - 1", math.MinInt32 - 1, 0, ErrUnderflow},
		{"max int", math.MaxInt, 0, ErrOverflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToInt32(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToInt32(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIntToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int16", math.MaxInt16, math.MaxInt16, nil},
		{"min int16", math.MinInt16, math.MinInt16, nil},
		{"max int16 + 1", math.MaxInt16 + 1, 0, ErrOverflow},
		{"min int16 - 1", math.MinInt16 - 1, 0, ErrUnderflow},
		{"max int", math.MaxInt, 0, ErrOverflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToInt16(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToInt16(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIntToInt8(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"positive", 100, 100, nil},
		{"negative", -100, -100, nil},
		{"max int8", math.MaxInt8, math.MaxInt8, nil},
		{"min int8", math.MinInt8, math.MinInt8, nil},
		{"max int8 + 1", math.MaxInt8 + 1, 0, ErrOverflow},
		{"min int8 - 1", math.MinInt8 - 1, 0, ErrUnderflow},
		{"max int", math.MaxInt, 0, ErrOverflow},
		{"min int", math.MinInt, 0, ErrUnderflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToInt8(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IntToInt8(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// ConversionError tests for int narrowing
// =============================================================================

func TestIntNarrowConversionErrors(t *testing.T) {
	t.Run("Int64ToInt32 overflow error details", func(t *testing.T) {
		_, err := Int64ToInt32(math.MaxInt32 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "int32" {
			t.Errorf("To = %q, want %q", convErr.To, "int32")
		}
		if convErr.Value != int64(math.MaxInt32+1) {
			t.Errorf("Value = %v, want %v", convErr.Value, int64(math.MaxInt32+1))
		}
		if !errors.Is(convErr, ErrOverflow) {
			t.Errorf("expected ErrOverflow, got %v", convErr.Err)
		}
	})

	t.Run("Int64ToInt32 underflow error details", func(t *testing.T) {
		_, err := Int64ToInt32(math.MinInt32 - 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "int32" {
			t.Errorf("To = %q, want %q", convErr.To, "int32")
		}
		if !errors.Is(convErr, ErrUnderflow) {
			t.Errorf("expected ErrUnderflow, got %v", convErr.Err)
		}
	})

	t.Run("Int64ToInt16 overflow error details", func(t *testing.T) {
		_, err := Int64ToInt16(math.MaxInt16 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
	})

	t.Run("Int64ToInt8 underflow error details", func(t *testing.T) {
		_, err := Int64ToInt8(math.MinInt8 - 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int64" {
			t.Errorf("From = %q, want %q", convErr.From, "int64")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("Int32ToInt16 overflow error details", func(t *testing.T) {
		_, err := Int32ToInt16(math.MaxInt16 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int32" {
			t.Errorf("From = %q, want %q", convErr.From, "int32")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
	})

	t.Run("Int32ToInt8 underflow error details", func(t *testing.T) {
		_, err := Int32ToInt8(math.MinInt8 - 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int32" {
			t.Errorf("From = %q, want %q", convErr.From, "int32")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("Int16ToInt8 overflow error details", func(t *testing.T) {
		_, err := Int16ToInt8(math.MaxInt8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int16" {
			t.Errorf("From = %q, want %q", convErr.From, "int16")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})

	t.Run("IntToInt32 overflow error details", func(t *testing.T) {
		_, err := IntToInt32(math.MaxInt32 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int" {
			t.Errorf("From = %q, want %q", convErr.From, "int")
		}
		if convErr.To != "int32" {
			t.Errorf("To = %q, want %q", convErr.To, "int32")
		}
	})

	t.Run("IntToInt16 underflow error details", func(t *testing.T) {
		_, err := IntToInt16(math.MinInt16 - 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int" {
			t.Errorf("From = %q, want %q", convErr.From, "int")
		}
		if convErr.To != "int16" {
			t.Errorf("To = %q, want %q", convErr.To, "int16")
		}
	})

	t.Run("IntToInt8 overflow error details", func(t *testing.T) {
		_, err := IntToInt8(math.MaxInt8 + 1)
		var convErr *ConversionError
		if !errors.As(err, &convErr) {
			t.Fatalf("expected ConversionError, got %T", err)
		}
		if convErr.From != "int" {
			t.Errorf("From = %q, want %q", convErr.From, "int")
		}
		if convErr.To != "int8" {
			t.Errorf("To = %q, want %q", convErr.To, "int8")
		}
	})
}
