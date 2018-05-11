package endian

import (
	"encoding/binary"
	"io"

	"github.com/ironzhang/pearls/ioutils"
)

func EncodeVarint(w io.Writer, x int64) (err error) {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(b, x)
	_, err = w.Write(b[:n])
	return err
}

func DecodeVarint(r io.Reader) (x int64, err error) {
	if br, ok := r.(io.ByteReader); ok {
		return binary.ReadVarint(br)
	}
	return binary.ReadVarint(ioutils.NewByteReader(r))
}

func EncodeUvarint(w io.Writer, x uint64) (err error) {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(b, x)
	_, err = w.Write(b[:n])
	return err
}

func DecodeUvarint(r io.Reader) (x uint64, err error) {
	if br, ok := r.(io.ByteReader); ok {
		return binary.ReadUvarint(br)
	}
	return binary.ReadUvarint(ioutils.NewByteReader(r))
}
