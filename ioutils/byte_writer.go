package ioutils

import "io"

type ByteWriter struct {
	w io.Writer
}

func NewByteWriter(w io.Writer) *ByteWriter {
	return &ByteWriter{w: w}
}

func (p *ByteWriter) WriteByte(b byte) error {
	_, err := p.w.Write([]byte{b})
	return err
}
