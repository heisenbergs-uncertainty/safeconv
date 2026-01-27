package safeconv

import (
	"math"
	"testing"
)

func TestClampInt64ToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int64", math.MaxInt64, math.MaxInt64},
		{"negative clamps to zero", -1, 0},
		{"min int64 clamps to zero", math.MinInt64, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToUint64(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt64ToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
		{"overflow clamps to max", math.MaxUint32 + 1, math.MaxUint32},
		{"large overflow clamps to max", math.MaxInt64, math.MaxUint32},
		{"negative clamps to zero", -1, 0},
		{"min int64 clamps to zero", math.MinInt64, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToUint32(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt64ToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
		{"overflow clamps to max", math.MaxUint16 + 1, math.MaxUint16},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToUint16(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt64ToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToUint8(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt64ToUint(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  uint
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt64ToUint(tt.input); got != tt.want {
				t.Errorf("ClampInt64ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt32ToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt32ToUint64(tt.input); got != tt.want {
				t.Errorf("ClampInt32ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt32ToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int32", math.MaxInt32, math.MaxInt32},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt32ToUint32(tt.input); got != tt.want {
				t.Errorf("ClampInt32ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt32ToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
		{"overflow clamps to max", math.MaxUint16 + 1, math.MaxUint16},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt32ToUint16(tt.input); got != tt.want {
				t.Errorf("ClampInt32ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt32ToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt32ToUint8(tt.input); got != tt.want {
				t.Errorf("ClampInt32ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt32ToUint(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  uint
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt32ToUint(tt.input); got != tt.want {
				t.Errorf("ClampInt32ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt16ToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt16ToUint64(tt.input); got != tt.want {
				t.Errorf("ClampInt16ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt16ToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt16ToUint32(tt.input); got != tt.want {
				t.Errorf("ClampInt16ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt16ToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int16", math.MaxInt16, math.MaxInt16},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt16ToUint16(tt.input); got != tt.want {
				t.Errorf("ClampInt16ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt16ToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt16ToUint8(tt.input); got != tt.want {
				t.Errorf("ClampInt16ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt16ToUint(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  uint
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt16ToUint(tt.input); got != tt.want {
				t.Errorf("ClampInt16ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt8ToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt8ToUint64(tt.input); got != tt.want {
				t.Errorf("ClampInt8ToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt8ToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt8ToUint32(tt.input); got != tt.want {
				t.Errorf("ClampInt8ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt8ToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt8ToUint16(tt.input); got != tt.want {
				t.Errorf("ClampInt8ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt8ToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"max int8", math.MaxInt8, math.MaxInt8},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt8ToUint8(tt.input); got != tt.want {
				t.Errorf("ClampInt8ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampInt8ToUint(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  uint
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampInt8ToUint(tt.input); got != tt.want {
				t.Errorf("ClampInt8ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  uint64
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToUint64(tt.input); got != tt.want {
				t.Errorf("ClampIntToUint64(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToUint32(tt.input); got != tt.want {
				t.Errorf("ClampIntToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
		{"overflow clamps to max", math.MaxUint16 + 1, math.MaxUint16},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToUint16(tt.input); got != tt.want {
				t.Errorf("ClampIntToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToUint8(tt.input); got != tt.want {
				t.Errorf("ClampIntToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampIntToUint(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  uint
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
		{"negative clamps to zero", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampIntToUint(tt.input); got != tt.want {
				t.Errorf("ClampIntToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
