package bstream

import (
	"reflect"
	"testing"
)

func TestWriteBit(t *testing.T) {
	b := NewBStreamWriter(5)
	b.WriteBit(one)
	if b.stream[0] != 128 {
		t.Error("first bit error")
	}
	b.WriteBit(one)
	if b.stream[0] != 192 {
		t.Error("second bit error")
	}

	b.WriteBit(one)
	if b.stream[0] != 224 {
		t.Error("third bit error")
	}
}

func TestWriteOneByte(t *testing.T) {
	b := NewBStreamWriter(5)
	b.WriteOneByte(0xff)
	if b.stream[0] != 255 {
		t.Error("first byte error")
	}
	b.WriteOneByte(0xa0)
	if b.stream[1] != 160 {
		t.Error("second byte error")
	}
	b.WriteOneByte(0x00)
	if b.stream[2] != 0 {
		t.Error("third byte error")
	}
	if l := len(b.Bytes()); l != 3 {
		t.Errorf("Expected %v bytes but got %v", 3, l)
	}
}

func TestWriteCombo(t *testing.T) {
	b := NewBStreamWriter(5)
	b.WriteBit(one)
	b.WriteOneByte(0xaa)
	if b.stream[0] != 0xD5 || b.stream[1] != 0x00 {
		t.Error("write bits wrong")
	}

	c := NewBStreamWriter(5)
	c.WriteBits(0xaa, 8)
	if c.stream[0] != 170 {
		t.Error("write bits wrong.")
	}

	c.WriteBits(0x0a0a, 8)
	if c.stream[1] != 0x0a {
		t.Error("write bit error when too few")
	}

	c.WriteBits(0x0a0a, 16)
	if len(c.stream) > 4 {
		t.Error("write bit error when too much")
	}
}

func TestReadBit(t *testing.T) {
	b := NewBStreamWriter(5)
	b.WriteBits(0xa0, 8)

	bit, err := b.ReadBit()

	if err != nil || bit == zero {
		t.Error("Read first bit error")
	}

	bit, err = b.ReadBit()

	if err != nil || bit == one {
		t.Error("Read second bit error")
	}
}

func TestReadByte(t *testing.T) {
	b := NewBStreamWriter(5)
	b.WriteBits(0xa5a5, 16)

	bit, err := b.ReadBit()

	if err != nil || bit == zero {
		t.Error("Read first bit error")
	}

	byt, err := b.ReadByte()
	if byt != 75 {
		t.Error("Read byte error")
	}
}

func TestWriteBits(t *testing.T) {
	b := NewBStreamWriter(24)
	b.WriteBits(0xa5a5, 16)
	b.WriteBits(0x07, 3)

	ret, err := b.ReadBits(12)
	if err != nil || ret != 2650 {
		t.Error("ReadBits error")
	}
	ret, err = b.ReadBits(4)
	if err != nil || ret != 5 {
		t.Error("ReadBits second error")
	}
	ret, err = b.ReadBits(3)
	if err != nil || ret != 7 {
		t.Error("ReadBits third error")
	}
}

func TestDualMode(t *testing.T) {
	b := NewBStreamWriter(1)

	b.WriteBit(one)
	b.WriteBit(zero)
	b.WriteBit(one)

	if bit, err := b.ReadBit(); err != nil || bit == zero {
		t.Error("Read first bit error")
	}

	b.WriteBit(one)

	if bit, err := b.ReadBit(); err != nil || bit == one {
		t.Error("Read second bit error")
	}
	if bit, err := b.ReadBit(); err != nil || bit == zero {
		t.Error("Read third bit error")
	}

	b.WriteBit(one)

	if bit, err := b.ReadBit(); err != nil || bit == zero {
		t.Error("Read fourth bit error")
	}
	if bit, err := b.ReadBit(); err != nil || bit == zero {
		t.Error("Read fifth bit error")
	}

	if len(b.stream) != 1 || b.stream[0] != 0xB8 {
		t.Error("Wrong stream result")
	}
}

func TestPreserveInput(t *testing.T) {
	var data = []byte{0xAA, 0xAA}
	input := make([]byte, len(data))
	copy(input, data)

	bs := NewBStreamReader(input)

	bs.ReadBits(1)
	bs.ReadByte()

	if !reflect.DeepEqual(data, input) {
		t.Error("Expected input to remain same after reading")
	}
}

func BenchmarkWriteBits(b *testing.B) {
	bb := NewBStreamWriter(255)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bb.WriteBits(uint64(i), 8)
	}
}

func BenchmarkReadBits(b *testing.B) {
	bb := NewBStreamWriter(255)

	for i := 0; i < b.N; i++ {
		bb.WriteBits(uint64(i), 8)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bb.ReadBits(2)
	}
}
