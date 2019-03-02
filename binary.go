package nbt

/*
	nbt

 	Copyright (c) 2018 beito

	This software is released under the MIT License.
	http://opensource.org/licenses/mit-license.php
*/

import (
	"github.com/beito123/binary"
)

func readString(stream *binary.OrderStream) (string, error) {
	ln, err := stream.Short()
	if err != nil {
		return "", err
	}

	return string(stream.Get(int(ln))), nil
}

func writeString(stream *binary.OrderStream, str string) error {
	b := []byte(str)

	err := stream.PutShort(uint16(len(b)))
	if err != nil {
		return err
	}

	return stream.Put(b)
}
