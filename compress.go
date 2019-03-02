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
	"io"
	"io/ioutil"
	"os"

	"github.com/beito123/binary"
)

var gzipHeader = []byte{0x1f, 0x8b, 0x08}

func hasGZipHeader(b []byte) bool {
	return bytes.HasPrefix(b, gzipHeader)
}

// FromFile returns new stream from file
// If the data is comressed, uncompresses with gzip
func FromFile(path string, order binary.Order) (*Stream, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return FromReader(file, order)
}

// FromReader returns new stream from Reader
// If the data is compressed, uncompresses with gzip
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
	}

	return NewStreamBytes(order, b), nil
}

// Compress compresses Stream's bytes
// This often is used for player and level data
func Compress(s *Stream) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	writer := gzip.NewWriter(buf)

	_, err := writer.Write(s.Bytes())
	if err != nil {
		return nil, err
	}

	writer.Flush()
	writer.Close()

	return buf.Bytes(), nil
}
