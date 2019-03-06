package nbt

/*
	nbt

 	Copyright (c) 2018 beito

	This software is released under the MIT License.
	http://opensource.org/licenses/mit-license.php
*/

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"os"

	"github.com/beito123/binary"
)

// DefaultCompressionLevel set the default compression level when it's compressing
const DefaultCompressionLevel = -1

// CompressType is a compression type
type CompressType int

// DefaultCompression returns the default compression level
func (ct CompressType) DefaultCompression() int {
	switch ct {
	case CompressGZip:
		return gzip.DefaultCompression
	case CompressZlib:
		return zlib.DefaultCompression
	}

	return 0
}

const (
	// CompressGZip compresses with gzip
	CompressGZip CompressType = iota

	// CompressZlib compresses with zlib
	CompressZlib
)

// Compression Detector

var gzipHeader = []byte{0x1f, 0x8b, 0x08}
var zlibHeader = []byte{0x78}

func hasGZipHeader(b []byte) bool {
	return bytes.HasPrefix(b, gzipHeader)
}

func hasZlibHeader(b []byte) bool {
	return bytes.HasPrefix(b, zlibHeader)
}

//

// FromFile returns new stream from file
// If the bytes is compressed, it will uncompresses
func FromFile(path string, order binary.Order) (*Stream, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return FromReader(file, order)
}

// FromReader returns new stream from Reader
// If the bytes is compressed, it will uncompresses
func FromReader(reader io.Reader, order binary.Order) (*Stream, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return FromBytes(b, order)
}

// FromBytes returns new stream with bytes
// If the bytes is compressed, it will uncompresses
func FromBytes(b []byte, order binary.Order) (*Stream, error) {
	if hasGZipHeader(b) {
		read, err := gzip.NewReader(bytes.NewBuffer(b))
		if err != nil {
			return nil, err
		}

		defer read.Close()

		b, err = ioutil.ReadAll(read)
		if err != nil {
			return nil, err
		}
	} else if hasZlibHeader(b) {
		read, err := zlib.NewReader(bytes.NewBuffer(b))
		if err != nil {
			return nil, err
		}

		defer read.Close()

		b, err = ioutil.ReadAll(read)
		if err != nil {
			return nil, err
		}
	}

	return NewStreamBytes(order, b), nil
}

// Compress compresses stream's bytes
// You can use compression level in "compress/gzip" and "compress/zlib" for level
// If you set the default compression level, you can set DefaultCompressionLevel
// This often is used for player and level data
func Compress(s *Stream, typ CompressType, level int) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	if level == DefaultCompressionLevel {
		level = typ.DefaultCompression()
	}

	switch typ {
	case CompressGZip:
		writer, err := gzip.NewWriterLevel(buf, level)
		if err != nil {
			return nil, err
		}

		_, err = writer.Write(s.Bytes())
		if err != nil {
			return nil, err
		}

		writer.Flush()
		writer.Close()
	case CompressZlib:
		writer, err := zlib.NewWriterLevel(buf, level)
		if err != nil {
			return nil, err
		}

		_, err = writer.Write(s.Bytes())
		if err != nil {
			return nil, err
		}

		writer.Flush()
		writer.Close()
	}

	return buf.Bytes(), nil
}
