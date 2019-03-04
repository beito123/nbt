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

// CompressType is a compression type
type CompressType int

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
// If the data is comressed, uncompresses with gzip
func FromFile(path string, order binary.Order) (*Stream, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return FromReader(file, order)
}

// FromReader returns new stream from Reader
// If the data is compressed, uncompresses
func FromReader(reader io.Reader, order binary.Order) (*Stream, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

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
// This often is used for player and level data
func Compress(s *Stream, typ CompressType, level int) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

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
