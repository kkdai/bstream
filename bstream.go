package bstream

type bit bool

const (
	zero bit = false
	one  bit = true
)

type BStream struct {
	stream      []byte
	remainCount uint8
}

//NewBStreamReader :
func NewBStreamReader(data []byte) *BStream {
	return &BStream{stream: data, remainCount: 8}
}

//NewBStreamWriter :
func NewBStreamWriter(length uint8) *BStream {
	return &BStream{stream: make([]byte, 0, length), remainCount: 0}
}

//WriteBit :
func (b *BStream) writeBit(input bit) {
	if b.remainCount == 0 {
		b.stream = append(b.stream, 0)
		b.remainCount = 8
	}

	latestIndex := len(b.stream) - 1
	if input {
		b.stream[latestIndex] |= 1 << (b.remainCount - 1)
	}
	b.remainCount--
}

//WriteByte :
func (b *BStream) writeByte(data byte) {
	if b.remainCount == 0 {
		b.stream = append(b.stream, 0)
		b.remainCount = 8
	}

	latestIndex := len(b.stream) - 1

	b.stream[latestIndex] |= data >> (8 - b.remainCount)
	b.stream = append(b.stream, 0)
	latestIndex++
	b.stream[latestIndex] = data << b.remainCount
}

//WriteBits :
func (b *BStream) WriteBits(data uint64, count int) {
	data <<= uint(64 - count)

	//handle write byte if count over 8
	for count >= 8 {
		byt := byte(data >> (64 - 8))
		b.writeByte(byt)

		data <<= 8
		count -= 8
	}

	//handle write bit
	for count > 0 {
		bi := data >> (64 - 1)
		b.writeBit(bi == 1)

		data <<= 1
		count -= 1
	}
}
