package safeconv

import "testing"

func TestTypeName(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"int", typeName[int](), "int"},
		{"int8", typeName[int8](), "int8"},
		{"int16", typeName[int16](), "int16"},
		{"int32", typeName[int32](), "int32"},
		{"int64", typeName[int64](), "int64"},
		{"uint", typeName[uint](), "uint"},
		{"uint8", typeName[uint8](), "uint8"},
		{"uint16", typeName[uint16](), "uint16"},
		{"uint32", typeName[uint32](), "uint32"},
		{"uint64", typeName[uint64](), "uint64"},
		{"float32", typeName[float32](), "float32"},
		{"float64", typeName[float64](), "float64"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("typeName[%s]() = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}
