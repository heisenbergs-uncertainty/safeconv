package safeconv

import (
	"math"
	"testing"
)

func TestClampUint64ToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
		{"overflow clamps to max", math.MaxUint32 + 1, math.MaxUint32},
		{"max uint64 clamps to max", math.MaxUint64, math.MaxUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToUint32(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint64ToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
		{"overflow clamps to max", math.MaxUint16 + 1, math.MaxUint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToUint16(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint64ToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToUint8(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint64ToUint(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint64ToUint(tt.input); got != tt.want {
				t.Errorf("ClampUint64ToUint(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint32ToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
		{"overflow clamps to max", math.MaxUint16 + 1, math.MaxUint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint32ToUint16(tt.input); got != tt.want {
				t.Errorf("ClampUint32ToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint32ToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint32ToUint8(tt.input); got != tt.want {
				t.Errorf("ClampUint32ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUint16ToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUint16ToUint8(tt.input); got != tt.want {
				t.Errorf("ClampUint16ToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint32
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint32", math.MaxUint32, math.MaxUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToUint32(tt.input); got != tt.want {
				t.Errorf("ClampUintToUint32(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint16
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint16", math.MaxUint16, math.MaxUint16},
		{"overflow clamps to max", math.MaxUint16 + 1, math.MaxUint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToUint16(tt.input); got != tt.want {
				t.Errorf("ClampUintToUint16(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClampUintToUint8(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint8
	}{
		{"zero", 0, 0},
		{"positive in range", 100, 100},
		{"max uint8", math.MaxUint8, math.MaxUint8},
		{"overflow clamps to max", math.MaxUint8 + 1, math.MaxUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClampUintToUint8(tt.input); got != tt.want {
				t.Errorf("ClampUintToUint8(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
