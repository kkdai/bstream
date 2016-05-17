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

}
