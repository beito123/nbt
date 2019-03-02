package nbt

/*
	nbt

 	Copyright (c) 2018 beito

	This software is released under the MIT License.
	http://opensource.org/licenses/mit-license.php
*/

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/beito123/binary"
)

func getTagByID(id byte) Tag {
	switch id {
	case IDTagEnd:
		return new(End)
	case IDTagByte:
		return new(Byte)
	case IDTagShort:
		return new(Short)
	case IDTagInt:
		return new(Int)
	case IDTagLong:
		return new(Long)
	case IDTagFloat:
		return new(Float)
	case IDTagDouble:
		return new(Double)
	case IDTagByteArray:
		return new(ByteArray)
	case IDTagString:
		return new(String)
	case IDTagList:
		return new(List)
	case IDTagCompound:
		return new(Compound)
	case IDTagIntArray:
		return new(IntArray)
	}

	return nil
}

// GetTagName returns tag name
func GetTagName(id byte) string {
	switch id {
	case IDTagEnd:
		return "End"
	case IDTagByte:
		return "Byte"
	case IDTagShort:
		return "Short"
	case IDTagInt:
		return "Int"
	case IDTagLong:
		return "Long"
	case IDTagFloat:
		return "Float"
	case IDTagDouble:
		return "Double"
	case IDTagByteArray:
		return "ByteArray"
	case IDTagString:
		return "String"
	case IDTagList:
		return "List"
	case IDTagCompound:
		return "Compound"
	case IDTagIntArray:
		return "IntArray"
	}

	return "Unknown"
}

// Tag is a nbt tag interface
type Tag interface {
	// ID returns tag id
	ID() byte

	// Name returns tag's name
	Name() string

	// SetName set name in tag
	SetName(name string)

	// Read reads tag from Stream
	Read(nbt *Stream) error

	// Write writes tag for Stream
	Write(nbt *Stream) error

	// ToBool returns value as bool
	ToBool() (bool, error)

	// ToByte returns value as byte
	ToByte() (byte, error)

	// ToRune returns value as rune
	ToRune() (rune, error)

	// ToInt returns value as int
	ToInt() (int, error)

	// ToUInt returns value as uint
	ToUInt() (uint, error)

	// ToUInt8 returns value as uint8
	ToUInt8() (uint8, error)

	// ToUInt16 returns value as uint16
	ToUInt16() (uint16, error)

	// ToUInt32 returns value as uint32
	ToUInt32() (uint32, error)

	// ToUInt64 returns value as uint64
	ToUInt64() (uint64, error)

	// ToInt8 returns value as int8
	ToInt8() (int8, error)

	// ToInt16 returns value as int16
	ToInt16() (int16, error)

	// ToInt32 returns value as int32
	ToInt32() (int32, error)

	// ToInt64 returns value as int64
	ToInt64() (int64, error)

	// ToFloat32 returns value as float32
	ToFloat32() (float32, error)

	// ToFloat64 returns value as float64
	ToFloat64() (float64, error)

	// ToByteArray returns value as []byte
	ToByteArray() ([]byte, error)

	// ToString returns value as string
	ToString() (string, error)

	// ToIntArray returns value as []int32
	ToIntArray() ([]int32, error)
}

// End is a end tag
// It shows stream end of tag data
type End struct {
	name string
}

// ID returns tag id
func (t *End) ID() byte {
	return IDTagEnd
}

// Name returns tag's name
func (t *End) Name() string {
	return t.name
}

// SetName set name in tag
func (t *End) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *End) Read(nbt *Stream) error {
	return nil
}

// Write writes tag for Stream
func (t *End) Write(nbt *Stream) error {
	return nil
}

// ToBool returns value as bool
func (t *End) ToBool() (bool, error) {
	return false, nil
}

// ToByte returns value as byte
func (t *End) ToByte() (byte, error) {
	return 0, nil
}

// ToRune returns value as rune
func (t *End) ToRune() (rune, error) {
	return 0, nil
}

// ToInt returns value as int
func (t *End) ToInt() (int, error) {
	return 0, nil
}

// ToUInt returns value as uint
func (t *End) ToUInt() (uint, error) {
	return 0, nil
}

// ToUInt8 returns value as uint8
func (t *End) ToUInt8() (uint8, error) {
	return 0, nil
}

// ToUInt16 returns value as uint16
func (t *End) ToUInt16() (uint16, error) {
	return 0, nil
}

// ToUInt32 returns value as uint32
func (t *End) ToUInt32() (uint32, error) {
	return 0, nil
}

// ToUInt64 returns value as uint64
func (t *End) ToUInt64() (uint64, error) {
	return 0, nil
}

// ToInt8 returns value as int8
func (t *End) ToInt8() (int8, error) {
	return 0, nil
}

// ToInt16 returns value as int16
func (t *End) ToInt16() (int16, error) {
	return 0, nil
}

// ToInt32 returns value as int32
func (t *End) ToInt32() (int32, error) {
	return 0, nil
}

// ToInt64 returns value as int64
func (t *End) ToInt64() (int64, error) {
	return 0, nil
}

// ToFloat32 returns value as float32
func (t *End) ToFloat32() (float32, error) {
	return 0, nil
}

// ToFloat64 returns value as float64
func (t *End) ToFloat64() (float64, error) {
	return 0, nil
}

// ToByteArray returns value as []byte
func (t *End) ToByteArray() ([]byte, error) {
	return []byte{}, nil
}

// ToString returns value as string
func (t *End) ToString() (string, error) {
	return "", nil
}

// ToIntArray returns value as []int32
func (t *End) ToIntArray() ([]int32, error) {
	return []int32{}, nil
}

// Byte is a tag for a byte
type Byte struct {
	name  string
	Value int8
}

// ID returns tag id
func (t *Byte) ID() byte {
	return IDTagByte
}

// Name returns tag's name
func (t *Byte) Name() string {
	return t.name
}

// SetName set name in tag
func (t *Byte) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *Byte) Read(n *Stream) (err error) {
	t.Value, err = n.Stream.SByte()

	return err
}

// Write writes tag for Stream
func (t *Byte) Write(n *Stream) error {
	return n.Stream.PutSByte(t.Value)
}

// ToBool returns value as bool
func (t *Byte) ToBool() (bool, error) {
	return t.Value != 0, nil
}

// ToByte returns value as byte
func (t *Byte) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToRune returns value as rune
func (t *Byte) ToRune() (rune, error) {
	return rune(t.Value), nil
}

// ToInt returns value as int
func (t *Byte) ToInt() (int, error) {
	return int(t.Value), nil
}

// ToUInt returns value as uint
func (t *Byte) ToUInt() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Byte) ToUInt8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Byte) ToUInt16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Byte) ToUInt32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Byte) ToUInt64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToInt8 returns value as int8
func (t *Byte) ToInt8() (int8, error) {
	return t.Value, nil
}

// ToInt16 returns value as int16
func (t *Byte) ToInt16() (int16, error) {
	return int16(t.Value), nil
}

// ToInt32 returns value as int32
func (t *Byte) ToInt32() (int32, error) {
	return int32(t.Value), nil
}

// ToInt64 returns value as int64
func (t *Byte) ToInt64() (int64, error) {
	return int64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Byte) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Byte) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}

// ToByteArray returns value as []byte
func (t *Byte) ToByteArray() ([]byte, error) {
	return []byte{byte(t.Value)}, nil
}

// ToString returns value as string
func (t *Byte) ToString() (string, error) {
	return strconv.Itoa(int(t.Value)), nil
}

// ToIntArray returns value as []int32
func (t *Byte) ToIntArray() ([]int32, error) {
	return []int32{int32(t.Value)}, nil
}

// Short is a tag for a short
type Short struct {
	name  string
	Value int16
}

// ID returns tag id
func (t *Short) ID() byte {
	return IDTagShort
}

// Name returns tag's name
func (t *Short) Name() string {
	return t.name
}

// SetName set name in tag
func (t *Short) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *Short) Read(n *Stream) (err error) {
	t.Value, err = n.Stream.SShort()

	return err
}

// Write writes tag for Stream
func (t *Short) Write(n *Stream) error {
	return n.Stream.PutSShort(t.Value)
}

// ToBool returns value as bool
func (t *Short) ToBool() (bool, error) {
	return t.Value != 0, nil
}

// ToByte returns value as byte
func (t *Short) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToRune returns value as rune
func (t *Short) ToRune() (rune, error) {
	return rune(t.Value), nil
}

// ToInt returns value as int
func (t *Short) ToInt() (int, error) {
	return int(t.Value), nil
}

// ToUInt returns value as uint
func (t *Short) ToUInt() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Short) ToUInt8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Short) ToUInt16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Short) ToUInt32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Short) ToUInt64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToInt8 returns value as int8
func (t *Short) ToInt8() (int8, error) {
	return int8(t.Value), nil
}

// ToInt16 returns value as int16
func (t *Short) ToInt16() (int16, error) {
	return t.Value, nil
}

// ToInt32 returns value as int32
func (t *Short) ToInt32() (int32, error) {
	return int32(t.Value), nil
}

// ToInt64 returns value as int64
func (t *Short) ToInt64() (int64, error) {
	return int64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Short) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Short) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}

// ToByteArray returns value as []byte
func (t *Short) ToByteArray() ([]byte, error) {
	return binary.WriteShort(t.Value), nil
}

// ToString returns value as string
func (t *Short) ToString() (string, error) {
	return strconv.Itoa(int(t.Value)), nil
}

// ToIntArray returns value as []int32
func (t *Short) ToIntArray() ([]int32, error) {
	return []int32{int32(t.Value)}, nil
}

// Int is a tag for an int
type Int struct {
	name  string
	Value int32
}

// ID returns tag id
func (t *Int) ID() byte {
	return IDTagInt
}

// Name returns tag's name
func (t *Int) Name() string {
	return t.name
}

// SetName set name in tag
func (t *Int) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *Int) Read(n *Stream) (err error) {
	t.Value, err = n.Stream.Int()

	return nil
}

// Write writes tag for Stream
func (t *Int) Write(n *Stream) error {
	return n.Stream.PutInt(t.Value)
}

// ToBool returns value as bool
func (t *Int) ToBool() (bool, error) {
	return t.Value != 0, nil
}

// ToByte returns value as byte
func (t *Int) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToRune returns value as rune
func (t *Int) ToRune() (rune, error) {
	return rune(t.Value), nil
}

// ToInt returns value as int
func (t *Int) ToInt() (int, error) {
	return int(t.Value), nil
}

// ToUInt returns value as uint
func (t *Int) ToUInt() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Int) ToUInt8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Int) ToUInt16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Int) ToUInt32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Int) ToUInt64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToInt8 returns value as int8
func (t *Int) ToInt8() (int8, error) {
	return int8(t.Value), nil
}

// ToInt16 returns value as int16
func (t *Int) ToInt16() (int16, error) {
	return int16(t.Value), nil
}

// ToInt32 returns value as int32
func (t *Int) ToInt32() (int32, error) {
	return int32(t.Value), nil
}

// ToInt64 returns value as int64
func (t *Int) ToInt64() (int64, error) {
	return int64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Int) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Int) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}

// ToByteArray returns value as []byte
func (t *Int) ToByteArray() ([]byte, error) {
	return binary.WriteInt(t.Value), nil
}

// ToString returns value as string
func (t *Int) ToString() (string, error) {
	return strconv.Itoa(int(t.Value)), nil
}

// ToIntArray returns value as []int32
func (t *Int) ToIntArray() ([]int32, error) {
	return []int32{int32(t.Value)}, nil
}

// Long is a tag for an int64
type Long struct {
	name  string
	Value int64
}

// ID returns tag id
func (t *Long) ID() byte {
	return IDTagLong
}

// Name returns tag's name
func (t *Long) Name() string {
	return t.name
}

// SetName set name in tag
func (t *Long) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *Long) Read(n *Stream) (err error) {
	t.Value, err = n.Stream.Long()

	return nil
}

// Write writes tag for Stream
func (t *Long) Write(n *Stream) error {
	return n.Stream.PutLong(t.Value)
}

// ToBool returns value as bool
func (t *Long) ToBool() (bool, error) {
	return t.Value != 0, nil
}

// ToByte returns value as byte
func (t *Long) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToRune returns value as rune
func (t *Long) ToRune() (rune, error) {
	return rune(t.Value), nil
}

// ToInt returns value as int
func (t *Long) ToInt() (int, error) {
	return int(t.Value), nil
}

// ToUInt returns value as uint
func (t *Long) ToUInt() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Long) ToUInt8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Long) ToUInt16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Long) ToUInt32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Long) ToUInt64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToInt8 returns value as int8
func (t *Long) ToInt8() (int8, error) {
	return int8(t.Value), nil
}

// ToInt16 returns value as int16
func (t *Long) ToInt16() (int16, error) {
	return int16(t.Value), nil
}

// ToInt32 returns value as int32
func (t *Long) ToInt32() (int32, error) {
	return int32(t.Value), nil
}

// ToInt64 returns value as int64
func (t *Long) ToInt64() (int64, error) {
	return int64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Long) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Long) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}

// ToByteArray returns value as []byte
func (t *Long) ToByteArray() ([]byte, error) {
	return binary.WriteLong(t.Value), nil
}

// ToString returns value as string
func (t *Long) ToString() (string, error) {
	return strconv.FormatInt(t.Value, 10), nil
}

// ToIntArray returns value as []int32
func (t *Long) ToIntArray() ([]int32, error) {
	return []int32{int32(t.Value >> 32), int32(t.Value)}, nil
}

// Float is a tag for a float
type Float struct {
	name  string
	Value float32
}

// ID returns tag id
func (t *Float) ID() byte {
	return IDTagFloat
}

// Name returns tag's name
func (t *Float) Name() string {
	return t.name
}

// SetName set name in tag
func (t *Float) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *Float) Read(n *Stream) (err error) {
	t.Value, err = n.Stream.Float()

	return err
}

// Write writes tag for Stream
func (t *Float) Write(n *Stream) error {
	return n.Stream.PutFloat(t.Value)
}

// ToBool returns value as bool
func (t *Float) ToBool() (bool, error) {
	return t.Value != 0, nil
}

// ToByte returns value as byte
func (t *Float) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToRune returns value as rune
func (t *Float) ToRune() (rune, error) {
	return rune(t.Value), nil
}

// ToInt returns value as int
func (t *Float) ToInt() (int, error) {
	return int(t.Value), nil
}

// ToUInt returns value as uint
func (t *Float) ToUInt() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Float) ToUInt8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Float) ToUInt16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Float) ToUInt32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Float) ToUInt64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToInt8 returns value as int8
func (t *Float) ToInt8() (int8, error) {
	return int8(t.Value), nil
}

// ToInt16 returns value as int16
func (t *Float) ToInt16() (int16, error) {
	return int16(t.Value), nil
}

// ToInt32 returns value as int32
func (t *Float) ToInt32() (int32, error) {
	return int32(t.Value), nil
}

// ToInt64 returns value as int64
func (t *Float) ToInt64() (int64, error) {
	return int64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Float) ToFloat32() (float32, error) {
	return t.Value, nil
}

// ToFloat64 returns value as float64
func (t *Float) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}

// ToByteArray returns value as []byte
func (t *Float) ToByteArray() ([]byte, error) {
	return binary.WriteFloat(t.Value), nil
}

// ToString returns value as string
func (t *Float) ToString() (string, error) {
	return strconv.FormatFloat(float64(t.Value), 'E', -1, 32), nil
}

// ToIntArray returns value as []int32
func (t *Float) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// Double is a tag for float64
type Double struct {
	name  string
	Value float64
}

// ID returns tag id
func (t *Double) ID() byte {
	return IDTagDouble
}

// Name returns tag's name
func (t *Double) Name() string {
	return t.name
}

// SetName set name in tag
func (t *Double) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *Double) Read(n *Stream) (err error) {
	t.Value, err = n.Stream.Double()

	return err
}

// Write writes tag for Stream
func (t *Double) Write(n *Stream) error {
	return n.Stream.PutDouble(t.Value)
}

// ToBool returns value as bool
func (t *Double) ToBool() (bool, error) {
	return t.Value != 0, nil
}

// ToByte returns value as byte
func (t *Double) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToRune returns value as rune
func (t *Double) ToRune() (rune, error) {
	return rune(t.Value), nil
}

// ToInt returns value as int
func (t *Double) ToInt() (int, error) {
	return int(t.Value), nil
}

// ToUInt returns value as uint
func (t *Double) ToUInt() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Double) ToUInt8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Double) ToUInt16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Double) ToUInt32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Double) ToUInt64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToInt8 returns value as int8
func (t *Double) ToInt8() (int8, error) {
	return int8(t.Value), nil
}

// ToInt16 returns value as int16
func (t *Double) ToInt16() (int16, error) {
	return int16(t.Value), nil
}

// ToInt32 returns value as int32
func (t *Double) ToInt32() (int32, error) {
	return int32(t.Value), nil
}

// ToInt64 returns value as int64
func (t *Double) ToInt64() (int64, error) {
	return int64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Double) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Double) ToFloat64() (float64, error) {
	return t.Value, nil
}

// ToByteArray returns value as []byte
func (t *Double) ToByteArray() ([]byte, error) {
	return binary.WriteDouble(t.Value), nil
}

// ToString returns value as string
func (t *Double) ToString() (string, error) {
	return strconv.FormatFloat(t.Value, 'E', -1, 64), nil
}

// ToIntArray returns value as []int32
func (t *Double) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// ByteArray is a tag for []byte
type ByteArray struct {
	name  string
	Value []byte
}

// ID returns tag id
func (t *ByteArray) ID() byte {
	return IDTagByteArray
}

// Name returns tag's name
func (t *ByteArray) Name() string {
	return t.name
}

// SetName set name in tag
func (t *ByteArray) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *ByteArray) Read(n *Stream) error {
	ln, err := n.Stream.Int()

	t.Value = n.Stream.Get(int(ln))

	return err
}

// Write writes tag for Stream
func (t *ByteArray) Write(n *Stream) error {
	return n.Stream.Put(t.Value)
}

// ToBool returns value as bool
func (t *ByteArray) ToBool() (bool, error) {
	return len(t.Value) > 0, nil
}

// ToByte returns value as byte
func (t *ByteArray) ToByte() (byte, error) {
	return 0, errors.New("couldn't cast to byte")
}

// ToRune returns value as rune
func (t *ByteArray) ToRune() (rune, error) {
	return 0, errors.New("couldn't cast to rune")
}

// ToInt returns value as int
func (t *ByteArray) ToInt() (int, error) {
	return 0, errors.New("couldn't cast to int")
}

// ToUInt returns value as uint
func (t *ByteArray) ToUInt() (uint, error) {
	return 0, errors.New("couldn't cast to uint")
}

// ToUInt8 returns value as uint8
func (t *ByteArray) ToUInt8() (uint8, error) {
	return 0, errors.New("couldn't cast to uint8")
}

// ToUInt16 returns value as uint16
func (t *ByteArray) ToUInt16() (uint16, error) {
	return 0, errors.New("couldn't cast to uint16")
}

// ToUInt32 returns value as uint32
func (t *ByteArray) ToUInt32() (uint32, error) {
	return 0, errors.New("couldn't cast to uint32")
}

// ToUInt64 returns value as uint64
func (t *ByteArray) ToUInt64() (uint64, error) {
	return 0, errors.New("couldn't cast to uint64")
}

// ToInt8 returns value as int8
func (t *ByteArray) ToInt8() (int8, error) {
	return 0, errors.New("couldn't cast to int8")
}

// ToInt16 returns value as int16
func (t *ByteArray) ToInt16() (int16, error) {
	return 0, errors.New("couldn't cast to int16")
}

// ToInt32 returns value as int32
func (t *ByteArray) ToInt32() (int32, error) {
	return 0, errors.New("couldn't cast to int32")
}

// ToInt64 returns value as int64
func (t *ByteArray) ToInt64() (int64, error) {
	return 0, errors.New("couldn't cast to int64")
}

// ToFloat32 returns value as float32
func (t *ByteArray) ToFloat32() (float32, error) {
	return 0, errors.New("couldn't cast to float32")
}

// ToFloat64 returns value as float64
func (t *ByteArray) ToFloat64() (float64, error) {
	return 0, errors.New("couldn't cast to float64")
}

// ToByteArray returns value as []byte
func (t *ByteArray) ToByteArray() ([]byte, error) {
	return t.Value, nil
}

// ToString returns value as string
func (t *ByteArray) ToString() (string, error) {
	str := "[ "
	for i, v := range t.Value {
		str += strconv.Itoa(int(v))
		if i != (len(t.Value) - 1) { // not last
			str += ", "
		}
	}

	return str + " ]", nil // output: [10, 255, 25, 100]
}

// ToIntArray returns value as []int32
func (t *ByteArray) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// String is a tag for string
type String struct {
	name  string
	Value string
}

// ID returns tag id
func (t *String) ID() byte {
	return IDTagString
}

// Name returns tag's name
func (t *String) Name() string {
	return t.name
}

// SetName set name in tag
func (t *String) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *String) Read(n *Stream) (err error) {
	t.Value, err = readString(n.Stream)

	return err
}

// Write writes tag for Stream
func (t *String) Write(n *Stream) error {
	return writeString(n.Stream, t.Value)
}

// ToBool returns value as bool
func (t *String) ToBool() (bool, error) {
	return len(t.Value) > 0, nil
}

// ToByte returns value as byte
func (t *String) ToByte() (byte, error) {
	v, err := strconv.Atoi(t.Value)
	return byte(v), err
}

// ToRune returns value as rune
func (t *String) ToRune() (rune, error) {
	v, err := strconv.Atoi(t.Value)
	return rune(v), err
}

// ToInt returns value as int
func (t *String) ToInt() (int, error) {
	v, err := strconv.Atoi(t.Value)
	return v, err
}

// ToUInt returns value as uint
func (t *String) ToUInt() (uint, error) {
	//Thanks: licensed CC BY 3.0 by korthaj
	// from https://yourbasic.org/golang/max-min-int-uint/
	v, err := strconv.ParseUint(t.Value, 10, 32<<(^uint(0)>>63))
	return uint(v), err
}

// ToUInt8 returns value as uint8
func (t *String) ToUInt8() (uint8, error) {
	v, err := strconv.ParseUint(t.Value, 10, 8)
	return uint8(v), err
}

// ToUInt16 returns value as uint16
func (t *String) ToUInt16() (uint16, error) {
	v, err := strconv.ParseUint(t.Value, 10, 16)
	return uint16(v), err
}

// ToUInt32 returns value as uint32
func (t *String) ToUInt32() (uint32, error) {
	v, err := strconv.ParseUint(t.Value, 10, 32)
	return uint32(v), err
}

// ToUInt64 returns value as uint64
func (t *String) ToUInt64() (uint64, error) {
	return strconv.ParseUint(t.Value, 10, 64)
}

// ToInt8 returns value as int8
func (t *String) ToInt8() (int8, error) {
	v, err := strconv.Atoi(t.Value)
	return int8(v), err
}

// ToInt16 returns value as int16
func (t *String) ToInt16() (int16, error) {
	v, err := strconv.Atoi(t.Value)
	return int16(v), err
}

// ToInt32 returns value as int32
func (t *String) ToInt32() (int32, error) {
	v, err := strconv.Atoi(t.Value)
	return int32(v), err
}

// ToInt64 returns value as int64
func (t *String) ToInt64() (int64, error) {
	return strconv.ParseInt(t.Value, 10, 64)
}

// ToFloat32 returns value as float32
func (t *String) ToFloat32() (float32, error) {
	v, err := strconv.ParseFloat(t.Value, 32)
	return float32(v), err
}

// ToFloat64 returns value as float64
func (t *String) ToFloat64() (float64, error) {
	return strconv.ParseFloat(t.Value, 64)
}

// ToByteArray returns value as []byte
func (t *String) ToByteArray() ([]byte, error) {
	return []byte(t.Value), nil
}

// ToString returns value as string
func (t *String) ToString() (string, error) {
	return t.Value, nil
}

// ToIntArray returns value as []int32
func (t *String) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// List is a container for tags
type List struct {
	name  string
	Value []Tag

	ListType byte
}

// ID returns tag id
func (t *List) ID() byte {
	return IDTagList
}

// Name returns tag's name
func (t *List) Name() string {
	return t.name
}

// SetName set name in tag
func (t *List) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *List) Read(n *Stream) (err error) {
	t.ListType, err = n.Stream.Byte()
	if err != nil {
		return err
	}

	ln, err := n.Stream.Int()
	if err != nil {
		return err
	}

	t.Value = make([]Tag, ln)

	for i := 0; i < int(ln); i++ {
		value := getTagByID(t.ListType)
		if value == nil {
			return errors.New("invalid type: " + strconv.Itoa(int(t.ListType)))
		}

		err = value.Read(n)
		if err != nil {
			return errors.New("not enough")
		}

		t.Value[i] = value
	}

	return err
}

// Write writes tag for Stream
func (t *List) Write(n *Stream) error {
	err := n.Stream.PutByte(t.ListType)
	if err != nil {
		return err
	}

	err = n.Stream.PutInt(int32(len(t.Value)))
	if err != nil {
		return err
	}

	for _, v := range t.Value {
		if v.ID() != t.ListType {
			continue // ignore
		}

		err = v.Write(n)
		if err != nil {
			return err
		}
	}

	return err
}

// ToBool returns value as bool
func (t *List) ToBool() (bool, error) {
	return len(t.Value) > 0, nil
}

// ToByte returns value as byte
func (t *List) ToByte() (byte, error) {
	return 0, errors.New("couldn't cast to byte")
}

// ToRune returns value as rune
func (t *List) ToRune() (rune, error) {
	return 0, errors.New("couldn't cast to rune")
}

// ToInt returns value as int
func (t *List) ToInt() (int, error) {
	return 0, errors.New("couldn't cast to int")
}

// ToUInt returns value as uint
func (t *List) ToUInt() (uint, error) {
	return 0, errors.New("couldn't cast to uint")
}

// ToUInt8 returns value as uint8
func (t *List) ToUInt8() (uint8, error) {
	return 0, errors.New("couldn't cast to uint8")
}

// ToUInt16 returns value as uint16
func (t *List) ToUInt16() (uint16, error) {
	return 0, errors.New("couldn't cast to uint16")
}

// ToUInt32 returns value as uint32
func (t *List) ToUInt32() (uint32, error) {
	return 0, errors.New("couldn't cast to uint32")
}

// ToUInt64 returns value as uint64
func (t *List) ToUInt64() (uint64, error) {
	return 0, errors.New("couldn't cast to uint64")
}

// ToInt8 returns value as int8
func (t *List) ToInt8() (int8, error) {
	return 0, errors.New("couldn't cast to int8")
}

// ToInt16 returns value as int16
func (t *List) ToInt16() (int16, error) {
	return 0, errors.New("couldn't cast to int16")
}

// ToInt32 returns value as int32
func (t *List) ToInt32() (int32, error) {
	return 0, errors.New("couldn't cast to int32")
}

// ToInt64 returns value as int64
func (t *List) ToInt64() (int64, error) {
	return 0, errors.New("couldn't cast to int64")
}

// ToFloat32 returns value as float32
func (t *List) ToFloat32() (float32, error) {
	return 0, errors.New("couldn't cast to float32")
}

// ToFloat64 returns value as float64
func (t *List) ToFloat64() (float64, error) {
	return 0, errors.New("couldn't cast to float64")
}

// ToByteArray returns value as []byte
func (t *List) ToByteArray() ([]byte, error) {
	return nil, errors.New("couldn't cast to []byte")
}

// ToString returns value as string
func (t *List) ToString() (string, error) {
	str := GetTagName(t.ListType) + "[ "
	for i, v := range t.Value {
		s, err := v.ToString()
		if err != nil {
			return "", err
		}

		str += s

		if i != (len(t.Value) - 1) { // not last
			str += ", "
		}
	}

	return str + " ]", nil
}

// ToIntArray returns value as []int32
func (t *List) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// Compound is a map contained tags
type Compound struct {
	name string

	Value map[string]Tag
}

// ID returns tag id
func (t *Compound) ID() byte {
	return IDTagCompound
}

// Name returns tag's name
func (t *Compound) Name() string {
	return t.name
}

// SetName set name in tag
func (t *Compound) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *Compound) Read(n *Stream) (err error) {
	t.Value = make(map[string]Tag)

	for {
		tag, err := n.ReadTag()
		if err != nil {
			return err
		}

		if tag.ID() == IDTagEnd {
			break
		}

		t.Value[tag.Name()] = tag
	}

	return err
}

// Write writes tag for Stream
func (t *Compound) Write(n *Stream) error {
	for name, v := range t.Value {
		v.SetName(name)
		err := n.WriteTag(v)
		if err != nil {
			return err
		}
	}

	return n.Stream.PutByte(byte(IDTagEnd))
}

// ToBool returns value as bool
func (t *Compound) ToBool() (bool, error) {
	return len(t.Value) > 0, nil
}

// ToByte returns value as byte
func (t *Compound) ToByte() (byte, error) {
	return 0, errors.New("couldn't cast to byte")
}

// ToRune returns value as rune
func (t *Compound) ToRune() (rune, error) {
	return 0, errors.New("couldn't cast to rune")
}

// ToInt returns value as int
func (t *Compound) ToInt() (int, error) {
	return 0, errors.New("couldn't cast to int")
}

// ToUInt returns value as uint
func (t *Compound) ToUInt() (uint, error) {
	return 0, errors.New("couldn't cast to uint")
}

// ToUInt8 returns value as uint8
func (t *Compound) ToUInt8() (uint8, error) {
	return 0, errors.New("couldn't cast to uint8")
}

// ToUInt16 returns value as uint16
func (t *Compound) ToUInt16() (uint16, error) {
	return 0, errors.New("couldn't cast to uint16")
}

// ToUInt32 returns value as uint32
func (t *Compound) ToUInt32() (uint32, error) {
	return 0, errors.New("couldn't cast to uint32")
}

// ToUInt64 returns value as uint64
func (t *Compound) ToUInt64() (uint64, error) {
	return 0, errors.New("couldn't cast to uint64")
}

// ToInt8 returns value as int8
func (t *Compound) ToInt8() (int8, error) {
	return 0, errors.New("couldn't cast to int8")
}

// ToInt16 returns value as int16
func (t *Compound) ToInt16() (int16, error) {
	return 0, errors.New("couldn't cast to int16")
}

// ToInt32 returns value as int32
func (t *Compound) ToInt32() (int32, error) {
	return 0, errors.New("couldn't cast to int32")
}

// ToInt64 returns value as int64
func (t *Compound) ToInt64() (int64, error) {
	return 0, errors.New("couldn't cast to int64")
}

// ToFloat32 returns value as float32
func (t *Compound) ToFloat32() (float32, error) {
	return 0, errors.New("couldn't cast to float32")
}

// ToFloat64 returns value as float64
func (t *Compound) ToFloat64() (float64, error) {
	return 0, errors.New("couldn't cast to float64")
}

// ToByteArray returns value as []byte
func (t *Compound) ToByteArray() ([]byte, error) {
	return nil, errors.New("couldn't cast to []byte")
}

// ToString returns value as string
func (t *Compound) ToString() (string, error) {
	str := "{ "

	var c int
	for name, v := range t.Value {
		s, err := v.ToString()
		if err != nil {
			return "", err
		}

		str += fmt.Sprintf("%s(%s): %s", name, GetTagName(v.ID()), s)

		c++

		if c != len(t.Value) { // not last
			str += ", "
		}
	}

	return str + " }", nil
}

// ToIntArray returns value as []int32
func (t *Compound) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// Get gets a tag with name from Compound
func (t *Compound) Get(name string) (Tag, bool) {
	tag, ok := t.Value[name]
	return tag, ok
}

// Set set a tag
func (t *Compound) Set(tag Tag) {
	t.Value[tag.Name()] = tag
}

// Has returns whether a tag
func (t *Compound) Has(name string) bool {
	_, ok := t.Value[name]

	return ok
}

// GetBool gets a tag with name as bool
func (t *Compound) GetBool(name string) (bool, error) {
	tag, ok := t.Get(name)
	if !ok {
		return false, errors.New("couldn't find tag " + name)
	}

	return tag.ToBool()
}

// GetByte gets a tag with name as byte
func (t *Compound) GetByte(name string) (byte, error) {
	tag, ok := t.Get(name)
	if !ok {
		return 0, errors.New("couldn't find tag " + name)
	}

	return tag.ToByte()
}

// GetShort gets a tag with name as int16
func (t *Compound) GetShort(name string) (int16, error) {
	tag, ok := t.Get(name)
	if !ok {
		return 0, errors.New("couldn't find tag " + name)
	}

	return tag.ToInt16()
}

// GetInt gets a tag with name as int32
func (t *Compound) GetInt(name string) (int32, error) {
	tag, ok := t.Get(name)
	if !ok {
		return 0, errors.New("couldn't find tag " + name)
	}

	return tag.ToInt32()
}

// GetLong gets a tag with name as int64
func (t *Compound) GetLong(name string) (int64, error) {
	tag, ok := t.Get(name)
	if !ok {
		return 0, errors.New("couldn't find tag " + name)
	}

	return tag.ToInt64()
}

// GetFloat gets a tag with name as float32
func (t *Compound) GetFloat(name string) (float32, error) {
	tag, ok := t.Get(name)
	if !ok {
		return 0, errors.New("couldn't find tag " + name)
	}

	return tag.ToFloat32()
}

// GetDouble gets a tag with name as float64
func (t *Compound) GetDouble(name string) (float64, error) {
	tag, ok := t.Get(name)
	if !ok {
		return 0, errors.New("couldn't find tag " + name)
	}

	return tag.ToFloat64()
}

// GetByteArray gets a tag with name as []byte
func (t *Compound) GetByteArray(name string) ([]byte, error) {
	tag, ok := t.Get(name)
	if !ok {
		return nil, errors.New("couldn't find tag " + name)
	}

	return tag.ToByteArray()
}

// GetString gets a tag with name as string
func (t *Compound) GetString(name string) (string, error) {
	tag, ok := t.Get(name)
	if !ok {
		return "", errors.New("couldn't find tag " + name)
	}

	return tag.ToString()
}

// GetList gets a tag with name as []Tag
func (t *Compound) GetList(name string) ([]Tag, error) {
	tag, ok := t.Get(name)
	if !ok {
		return nil, errors.New("couldn't find tag " + name)
	}

	list, ok := tag.(*List)
	if !ok {
		return nil, errors.New("couldn't cast " + tag.Name() + " List")
	}

	return list.Value, nil
}

// GetCompound gets a tag with name as *Compound
func (t *Compound) GetCompound(name string) (*Compound, error) {
	tag, ok := t.Get(name)
	if !ok {
		return nil, errors.New("couldn't find tag " + name)
	}

	com, ok := tag.(*Compound)
	if !ok {
		return nil, errors.New("couldn't cast " + tag.Name() + " Compound")
	}

	return com, nil
}

// GetIntArray gets a tag with name as []int32
func (t *Compound) GetIntArray(name string) ([]int32, error) {
	tag, ok := t.Get(name)
	if !ok {
		return nil, errors.New("couldn't find tag " + name)
	}

	return tag.ToIntArray()
}

// IntArray is a tag for a int array
type IntArray struct {
	name string

	Value []int32
}

// ID returns tag id
func (t *IntArray) ID() byte {
	return IDTagIntArray
}

// Name returns tag's name
func (t *IntArray) Name() string {
	return t.name
}

// SetName set name in tag
func (t *IntArray) SetName(name string) {
	t.name = name
}

// Read reads tag from Stream
func (t *IntArray) Read(n *Stream) (err error) {
	ln, err := n.Stream.Int()
	if err != nil {
		return err
	}

	for i := 0; i < int(ln); i++ {
		value, err := n.Stream.Int()
		if err != nil {
			return err
		}

		t.Value[i] = value
	}

	return err
}

// Write writes tag for Stream
func (t *IntArray) Write(n *Stream) error {
	err := n.Stream.PutInt(int32(len(t.Value)))
	if err != nil {
		return err
	}

	for _, value := range t.Value {
		err = n.Stream.PutInt(int32(value))
		if err != nil {
			return err
		}
	}

	return err
}

// ToBool returns value as bool
func (t *IntArray) ToBool() (bool, error) {
	return len(t.Value) > 0, nil
}

// ToByte returns value as byte
func (t *IntArray) ToByte() (byte, error) {
	return 0, errors.New("couldn't cast to byte")
}

// ToRune returns value as rune
func (t *IntArray) ToRune() (rune, error) {
	return 0, errors.New("couldn't cast to rune")
}

// ToInt returns value as int
func (t *IntArray) ToInt() (int, error) {
	return 0, errors.New("couldn't cast to int")
}

// ToUInt returns value as uint
func (t *IntArray) ToUInt() (uint, error) {
	return 0, errors.New("couldn't cast to uint")
}

// ToUInt8 returns value as uint8
func (t *IntArray) ToUInt8() (uint8, error) {
	return 0, errors.New("couldn't cast to uint8")
}

// ToUInt16 returns value as uint16
func (t *IntArray) ToUInt16() (uint16, error) {
	return 0, errors.New("couldn't cast to uint16")
}

// ToUInt32 returns value as uint32
func (t *IntArray) ToUInt32() (uint32, error) {
	return 0, errors.New("couldn't cast to uint32")
}

// ToUInt64 returns value as uint64
func (t *IntArray) ToUInt64() (uint64, error) {
	return 0, errors.New("couldn't cast to uint64")
}

// ToInt8 returns value as int8
func (t *IntArray) ToInt8() (int8, error) {
	return 0, errors.New("couldn't cast to int8")
}

// ToInt16 returns value as int16
func (t *IntArray) ToInt16() (int16, error) {
	return 0, errors.New("couldn't cast to int16")
}

// ToInt32 returns value as int32
func (t *IntArray) ToInt32() (int32, error) {
	return 0, errors.New("couldn't cast to int32")
}

// ToInt64 returns value as int64
func (t *IntArray) ToInt64() (int64, error) {
	return 0, errors.New("couldn't cast to int64")
}

// ToFloat32 returns value as float32
func (t *IntArray) ToFloat32() (float32, error) {
	return 0, errors.New("couldn't cast to float32")
}

// ToFloat64 returns value as float64
func (t *IntArray) ToFloat64() (float64, error) {
	return 0, errors.New("couldn't cast to float64")
}

// ToByteArray returns value as []byte
func (t *IntArray) ToByteArray() ([]byte, error) {
	return nil, errors.New("couldn't cast to []byte")
}

// ToString returns value as string
func (t *IntArray) ToString() (string, error) {
	str := "[ "
	for i, v := range t.Value {
		str += strconv.Itoa(int(v))
		if i != (len(t.Value) - 1) { // not last
			str += ", "
		}
	}

	return str + " ]", nil
}

// ToIntArray returns value as []int32
func (t *IntArray) ToIntArray() ([]int32, error) {
	return t.Value, nil
}
