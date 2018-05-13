package endian

import (
	"bytes"
	"testing"
)

type ST struct {
	u8  uint8
	u16 uint16
	u32 uint32
	u64 uint64
}

func testEndian(t *testing.T, b ByteOrder) {
	var err error
	tests := []ST{
		{
			u8:  8,
			u16: 16,
			u32: 32,
			u64: 64,
		},
	}
	for i, tt := range tests {
		var buf bytes.Buffer
		if err = b.EncodeUint8(&buf, tt.u8); err != nil {
			t.Fatalf("%d: EncodeUint8: %v", i, err)
		}
		if err = b.EncodeUint16(&buf, tt.u16); err != nil {
			t.Fatalf("%d: EncodeUint16: %v", i, err)
		}
		if err = b.EncodeUint32(&buf, tt.u32); err != nil {
			t.Fatalf("%d: EncodeUint32: %v", i, err)
		}
		if err = b.EncodeUint64(&buf, tt.u64); err != nil {
			t.Fatalf("%d: EncodeUint64: %v", i, err)
		}
		data := buf.Bytes()

		var got ST
		if got.u8, err = b.DecodeUint8(&buf); err != nil {
			t.Fatalf("%d: DecodeUint8: %v", i, err)
		}
		if got.u16, err = b.DecodeUint16(&buf); err != nil {
			t.Fatalf("%d: DecodeUint16: %v", i, err)
		}
		if got.u32, err = b.DecodeUint32(&buf); err != nil {
			t.Fatalf("%d: DecodeUint32: %v", i, err)
		}
		if got.u64, err = b.DecodeUint64(&buf); err != nil {
			t.Fatalf("%d: DecodeUint64: %v", i, err)
		}
		if want := tt; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v, data: %v", i, got, data)
		}
	}
}

func TestLittleEndian(t *testing.T) {
	testEndian(t, LittleEndian)
}

func TestBigEndian(t *testing.T) {
	testEndian(t, BigEndian)
}
