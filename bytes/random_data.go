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
	size  int64
	start int64
}

// Capacity --
func (data *randomDataCommon) Capacity() int64 {
	return data.size
}

// Start --
func (data *randomDataCommon) Start() int64 {
	return data.start
}

func (data *randomDataCommon) pointerForOffset(offset int64) unsafe.Pointer {
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

// **************************************
//
//
//
// **************************************

// RandomDataOut --
type RandomDataOut struct {
	randomDataCommon
}
