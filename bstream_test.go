package bstream

import (
	"log"
	"testing"
)

func TestWriteBit(t *testing.T) {
	b := NewBStreamWriter(5)
	b.writeBit(one)
	if b.stream[0] != 128 {
		t.Error("first bit error")
	}
	b.writeBit(one)
	if b.stream[0] != 192 {
		t.Error("second bit error")
	}

	b.writeBit(one)
	if b.stream[0] != 224 {
		t.Error("third bit error")
	}
}

func TestWriteByte(t *testing.T) {
	b := NewBStreamWriter(5)
	b.writeByte(0xff)
	if b.stream[0] != 255 {
		t.Error("first byte error")
	}
	b.writeByte(0xa0)
	if b.stream[1] != 160 {
		t.Error("second byte error")
	}

	b.writeByte(0x00)
	if b.stream[2] != 0 {
		t.Error("third byte error")
	}
}

func TestWriteCombo(t *testing.T) {
	b := NewBStreamWriter(5)
	b.writeBit(one)
	b.writeByte(0xaa)

	c := NewBStreamWriter(5)
	c.WriteBits(0xaa, 8)
	if c.stream[0] != 170 {
		t.Error("write bits wrong.")
	}

	c.WriteBits(0x0a0a, 8)
	log.Println(c.stream[1])
	if c.stream[1] != 0x0a {
		t.Error("write bit error when too few")
	}

	c.WriteBits(0x0a0a, 16)
	log.Println(c.stream[3])
	if c.stream[3] != 0x0 {
		t.Error("write bit error when too much")
	}
}

func TestReadBit(t *testing.T) {
	b := NewBStreamWriter(5)
	b.WriteBits(0xa0, 8)

	bit, err := b.readBit()

	if err != nil || bit == zero {
		t.Error("Read first bit error")
	}

	bit, err = b.readBit()

	if err != nil || bit == one {
		t.Error("Read second bit error")
	}
}

func TestReadByte(t *testing.T) {
	b := NewBStreamWriter(5)
	b.WriteBits(0xa5a5, 16)

	bit, err := b.readBit()

	if err != nil || bit == zero {
		t.Error("Read first bit error")
	}

	byt, err := b.readByte()
	if byt != 75 {
		t.Error("Read byte error")
	}
}

func TestWriteBits(t *testing.T) {
	b := NewBStreamWriter(24)
	b.WriteBits(0xa5a5, 16)

	ret, err := b.ReadBits(12)
	if err != nil || ret != 2650 {
		t.Error("ReadBits error")
	}

	ret, err = b.ReadBits(4)
	if err != nil || ret != 5 {
		t.Error("ReadBits second error")
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
