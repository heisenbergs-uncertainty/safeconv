package safeconv

import (
	"math"
	"testing"
)

func TestClampUint64ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  int64
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int64", math.MaxInt64, math.MaxInt64},
		{"overflow clamps to max", math.MaxInt64 + 1, math.MaxInt64},
		{"max uint64 clamps to max", math.MaxUint64, math.MaxInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToInt64(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToInt64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint64ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  int32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"overflow clamps to max", math.MaxInt32 + 1, math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToInt32(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint64ToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  int16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"overflow clamps to max", math.MaxInt16 + 1, math.MaxInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToInt16(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint64ToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToInt8(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint64ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  int
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int", math.MaxInt, math.MaxInt},
		{"overflow clamps to max", uint64(math.MaxInt) + 1, math.MaxInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToInt(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint32ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  int32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"overflow clamps to max", math.MaxInt32 + 1, math.MaxInt32},
		{"max uint32 clamps to max", math.MaxUint32, math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint32ToInt32(tt.input); got != tt.want {
				t.Errorf("ClampUint32ToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint32ToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  int16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"overflow clamps to max", math.MaxInt16 + 1, math.MaxInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint32ToInt16(tt.input); got != tt.want {
				t.Errorf("ClampUint32ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint32ToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint32ToInt8(tt.input); got != tt.want {
				t.Errorf("ClampUint32ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint16ToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  int16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"overflow clamps to max", math.MaxInt16 + 1, math.MaxInt16},
		{"max uint16 clamps to max", math.MaxUint16, math.MaxInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint16ToInt16(tt.input); got != tt.want {
				t.Errorf("ClampUint16ToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint16ToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint16ToInt8(tt.input); got != tt.want {
				t.Errorf("ClampUint16ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint8ToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
		{"max uint8 clamps to max", math.MaxUint8, math.MaxInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint8ToInt8(tt.input); got != tt.want {
				t.Errorf("ClampUint8ToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  int64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToInt64(tt.input); got != tt.want {
				t.Errorf("ClampUintToInt64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  int32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"overflow clamps to max", math.MaxInt32 + 1, math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToInt32(tt.input); got != tt.want {
				t.Errorf("ClampUintToInt32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToInt16(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  int16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"overflow clamps to max", math.MaxInt16 + 1, math.MaxInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToInt16(tt.input); got != tt.want {
				t.Errorf("ClampUintToInt16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToInt8(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  int8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"overflow clamps to max", math.MaxInt8 + 1, math.MaxInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToInt8(tt.input); got != tt.want {
				t.Errorf("ClampUintToInt8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToInt(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  int
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max int", uint(math.MaxInt), math.MaxInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToInt(tt.input); got != tt.want {
				t.Errorf("ClampUintToInt(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
