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
}

func TestWriteCombo(t *testing.T) {
	b := NewBStreamWriter(5)
	b.writeBit(one)
	b.writeByte(0xaa)
	log.Println(b.stream[0], b.stream[1])

	c := NewBStreamWriter(5)
	c.WriteBits(0xaa, 8)
	if c.stream[0] != 170 {
		t.Error("write bits wrong.")
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
