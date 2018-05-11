package ioutils

import "io"

type ByteReader struct {
	r io.Reader
}

func NewByteReader(r io.Reader) *ByteReader {
	return &ByteReader{r: r}
}

func (p *ByteReader) ReadByte() (byte, error) {
	b := make([]byte, 1)
	_, err := io.ReadFull(p.r, b)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}
