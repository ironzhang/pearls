package endian

import (
	"bytes"
	"testing"
)

func testVarint(t *testing.T, x int64) {
	var y int64
	var err error
	var buf bytes.Buffer
	if err = EncodeVarint(&buf, x); err != nil {
		t.Errorf("Varint(%d): EncodeVarint: %v", x, err)
		return
	}
	data := buf.Bytes()
	if y, err = DecodeVarint(&buf); err != nil {
		t.Errorf("Varint(%d): DecodeVarint: %v", x, err)
		return
	}
	if x != y {
		t.Errorf("Varint(%d): got %d", x, y)
	} else {
		t.Logf("Varint(%d): got %d, data: %v", x, y, data)
	}
}

func testUvarint(t *testing.T, x uint64) {
	var y uint64
	var err error
	var buf bytes.Buffer
	if err = EncodeUvarint(&buf, x); err != nil {
		t.Errorf("Uvarint(%d): EncodeUvarint: %v", x, err)
		return
	}
	data := buf.Bytes()
	if y, err = DecodeUvarint(&buf); err != nil {
		t.Errorf("Uvarint(%d): EncodeUvarint: %v", x, err)
		return
	}
	if x != y {
		t.Errorf("Uvarint(%d): got %d, want %d", x, y, x)
	} else {
		t.Logf("Uvarint(%d): got %d, data: %v", x, y, data)
	}
}

var tests = []int64{
	-1 << 63,
	-1<<63 + 1,
	-1,
	0,
	1,
	2,
	10,
	20,
	63,
	64,
	65,
	127,
	128,
	129,
	255,
	256,
	257,
	1<<63 - 1,
}

func TestVarint(t *testing.T) {
	for _, x := range tests {
		testVarint(t, x)
	}
}

func TestUvarint(t *testing.T) {
	for _, x := range tests {
		testUvarint(t, uint64(x))
	}
}
