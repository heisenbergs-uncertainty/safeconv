package safeconv

import (
	"math"
	"testing"
)

func TestClampInt64ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  int32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"min int32", math.MinInt32, math.MinInt32},
		{"overflow clamps to max", math.MaxInt32 + 1, math.MaxInt32},
		{"underflow clamps to min", math.MinInt32 - 1, math.MinInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToInt32(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt64ToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  int16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"min int16", math.MinInt16, math.MinInt16},
		{"overflow clamps to max", math.MaxInt16 + 1, math.MaxInt16},
		{"underflow clamps to min", math.MinInt16 - 1, math.MinInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToInt16(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt64ToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"min int8", math.MinInt8, math.MinInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
		{"underflow clamps to min", math.MinInt8 - 1, math.MinInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToInt8(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt64ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative", -100, -100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToInt(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt32ToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  int16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"min int16", math.MinInt16, math.MinInt16},
		{"overflow clamps to max", math.MaxInt16 + 1, math.MaxInt16},
		{"underflow clamps to min", math.MinInt16 - 1, math.MinInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt32ToInt16(tt.input); got != tt.want {
				t.Errorf("ClampInt32ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt32ToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"min int8", math.MinInt8, math.MinInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
		{"underflow clamps to min", math.MinInt8 - 1, math.MinInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt32ToInt8(tt.input); got != tt.want {
				t.Errorf("ClampInt32ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt16ToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"min int8", math.MinInt8, math.MinInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
		{"underflow clamps to min", math.MinInt8 - 1, math.MinInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt16ToInt8(tt.input); got != tt.want {
				t.Errorf("ClampInt16ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"min int32", math.MinInt32, math.MinInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToInt32(tt.input); got != tt.want {
				t.Errorf("ClampIntToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"min int16", math.MinInt16, math.MinInt16},
		{"overflow clamps to max", math.MaxInt16 + 1, math.MaxInt16},
		{"underflow clamps to min", math.MinInt16 - 1, math.MinInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToInt16(tt.input); got != tt.want {
				t.Errorf("ClampIntToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"negative in range", -100, -100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"min int8", math.MinInt8, math.MinInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
		{"underflow clamps to min", math.MinInt8 - 1, math.MinInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToInt8(tt.input); got != tt.want {
				t.Errorf("ClampIntToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
