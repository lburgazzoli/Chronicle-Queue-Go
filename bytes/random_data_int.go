package bytes

import "sync/atomic"

// **************************************
//
// INT
//
// **************************************

// ReadInt8 --
func (data *RandomDataIn) ReadInt8(offset int64) int8 {
	CheckBounds(offset, 1, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*int8)(ptr)
}

// ReadInt16 --
func (data *RandomDataIn) ReadInt16(offset int64) int16 {
	CheckBounds(offset, 2, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*int16)(ptr)
}

// ReadInt32 --
func (data *RandomDataIn) ReadInt32(offset int64) int32 {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*int32)(ptr)
}

// ReadInt64 --
func (data *RandomDataIn) ReadInt64(offset int64) int64 {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*int64)(ptr)
}

// VolatileReadInt8 --
func (data *RandomDataIn) VolatileReadInt8(offset int64) int8 {
	return data.ReadInt8(offset)
}

// VolatileReadInt16 --
func (data *RandomDataIn) VolatileReadInt16(offset int64) int16 {
	return data.ReadInt16(offset)
}

// VolatileReadInt32 --
func (data *RandomDataIn) VolatileReadInt32(offset int64) int32 {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)
	val := atomic.LoadInt32((*int32)(ptr))

	return int32(val)
}

// VolatileReadInt64 --
func (data *RandomDataIn) VolatileReadInt64(offset int64) int64 {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)
	val := atomic.LoadInt64((*int64)(ptr))

	return int64(val)
}

// WriteInt8 --
func (data *RandomDataIn) WriteInt8(offset int64, value int8) {
	CheckBounds(offset, 1, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int8)(ptr) = value
}

// WriteInt16 --
func (data *RandomDataIn) WriteInt16(offset int64, value int16) {
	CheckBounds(offset, 2, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int16)(ptr) = value
}

// WriteInt32 --
func (data *RandomDataIn) WriteInt32(offset int64, value int32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int32)(ptr) = value
}

// WriteInt64 --
func (data *RandomDataIn) WriteInt64(offset int64, value int64) {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int64)(ptr) = value
}

// WriteInt32Ordered --
func (data *RandomDataIn) WriteInt32Ordered(offset int64, value int32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)
	dst := (*int32)(ptr)

	atomic.StoreInt32(dst, value)
}

// WriteInt64Ordered --
func (data *RandomDataIn) WriteInt64Ordered(offset int64, value int64) {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)
	dst := (*int64)(ptr)

	atomic.StoreInt64(dst, value)
}

// **************************************
//
// UINT
//
// **************************************

// ReadUnsignedInt8 --
func (data *RandomDataIn) ReadUnsignedInt8(offset int64) uint8 {
	CheckBounds(offset, 1, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*uint8)(ptr)
}

// ReadUnsignedInt16 --
func (data *RandomDataIn) ReadUnsignedInt16(offset int64) uint16 {
	CheckBounds(offset, 2, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*uint16)(ptr)
}

// ReadUnsignedInt32 --
func (data *RandomDataIn) ReadUnsignedInt32(offset int64) uint32 {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*uint32)(ptr)
}

// ReadUnsignedInt64 --
func (data *RandomDataIn) ReadUnsignedInt64(offset int64) uint64 {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)

	return *(*uint64)(ptr)
}

// VolatileReadUnsignedInt8 --
func (data *RandomDataIn) VolatileReadUnsignedInt8(offset int64) uint8 {
	return data.ReadUnsignedInt8(offset)
}

// VolatileReadUnsignedInt16 --
func (data *RandomDataIn) VolatileReadUnsignedInt16(offset int64) uint16 {
	return data.ReadUnsignedInt16(offset)
}

// VolatileReadUnsignedInt32 --
func (data *RandomDataIn) VolatileReadUnsignedInt32(offset int64) uint32 {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)
	val := atomic.LoadUint32((*uint32)(ptr))

	return uint32(val)
}

// VolatileReadUnsignedInt64 --
func (data *RandomDataIn) VolatileReadUnsignedInt64(offset int64) uint64 {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)
	val := atomic.LoadUint64((*uint64)(ptr))

	return uint64(val)
}

// WriteUnsignedInt8 --
func (data *RandomDataIn) WriteUnsignedInt8(offset int64, value uint8) {
	CheckBounds(offset, 1, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint8)(ptr) = value
}

// WriteUnsignedInt16 --
func (data *RandomDataIn) WriteUnsignedInt16(offset int64, value uint16) {
	CheckBounds(offset, 2, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint16)(ptr) = value
}

// WriteUnsignedInt32 --
func (data *RandomDataIn) WriteUnsignedInt32(offset int64, value uint32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint32)(ptr) = value
}

// WriteUnsignedInt64 --
func (data *RandomDataIn) WriteUnsignedInt64(offset int64, value uint64) {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint64)(ptr) = value
}

// WriteUnsignedInt32Ordered --
func (data *RandomDataIn) WriteUnsignedInt32Ordered(offset int64, value uint32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)
	dst := (*uint32)(ptr)

	atomic.StoreUint32(dst, value)
}

// WriteUnsignedInt64Ordered --
func (data *RandomDataIn) WriteUnsignedInt64Ordered(offset int64, value uint64) {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)
	dst := (*uint64)(ptr)

	atomic.StoreUint64(dst, value)
}
