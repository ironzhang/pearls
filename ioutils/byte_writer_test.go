package ioutils

import (
	"bytes"
	"reflect"
	"testing"
)

func TestByteWriter(t *testing.T) {
	ins := []byte{0, 1, 2, 3, 4, 55}
	outs := new(bytes.Buffer)

	w := NewByteWriter(outs)
	for _, b := range ins {
		if err := w.WriteByte(b); err != nil {
			t.Fatalf("write byte: %v", err)
		}
	}
	if got, want := outs.Bytes(), ins; !reflect.DeepEqual(got, want) {
		t.Fatalf("got(%v) != want(%v)", got, want)
	} else {
		t.Logf("got(%v) == want(%v)", got, want)
	}
}
