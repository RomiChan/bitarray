package bitarray

import (
	"math/bits"
	"strings"
)

type SignificantBit bool

// 位序
const (
	LSB SignificantBit = false
	MSB SignificantBit = true
)

type BitArray struct {
	d     []byte
	p     int
	ismsb SignificantBit
}

func NewBitArray(buf []byte, sb SignificantBit) *BitArray {
	return &BitArray{d: buf, p: 0, ismsb: sb}
}

// Len bit数
func (ba *BitArray) Len() int {
	return ba.p
}

func (ba *BitArray) Append(bit bool) {
	index := ba.p / 8       // 看应该对那个字节动手
	if index == len(ba.d) { // 不够长 把里面的append一下
		ba.d = append(ba.d, 0) // 填0
	}
	offset := 0
	if ba.ismsb {
		offset = 7 - ba.p%8
	} else {
		offset = ba.p % 8
	}
	if bit {
		ba.d[index] |= 1 << offset
	}
	ba.p++
}

func (ba *BitArray) Extend(bits []bool) {
	for _, v := range bits {
		ba.Append(v)
	}
}
func (ba *BitArray) AppendByte(data byte) {
	var i uint
	if ba.ismsb {
		for i = 128; i > 0; i >>= 1 {
			ba.Append(uint(data)&i > 0)
		}
	} else {
		for i = 1; i <= 128; i <<= 1 {
			ba.Append(uint(data)&i > 0)
		}
	}
}

func (ba *BitArray) ExtendBytes(data []byte) {
	for _, v := range data {
		ba.AppendByte(v)
	}
}

func (ba *BitArray) GetBit(index int) bool {
	var offset uint
	if ba.ismsb {
		offset = 7 - uint(index%8)
	} else {
		offset = uint(index % 8)
	}
	return uint(ba.d[index/8])&(1<<offset) > 0
}

func (ba *BitArray) SetBit(index int) {
	mod := index / 8
	offset := 0
	if ba.ismsb {
		offset = 7 - index%8
	} else {
		offset = index % 8
	}
	ba.d[mod] |= 1 << offset
}

func (ba *BitArray) ClearBit(index int) {
	mod := index / 8
	offset := 0
	if ba.ismsb {
		offset = 7 - index%8
	} else {
		offset = index % 8
	}
	ba.d[mod] &^= 1 << offset
}

func (ba *BitArray) Count() (ret int) {
	for _, b := range ba.d {
		ret += bits.OnesCount8(b)
	}
	return
}

func (ba *BitArray) Add(array *BitArray) (nba *BitArray) {
	nba = &BitArray{d: ba.d, p: ba.p, ismsb: ba.ismsb}
	for i := 0; i < array.Len(); i++ {
		nba.Append(array.GetBit(i))
	}
	return
}

func (ba *BitArray) Bytes() []byte {
	mod := ba.p % 8
	l := ba.p / 8
	if mod > 0 {
		l++
	}
	return ba.d[:l]
}

func (ba *BitArray) String() string {
	var b strings.Builder
	for i := 0; i < ba.p; i++ {
		if ba.GetBit(i) {
			b.WriteString("1")
		} else {
			b.WriteString("0")
		}
	}
	return b.String()
}
