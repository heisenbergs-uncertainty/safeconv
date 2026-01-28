package safeconv

import (
	"errors"
	"testing"
	"unicode"
)

func TestInt64ToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII letter", 65, 'A', nil},
		{"ASCII max", 127, 127, nil},
		{"BMP character", 0x4E2D, '中', nil},      // Chinese character
		{"emoji", 0x1F600, '😀', nil},             // Grinning face emoji
		{"max unicode", 0x10FFFF, 0x10FFFF, nil}, // Maximum valid code point
		{"surrogate start", 0xD800, replacementChar, ErrInvalidUnicode},
		{"surrogate middle", 0xDBFF, replacementChar, ErrInvalidUnicode},
		{"surrogate end", 0xDFFF, replacementChar, ErrInvalidUnicode},
		{"above max unicode", 0x110000, replacementChar, ErrInvalidUnicode},
		{"negative", -1, replacementChar, ErrUnderflow},
		{"large negative", -1000000, replacementChar, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Int64ToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Int64ToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt32ToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"BMP", 0x4E2D, '中', nil},
		{"surrogate", 0xD800, replacementChar, ErrInvalidUnicode},
		{"negative", -1, replacementChar, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int32ToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Int32ToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Int32ToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Int32ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestIntToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"emoji", 0x1F600, '😀', nil},
		{"surrogate", 0xD800, replacementChar, ErrInvalidUnicode},
		{"negative", -1, replacementChar, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("IntToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("IntToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("IntToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt16ToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   int16
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"max int16", 32767, 32767, nil},
		{"negative", -1, replacementChar, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int16ToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Int16ToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Int16ToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Int16ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestInt8ToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   int8
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"max int8", 127, 127, nil},
		{"negative", -1, replacementChar, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int8ToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Int8ToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Int8ToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Int8ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint64ToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"emoji", 0x1F600, '😀', nil},
		{"max unicode", 0x10FFFF, 0x10FFFF, nil},
		{"surrogate start", 0xD800, replacementChar, ErrInvalidUnicode},
		{"surrogate end", 0xDFFF, replacementChar, ErrInvalidUnicode},
		{"above max", 0x110000, replacementChar, ErrInvalidUnicode},
		{"very large", 0xFFFFFFFF, replacementChar, ErrInvalidUnicode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Uint64ToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Uint64ToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint32ToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"max unicode", 0x10FFFF, 0x10FFFF, nil},
		{"surrogate", 0xD800, replacementChar, ErrInvalidUnicode},
		{"above max", 0x110000, replacementChar, ErrInvalidUnicode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint32ToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Uint32ToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Uint32ToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestUintToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"max unicode", 0x10FFFF, 0x10FFFF, nil},
		{"surrogate", 0xD800, replacementChar, ErrInvalidUnicode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UintToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("UintToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("UintToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("UintToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint16ToRune(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		want    rune
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 65, 'A', nil},
		{"BMP max", 0xFFFF, 0xFFFF, nil},
		{"surrogate start", 0xD800, replacementChar, ErrInvalidUnicode},
		{"surrogate end", 0xDFFF, replacementChar, ErrInvalidUnicode},
		{"just before surrogate", 0xD7FF, 0xD7FF, nil},
		{"just after surrogate", 0xE000, 0xE000, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint16ToRune(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Uint16ToRune(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Uint16ToRune(%d) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Uint16ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint8ToRune(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  rune
	}{
		{"zero", 0, 0},
		{"ASCII", 65, 'A'},
		{"max uint8", 255, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToRune(tt.input)
			if got != tt.want {
				t.Errorf("Uint8ToRune(%d) = %U, want %U", tt.input, got, tt.want)
			}
		})
	}
}

func TestRuneToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		want  int64
	}{
		{"zero", 0, 0},
		{"ASCII", 'A', 65},
		{"emoji", '😀', 0x1F600},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RuneToInt64(tt.input)
			if got != tt.want {
				t.Errorf("RuneToInt64(%U) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestRuneToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		want  int32
	}{
		{"zero", 0, 0},
		{"ASCII", 'A', 65},
		{"emoji", '😀', 0x1F600},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RuneToInt32(tt.input)
			if got != tt.want {
				t.Errorf("RuneToInt32(%U) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestRuneToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   rune
		want    uint64
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 'A', 65, nil},
		{"emoji", '😀', 0x1F600, nil},
		{"negative rune", rune(-1), 0, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RuneToUint64(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("RuneToUint64(%U) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("RuneToUint64(%U) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("RuneToUint64(%U) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestRuneToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   rune
		want    uint16
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 'A', 65, nil},
		{"BMP max", 0xFFFF, 0xFFFF, nil},
		{"emoji overflow", '😀', 0, ErrOverflow},
		{"negative rune", rune(-1), 0, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RuneToUint16(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("RuneToUint16(%U) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("RuneToUint16(%U) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("RuneToUint16(%U) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestRuneToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   rune
		want    uint8
		wantErr error
	}{
		{"zero", 0, 0, nil},
		{"ASCII", 'A', 65, nil},
		{"max uint8", 255, 255, nil},
		{"overflow", 256, 0, ErrOverflow},
		{"emoji overflow", '😀', 0, ErrOverflow},
		{"negative rune", rune(-1), 0, ErrUnderflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RuneToUint8(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("RuneToUint8(%U) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("RuneToUint8(%U) unexpected error: %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("RuneToUint8(%U) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsValidRune(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		want  bool
	}{
		{"zero", 0, true},
		{"ASCII", 'A', true},
		{"emoji", '😀', true},
		{"max unicode", 0x10FFFF, true},
		{"surrogate", 0xD800, false},
		{"replacement char", unicode.ReplacementChar, true},
		{"negative", rune(-1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidRune(tt.input); got != tt.want {
				t.Errorf("IsValidRune(%U) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMustInt64ToRune_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustInt64ToRune(-1) should have panicked")
		}
	}()
	MustInt64ToRune(-1)
}

func TestMustUint64ToRune_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustUint64ToRune(0x110000) should have panicked")
		}
	}()
	MustUint64ToRune(0x110000)
}

func TestConversionError_IsInvalidUnicode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"overflow error", ErrOverflow, false},
		{"underflow error", ErrUnderflow, false},
		{"nan error", ErrNaN, false},
		{"infinity error", ErrInfinity, false},
		{"invalid unicode error", ErrInvalidUnicode, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convErr := &ConversionError{
				From:  "int64",
				To:    "rune",
				Value: int64(0),
				Err:   tt.err,
			}
			if got := convErr.IsInvalidUnicode(); got != tt.want {
				t.Errorf("IsInvalidUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
