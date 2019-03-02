package nbt

/*
	mc-nbt

 	Copyright (c) 2019 beito

	This software is released under the MIT License.
	http://opensource.org/licenses/mit-license.php
*/

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/beito123/binary"
)

const (
	// IDTagEnd signifies end for compound tag
	// Payload: none
	IDTagEnd = iota

	// IDTagByte is a signed byte (-128 to 127)
	// Payload is 1byte
	IDTagByte

	// IDTagShort is a signed short (-32768 to 32767)
	// Payload is 2 bytes
	IDTagShort

	// IDTagInt is a signed int (-2147483648 to 2147483647)
	// Payload is 4 bytes
	IDTagInt

	// IDTagLong is a signed long (-9223372036854775808 to 9223372036854775807)
	// Payload is 8 bytes
	IDTagLong

	// IDTagFloat is a signed single float (IEEE-754)
	// Payload is 4 bytes
	IDTagFloat

	// IDTagDouble is a signed single double (IEEE-754)
	// Payload is 8 bytes
	IDTagDouble

	// IDTagByteArray is a array of signed bytes
	// Payload is 4 bytes(len of data with signed int) + len bytes
	IDTagByteArray

	// IDTagString is a UTF-8 string (max 32767 bytes)
	// Payload is 2 bytes (len of data with short) + len bytes
	IDTagString

	// IDTagList is a list of nameless tags, all tags need same type.
	// Payload is 1 byte(tag type with byte) + 4 bytes(len with signed int) + len bytes
	IDTagList

	// IDTagCompound is a list of named tags.
	IDTagCompound

	// IDTagIntArray is a list of int.
	IDTagIntArray
)

var (
	// BigEndian is for MCJE, anvil and more...
	BigEndian = binary.BigEndian

	// LittleEndian is for MCBE leveldb format
	LittleEndian = binary.LittleEndian
)

// NewStream returns new Stream
func NewStream(order binary.Order) *Stream {
	return NewStreamBytes(order, []byte{})
}

// NewStreamBytes returns new Stream with bytes data
func NewStreamBytes(order binary.Order, b []byte) *Stream {
	return &Stream{
		Stream: binary.NewOrderStreamBytes(order, b),
	}
}

// Stream is binary nbt stream
type Stream struct {
	Stream *binary.OrderStream
}

// Reset resets buffer
func (s *Stream) Reset() {
	s.Stream.Reset()
}

// Bytes returns stream's all buffer
func (s *Stream) Bytes() []byte {
	return s.Stream.AllBytes()
}

// ReadTag reads tag from buffer
func (s *Stream) ReadTag() (Tag, error) {
	errHandle := func(err error) error {
		return fmt.Errorf("nbt: happened errors while it's parsing near %d Error: %s", s.Stream.Off(), err.Error())
	}

	id, err := s.Stream.Byte()
	if err != nil {
		return nil, errHandle(err)
	}

	tag := getTagByID(id)
	if tag == nil {
		return nil, errors.New("mc.nbt: invalid tag, " + strconv.Itoa(int(id)))
	}

	if tag.ID() == IDTagEnd {
		return tag, nil
	}

	name, err := readString(s.Stream)
	if err != nil {
		return nil, errHandle(err)
	}

	tag.SetName(name)

	err = tag.Read(s)
	if err != nil {
		return nil, errHandle(err)
	}

	return tag, nil
}

// WriteTag writes tag to buffer
func (s *Stream) WriteTag(tag Tag) error {
	err := s.Stream.PutByte(tag.ID())
	if err != nil {
		return err
	}

	err = writeString(s.Stream, tag.Name())
	if err != nil {
		return err
	}

	return tag.Write(s)
}
