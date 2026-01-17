package safeconv

// must is an internal helper that panics on error.
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// =============================================================================
// Int to Uint Must* variants (from int_to_uint.go)
// =============================================================================

// MustInt64ToUint64 converts int64 to uint64, panicking on error.
func MustInt64ToUint64(x int64) uint64 { return must(Int64ToUint64(x)) }

// MustInt64ToUint32 converts int64 to uint32, panicking on error.
func MustInt64ToUint32(x int64) uint32 { return must(Int64ToUint32(x)) }

// MustInt64ToUint16 converts int64 to uint16, panicking on error.
func MustInt64ToUint16(x int64) uint16 { return must(Int64ToUint16(x)) }

// MustInt64ToUint8 converts int64 to uint8, panicking on error.
func MustInt64ToUint8(x int64) uint8 { return must(Int64ToUint8(x)) }

// MustInt64ToUint converts int64 to uint, panicking on error.
func MustInt64ToUint(x int64) uint { return must(Int64ToUint(x)) }

// MustInt32ToUint64 converts int32 to uint64, panicking on error.
func MustInt32ToUint64(x int32) uint64 { return must(Int32ToUint64(x)) }

// MustInt32ToUint32 converts int32 to uint32, panicking on error.
func MustInt32ToUint32(x int32) uint32 { return must(Int32ToUint32(x)) }

// MustInt32ToUint16 converts int32 to uint16, panicking on error.
func MustInt32ToUint16(x int32) uint16 { return must(Int32ToUint16(x)) }

// MustInt32ToUint8 converts int32 to uint8, panicking on error.
func MustInt32ToUint8(x int32) uint8 { return must(Int32ToUint8(x)) }

// MustInt32ToUint converts int32 to uint, panicking on error.
func MustInt32ToUint(x int32) uint { return must(Int32ToUint(x)) }

// MustInt16ToUint64 converts int16 to uint64, panicking on error.
func MustInt16ToUint64(x int16) uint64 { return must(Int16ToUint64(x)) }

// MustInt16ToUint32 converts int16 to uint32, panicking on error.
func MustInt16ToUint32(x int16) uint32 { return must(Int16ToUint32(x)) }

// MustInt16ToUint16 converts int16 to uint16, panicking on error.
func MustInt16ToUint16(x int16) uint16 { return must(Int16ToUint16(x)) }

// MustInt16ToUint8 converts int16 to uint8, panicking on error.
func MustInt16ToUint8(x int16) uint8 { return must(Int16ToUint8(x)) }

// MustInt16ToUint converts int16 to uint, panicking on error.
func MustInt16ToUint(x int16) uint { return must(Int16ToUint(x)) }

// MustInt8ToUint64 converts int8 to uint64, panicking on error.
func MustInt8ToUint64(x int8) uint64 { return must(Int8ToUint64(x)) }

// MustInt8ToUint32 converts int8 to uint32, panicking on error.
func MustInt8ToUint32(x int8) uint32 { return must(Int8ToUint32(x)) }

// MustInt8ToUint16 converts int8 to uint16, panicking on error.
func MustInt8ToUint16(x int8) uint16 { return must(Int8ToUint16(x)) }

// MustInt8ToUint8 converts int8 to uint8, panicking on error.
func MustInt8ToUint8(x int8) uint8 { return must(Int8ToUint8(x)) }

// MustInt8ToUint converts int8 to uint, panicking on error.
func MustInt8ToUint(x int8) uint { return must(Int8ToUint(x)) }

// MustIntToUint64 converts int to uint64, panicking on error.
func MustIntToUint64(x int) uint64 { return must(IntToUint64(x)) }

// MustIntToUint32 converts int to uint32, panicking on error.
func MustIntToUint32(x int) uint32 { return must(IntToUint32(x)) }

// MustIntToUint16 converts int to uint16, panicking on error.
func MustIntToUint16(x int) uint16 { return must(IntToUint16(x)) }

// MustIntToUint8 converts int to uint8, panicking on error.
func MustIntToUint8(x int) uint8 { return must(IntToUint8(x)) }

// MustIntToUint converts int to uint, panicking on error.
func MustIntToUint(x int) uint { return must(IntToUint(x)) }

// =============================================================================
// Uint to Int Must* variants (from uint_to_int.go)
// =============================================================================

// MustUint64ToInt64 converts uint64 to int64, panicking on error.
func MustUint64ToInt64(x uint64) int64 { return must(Uint64ToInt64(x)) }

// MustUint64ToInt32 converts uint64 to int32, panicking on error.
func MustUint64ToInt32(x uint64) int32 { return must(Uint64ToInt32(x)) }

// MustUint64ToInt16 converts uint64 to int16, panicking on error.
func MustUint64ToInt16(x uint64) int16 { return must(Uint64ToInt16(x)) }

// MustUint64ToInt8 converts uint64 to int8, panicking on error.
func MustUint64ToInt8(x uint64) int8 { return must(Uint64ToInt8(x)) }

// MustUint64ToInt converts uint64 to int, panicking on error.
func MustUint64ToInt(x uint64) int { return must(Uint64ToInt(x)) }

// MustUint32ToInt32 converts uint32 to int32, panicking on error.
func MustUint32ToInt32(x uint32) int32 { return must(Uint32ToInt32(x)) }

// MustUint32ToInt16 converts uint32 to int16, panicking on error.
func MustUint32ToInt16(x uint32) int16 { return must(Uint32ToInt16(x)) }

// MustUint32ToInt8 converts uint32 to int8, panicking on error.
func MustUint32ToInt8(x uint32) int8 { return must(Uint32ToInt8(x)) }

// MustUint16ToInt16 converts uint16 to int16, panicking on error.
func MustUint16ToInt16(x uint16) int16 { return must(Uint16ToInt16(x)) }

// MustUint16ToInt8 converts uint16 to int8, panicking on error.
func MustUint16ToInt8(x uint16) int8 { return must(Uint16ToInt8(x)) }

// MustUint8ToInt8 converts uint8 to int8, panicking on error.
func MustUint8ToInt8(x uint8) int8 { return must(Uint8ToInt8(x)) }

// MustUintToInt64 converts uint to int64, panicking on error.
func MustUintToInt64(x uint) int64 { return must(UintToInt64(x)) }

// MustUintToInt32 converts uint to int32, panicking on error.
func MustUintToInt32(x uint) int32 { return must(UintToInt32(x)) }

// MustUintToInt16 converts uint to int16, panicking on error.
func MustUintToInt16(x uint) int16 { return must(UintToInt16(x)) }

// MustUintToInt8 converts uint to int8, panicking on error.
func MustUintToInt8(x uint) int8 { return must(UintToInt8(x)) }

// MustUintToInt converts uint to int, panicking on error.
func MustUintToInt(x uint) int { return must(UintToInt(x)) }

// =============================================================================
// Int narrowing Must* variants (from int_narrow.go)
// =============================================================================

// MustInt64ToInt32 converts int64 to int32, panicking on error.
func MustInt64ToInt32(x int64) int32 { return must(Int64ToInt32(x)) }

// MustInt64ToInt16 converts int64 to int16, panicking on error.
func MustInt64ToInt16(x int64) int16 { return must(Int64ToInt16(x)) }

// MustInt64ToInt8 converts int64 to int8, panicking on error.
func MustInt64ToInt8(x int64) int8 { return must(Int64ToInt8(x)) }

// MustInt64ToInt converts int64 to int, panicking on error.
func MustInt64ToInt(x int64) int { return must(Int64ToInt(x)) }

// MustInt32ToInt16 converts int32 to int16, panicking on error.
func MustInt32ToInt16(x int32) int16 { return must(Int32ToInt16(x)) }

// MustInt32ToInt8 converts int32 to int8, panicking on error.
func MustInt32ToInt8(x int32) int8 { return must(Int32ToInt8(x)) }

// MustInt16ToInt8 converts int16 to int8, panicking on error.
func MustInt16ToInt8(x int16) int8 { return must(Int16ToInt8(x)) }

// MustIntToInt32 converts int to int32, panicking on error.
func MustIntToInt32(x int) int32 { return must(IntToInt32(x)) }

// MustIntToInt16 converts int to int16, panicking on error.
func MustIntToInt16(x int) int16 { return must(IntToInt16(x)) }

// MustIntToInt8 converts int to int8, panicking on error.
func MustIntToInt8(x int) int8 { return must(IntToInt8(x)) }

// =============================================================================
// Uint narrowing Must* variants (from uint_narrow.go)
// =============================================================================

// MustUint64ToUint32 converts uint64 to uint32, panicking on error.
func MustUint64ToUint32(x uint64) uint32 { return must(Uint64ToUint32(x)) }

// MustUint64ToUint16 converts uint64 to uint16, panicking on error.
func MustUint64ToUint16(x uint64) uint16 { return must(Uint64ToUint16(x)) }

// MustUint64ToUint8 converts uint64 to uint8, panicking on error.
func MustUint64ToUint8(x uint64) uint8 { return must(Uint64ToUint8(x)) }

// MustUint64ToUint converts uint64 to uint, panicking on error.
func MustUint64ToUint(x uint64) uint { return must(Uint64ToUint(x)) }

// MustUint32ToUint16 converts uint32 to uint16, panicking on error.
func MustUint32ToUint16(x uint32) uint16 { return must(Uint32ToUint16(x)) }

// MustUint32ToUint8 converts uint32 to uint8, panicking on error.
func MustUint32ToUint8(x uint32) uint8 { return must(Uint32ToUint8(x)) }

// MustUint16ToUint8 converts uint16 to uint8, panicking on error.
func MustUint16ToUint8(x uint16) uint8 { return must(Uint16ToUint8(x)) }

// MustUintToUint32 converts uint to uint32, panicking on error.
func MustUintToUint32(x uint) uint32 { return must(UintToUint32(x)) }

// MustUintToUint16 converts uint to uint16, panicking on error.
func MustUintToUint16(x uint) uint16 { return must(UintToUint16(x)) }

// MustUintToUint8 converts uint to uint8, panicking on error.
func MustUintToUint8(x uint) uint8 { return must(UintToUint8(x)) }

// =============================================================================
// Float64 truncation Must* variants (from float_trunc.go)
// =============================================================================

// MustFloat64TruncToInt64 converts float64 to int64 by truncating, panicking on error.
func MustFloat64TruncToInt64(x float64) int64 { return must(Float64TruncToInt64(x)) }

// MustFloat64TruncToInt32 converts float64 to int32 by truncating, panicking on error.
func MustFloat64TruncToInt32(x float64) int32 { return must(Float64TruncToInt32(x)) }

// MustFloat64TruncToInt16 converts float64 to int16 by truncating, panicking on error.
func MustFloat64TruncToInt16(x float64) int16 { return must(Float64TruncToInt16(x)) }

// MustFloat64TruncToInt8 converts float64 to int8 by truncating, panicking on error.
func MustFloat64TruncToInt8(x float64) int8 { return must(Float64TruncToInt8(x)) }

// MustFloat64TruncToInt converts float64 to int by truncating, panicking on error.
func MustFloat64TruncToInt(x float64) int { return must(Float64TruncToInt(x)) }

// MustFloat64TruncToUint64 converts float64 to uint64 by truncating, panicking on error.
func MustFloat64TruncToUint64(x float64) uint64 { return must(Float64TruncToUint64(x)) }

// MustFloat64TruncToUint32 converts float64 to uint32 by truncating, panicking on error.
func MustFloat64TruncToUint32(x float64) uint32 { return must(Float64TruncToUint32(x)) }

// MustFloat64TruncToUint16 converts float64 to uint16 by truncating, panicking on error.
func MustFloat64TruncToUint16(x float64) uint16 { return must(Float64TruncToUint16(x)) }

// MustFloat64TruncToUint8 converts float64 to uint8 by truncating, panicking on error.
func MustFloat64TruncToUint8(x float64) uint8 { return must(Float64TruncToUint8(x)) }

// MustFloat64TruncToUint converts float64 to uint by truncating, panicking on error.
func MustFloat64TruncToUint(x float64) uint { return must(Float64TruncToUint(x)) }

// =============================================================================
// Float32 truncation Must* variants (from float_trunc.go)
// =============================================================================

// MustFloat32TruncToInt64 converts float32 to int64 by truncating, panicking on error.
func MustFloat32TruncToInt64(x float32) int64 { return must(Float32TruncToInt64(x)) }

// MustFloat32TruncToInt32 converts float32 to int32 by truncating, panicking on error.
func MustFloat32TruncToInt32(x float32) int32 { return must(Float32TruncToInt32(x)) }

// MustFloat32TruncToInt16 converts float32 to int16 by truncating, panicking on error.
func MustFloat32TruncToInt16(x float32) int16 { return must(Float32TruncToInt16(x)) }

// MustFloat32TruncToInt8 converts float32 to int8 by truncating, panicking on error.
func MustFloat32TruncToInt8(x float32) int8 { return must(Float32TruncToInt8(x)) }

// MustFloat32TruncToInt converts float32 to int by truncating, panicking on error.
func MustFloat32TruncToInt(x float32) int { return must(Float32TruncToInt(x)) }

// MustFloat32TruncToUint64 converts float32 to uint64 by truncating, panicking on error.
func MustFloat32TruncToUint64(x float32) uint64 { return must(Float32TruncToUint64(x)) }

// MustFloat32TruncToUint32 converts float32 to uint32 by truncating, panicking on error.
func MustFloat32TruncToUint32(x float32) uint32 { return must(Float32TruncToUint32(x)) }

// MustFloat32TruncToUint16 converts float32 to uint16 by truncating, panicking on error.
func MustFloat32TruncToUint16(x float32) uint16 { return must(Float32TruncToUint16(x)) }

// MustFloat32TruncToUint8 converts float32 to uint8 by truncating, panicking on error.
func MustFloat32TruncToUint8(x float32) uint8 { return must(Float32TruncToUint8(x)) }

// MustFloat32TruncToUint converts float32 to uint by truncating, panicking on error.
func MustFloat32TruncToUint(x float32) uint { return must(Float32TruncToUint(x)) }

// =============================================================================
// Float64 rounding Must* variants (from float_round.go)
// =============================================================================

// MustFloat64RoundToInt64 converts float64 to int64 by rounding, panicking on error.
func MustFloat64RoundToInt64(x float64) int64 { return must(Float64RoundToInt64(x)) }

// MustFloat64RoundToInt32 converts float64 to int32 by rounding, panicking on error.
func MustFloat64RoundToInt32(x float64) int32 { return must(Float64RoundToInt32(x)) }

// MustFloat64RoundToInt16 converts float64 to int16 by rounding, panicking on error.
func MustFloat64RoundToInt16(x float64) int16 { return must(Float64RoundToInt16(x)) }

// MustFloat64RoundToInt8 converts float64 to int8 by rounding, panicking on error.
func MustFloat64RoundToInt8(x float64) int8 { return must(Float64RoundToInt8(x)) }

// MustFloat64RoundToInt converts float64 to int by rounding, panicking on error.
func MustFloat64RoundToInt(x float64) int { return must(Float64RoundToInt(x)) }

// MustFloat64RoundToUint64 converts float64 to uint64 by rounding, panicking on error.
func MustFloat64RoundToUint64(x float64) uint64 { return must(Float64RoundToUint64(x)) }

// MustFloat64RoundToUint32 converts float64 to uint32 by rounding, panicking on error.
func MustFloat64RoundToUint32(x float64) uint32 { return must(Float64RoundToUint32(x)) }

// MustFloat64RoundToUint16 converts float64 to uint16 by rounding, panicking on error.
func MustFloat64RoundToUint16(x float64) uint16 { return must(Float64RoundToUint16(x)) }

// MustFloat64RoundToUint8 converts float64 to uint8 by rounding, panicking on error.
func MustFloat64RoundToUint8(x float64) uint8 { return must(Float64RoundToUint8(x)) }

// MustFloat64RoundToUint converts float64 to uint by rounding, panicking on error.
func MustFloat64RoundToUint(x float64) uint { return must(Float64RoundToUint(x)) }

// =============================================================================
// Float32 rounding Must* variants (from float_round.go)
// =============================================================================

// MustFloat32RoundToInt64 converts float32 to int64 by rounding, panicking on error.
func MustFloat32RoundToInt64(x float32) int64 { return must(Float32RoundToInt64(x)) }

// MustFloat32RoundToInt32 converts float32 to int32 by rounding, panicking on error.
func MustFloat32RoundToInt32(x float32) int32 { return must(Float32RoundToInt32(x)) }

// MustFloat32RoundToInt16 converts float32 to int16 by rounding, panicking on error.
func MustFloat32RoundToInt16(x float32) int16 { return must(Float32RoundToInt16(x)) }

// MustFloat32RoundToInt8 converts float32 to int8 by rounding, panicking on error.
func MustFloat32RoundToInt8(x float32) int8 { return must(Float32RoundToInt8(x)) }

// MustFloat32RoundToInt converts float32 to int by rounding, panicking on error.
func MustFloat32RoundToInt(x float32) int { return must(Float32RoundToInt(x)) }

// MustFloat32RoundToUint64 converts float32 to uint64 by rounding, panicking on error.
func MustFloat32RoundToUint64(x float32) uint64 { return must(Float32RoundToUint64(x)) }

// MustFloat32RoundToUint32 converts float32 to uint32 by rounding, panicking on error.
func MustFloat32RoundToUint32(x float32) uint32 { return must(Float32RoundToUint32(x)) }

// MustFloat32RoundToUint16 converts float32 to uint16 by rounding, panicking on error.
func MustFloat32RoundToUint16(x float32) uint16 { return must(Float32RoundToUint16(x)) }

// MustFloat32RoundToUint8 converts float32 to uint8 by rounding, panicking on error.
func MustFloat32RoundToUint8(x float32) uint8 { return must(Float32RoundToUint8(x)) }

// MustFloat32RoundToUint converts float32 to uint by rounding, panicking on error.
func MustFloat32RoundToUint(x float32) uint { return must(Float32RoundToUint(x)) }

// =============================================================================
// Float narrowing Must* variants (from float_narrow.go)
// =============================================================================

// MustFloat64ToFloat32 converts float64 to float32, panicking on error.
func MustFloat64ToFloat32(x float64) float32 { return must(Float64ToFloat32(x)) }
