package bytes

import (
	"unsafe"
)

// **************************************
//
//
//
// **************************************

type randomDataCommon struct {
	ptr   unsafe.Pointer
	size  int32
	start int32
}

// Capacity --
func (data *randomDataCommon) Capacity() int32 {
	return data.size
}

// Start --
func (data *randomDataCommon) Start() int32 {
	return data.start
}

func (data *randomDataCommon) pointerForOffset(offset int32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(data.ptr) + uintptr(offset))
}

// **************************************
//
//
//
// **************************************

// RandomDataIn --
type RandomDataIn struct {
	randomDataCommon
}

// UnsignedInt8 --
func (data *RandomDataIn) UnsignedInt8(offset int32) uint8 {
	CheckBounds(offset, 1, data.size)

	uptr := data.pointerForOffset(offset)

	return *(*uint8)(uptr)
}

// UnsignedInt16 --
func (data *RandomDataIn) UnsignedInt16(offset int32) uint16 {
	CheckBounds(offset, 2, data.size)

	uptr := data.pointerForOffset(offset)

	return *(*uint16)(uptr)
}

// UnsignedInt32 --
func (data *RandomDataIn) UnsignedInt32(offset int32) uint32 {
	CheckBounds(offset, 4, data.size)

	uptr := data.pointerForOffset(offset)

	return *(*uint32)(uptr)
}

// UnsignedInt64 --
func (data *RandomDataIn) UnsignedInt64(offset int32) uint64 {
	CheckBounds(offset, 8, data.size)

	uptr := data.pointerForOffset(offset)

	return *(*uint64)(uptr)
}

// **************************************
//
//
//
// **************************************

// RandomDataOut --
type RandomDataOut struct {
	randomDataCommon
}
