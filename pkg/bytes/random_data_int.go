//
// Copyright 2018 higherfrequencytrading.com
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
func (data *RandomDataOut) WriteInt8(offset int64, value int8) {
	CheckBounds(offset, 1, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int8)(ptr) = value
}

// WriteInt16 --
func (data *RandomDataOut) WriteInt16(offset int64, value int16) {
	CheckBounds(offset, 2, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int16)(ptr) = value
}

// WriteInt32 --
func (data *RandomDataOut) WriteInt32(offset int64, value int32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int32)(ptr) = value
}

// WriteInt64 --
func (data *RandomDataOut) WriteInt64(offset int64, value int64) {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)

	*(*int64)(ptr) = value
}

// WriteInt32Ordered --
func (data *RandomDataOut) WriteInt32Ordered(offset int64, value int32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)
	dst := (*int32)(ptr)

	atomic.StoreInt32(dst, value)
}

// WriteInt64Ordered --
func (data *RandomDataOut) WriteInt64Ordered(offset int64, value int64) {
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
func (data *RandomDataOut) WriteUnsignedInt8(offset int64, value uint8) {
	CheckBounds(offset, 1, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint8)(ptr) = value
}

// WriteUnsignedInt16 --
func (data *RandomDataOut) WriteUnsignedInt16(offset int64, value uint16) {
	CheckBounds(offset, 2, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint16)(ptr) = value
}

// WriteUnsignedInt32 --
func (data *RandomDataOut) WriteUnsignedInt32(offset int64, value uint32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint32)(ptr) = value
}

// WriteUnsignedInt64 --
func (data *RandomDataOut) WriteUnsignedInt64(offset int64, value uint64) {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)

	*(*uint64)(ptr) = value
}

// WriteUnsignedInt32Ordered --
func (data *RandomDataOut) WriteUnsignedInt32Ordered(offset int64, value uint32) {
	CheckBounds(offset, 4, data.size)

	ptr := data.pointerForOffset(offset)
	dst := (*uint32)(ptr)

	atomic.StoreUint32(dst, value)
}

// WriteUnsignedInt64Ordered --
func (data *RandomDataOut) WriteUnsignedInt64Ordered(offset int64, value uint64) {
	CheckBounds(offset, 8, data.size)

	ptr := data.pointerForOffset(offset)
	dst := (*uint64)(ptr)

	atomic.StoreUint64(dst, value)
}
