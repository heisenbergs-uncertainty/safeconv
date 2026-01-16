package safeconv

// SignedInt is a constraint for signed integer types.
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInt is a constraint for unsigned integer types.
type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer is a constraint for all integer types.
type Integer interface {
	SignedInt | UnsignedInt
}

// Float is a constraint for floating-point types.
type Float interface {
	~float32 | ~float64
}
