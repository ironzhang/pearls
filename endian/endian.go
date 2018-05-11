package endian

import (
	"encoding/binary"
	"io"

	"github.com/ironzhang/pearls/ioutils"
)

type ByteOrder interface {
	String() string
	EncodeUint8(w io.Writer, x uint8) (err error)
	DecodeUint8(r io.Reader) (x uint8, err error)
	EncodeUint16(w io.Writer, x uint16) (err error)
	DecodeUint16(r io.Reader) (x uint16, err error)
	EncodeUint32(w io.Writer, x uint32) (err error)
	DecodeUint32(r io.Reader) (x uint32, err error)
	EncodeUint64(w io.Writer, x uint64) (err error)
	DecodeUint64(r io.Reader) (x uint64, err error)
}

// LittleEndian 小端字节序编解码器
var LittleEndian littleEndian

// BigEndian 大端字节序编解码器
var BigEndian bigEndian

type littleEndian struct{}

func (littleEndian) String() string { return "LittleEndian" }

func (littleEndian) EncodeUint8(w io.Writer, x uint8) (err error) {
	if bw, ok := w.(io.ByteWriter); ok {
		return bw.WriteByte(x)
	}
	return ioutils.NewByteWriter(w).WriteByte(x)
}

func (littleEndian) DecodeUint8(r io.Reader) (x uint8, err error) {
	if br, ok := r.(io.ByteReader); ok {
		return br.ReadByte()
	}
	return ioutils.NewByteReader(r).ReadByte()
}

func (littleEndian) EncodeUint16(w io.Writer, x uint16) (err error) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, x)
	_, err = w.Write(b)
	return err
}

func (littleEndian) DecodeUint16(r io.Reader) (x uint16, err error) {
	b := make([]byte, 2)
	if _, err = io.ReadFull(r, b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b), nil
}

func (littleEndian) EncodeUint32(w io.Writer, x uint32) (err error) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, x)
	_, err = w.Write(b)
	return err
}

func (littleEndian) DecodeUint32(r io.Reader) (x uint32, err error) {
	b := make([]byte, 4)
	if _, err = io.ReadFull(r, b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}

func (littleEndian) EncodeUint64(w io.Writer, x uint64) (err error) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, x)
	_, err = w.Write(b)
	return err
}

func (littleEndian) DecodeUint64(r io.Reader) (x uint64, err error) {
	b := make([]byte, 8)
	if _, err = io.ReadFull(r, b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b), nil
}

type bigEndian struct{}

func (bigEndian) String() string { return "BigEndian" }

func (bigEndian) EncodeUint8(w io.Writer, x uint8) (err error) {
	if bw, ok := w.(io.ByteWriter); ok {
		return bw.WriteByte(x)
	}
	return ioutils.NewByteWriter(w).WriteByte(x)
}

func (bigEndian) DecodeUint8(r io.Reader) (x uint8, err error) {
	if br, ok := r.(io.ByteReader); ok {
		return br.ReadByte()
	}
	return ioutils.NewByteReader(r).ReadByte()
}

func (bigEndian) EncodeUint16(w io.Writer, x uint16) (err error) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, x)
	_, err = w.Write(b)
	return err
}

func (bigEndian) DecodeUint16(r io.Reader) (x uint16, err error) {
	b := make([]byte, 2)
	if _, err = io.ReadFull(r, b); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(b), nil
}

func (bigEndian) EncodeUint32(w io.Writer, x uint32) (err error) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, x)
	_, err = w.Write(b)
	return err
}

func (bigEndian) DecodeUint32(r io.Reader) (x uint32, err error) {
	b := make([]byte, 4)
	if _, err = io.ReadFull(r, b); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(b), nil
}

func (bigEndian) EncodeUint64(w io.Writer, x uint64) (err error) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, x)
	_, err = w.Write(b)
	return err
}

func (bigEndian) DecodeUint64(r io.Reader) (x uint64, err error) {
	b := make([]byte, 8)
	if _, err = io.ReadFull(r, b); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b), nil
}
