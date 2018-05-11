package ioutils

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestByteReader(t *testing.T) {
	ins := []byte{0, 1, 2, 3, 4}
	outs := make([]byte, 0)

	r := NewByteReader(bytes.NewBuffer(ins))
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("read byte: %v", err)
		}
		outs = append(outs, b)
	}
	if got, want := outs, ins; !reflect.DeepEqual(got, want) {
		t.Fatalf("got(%v) != want(%v)", got, want)
	} else {
		t.Logf("got(%v) == want(%v)", got, want)
	}
}
