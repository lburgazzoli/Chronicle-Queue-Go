package bytes

import "math"

// **************************************
//
// FLOAT
//
// **************************************

// ReadFloat32 --
func (data *RandomDataIn) ReadFloat32(offset int64) float32 {
	val := data.ReadUnsignedInt32(offset)
	ret := math.Float32frombits(val)

	return ret
}

// ReadFloat64 --
func (data *RandomDataIn) ReadFloat64(offset int64) float64 {
	val := data.ReadUnsignedInt64(offset)
	ret := math.Float64frombits(val)

	return ret
}

// VolatileReadFloat32 --
func (data *RandomDataIn) VolatileReadFloat32(offset int64) float32 {
	val := data.VolatileReadUnsignedInt32(offset)
	ret := math.Float32frombits(val)

	return ret
}

// VolatileReadFloat64 --
func (data *RandomDataIn) VolatileReadFloat64(offset int64) float64 {
	val := data.VolatileReadUnsignedInt64(offset)
	ret := math.Float64frombits(val)

	return ret
}

// WriteFloat32 --
func (data *RandomDataOut) WriteFloat32(offset int64, value float32) {
	val := math.Float32bits(value)

	data.WriteUnsignedInt32(offset, val)
}

// WriteFloat64 --
func (data *RandomDataOut) WriteFloat64(offset int64, value float64) {
	val := math.Float64bits(value)

	data.WriteUnsignedInt64(offset, val)
}

// WriteFloat32Ordered --
func (data *RandomDataOut) WriteFloat32Ordered(offset int64, value float32) {
	val := math.Float32bits(value)

	data.WriteUnsignedInt32Ordered(offset, val)
}

// WriteFloat64Ordered --
func (data *RandomDataOut) WriteFloat64Ordered(offset int64, value float64) {
	val := math.Float64bits(value)

	data.WriteUnsignedInt64Ordered(offset, val)
}
