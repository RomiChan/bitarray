package bitarray

import (
	"strings"
)

const (
	BIG = iota
	LITTLE
)

type BitArray struct {
	data   []byte
	bits   int
	endian int
}

func setbit(bit byte, offset int) byte {
	return bit | (1 << offset)
}

func getbit(bit byte, offset int) bool {
	if (bit & (1 << offset)) > 0 {
		return true
	}
	return false
}

func clearbit(bit byte, offset int) byte {
	//return ^(^bit | (1 << offset))
	return bit &^ byte(1<<offset)
}

func NewBitArray(endian int) *BitArray {
	return &BitArray{data: make([]byte, 0, 10), bits: 0, endian: endian}
}

// bit数
func (self *BitArray) Len() int {
	return self.bits
}

func (self *BitArray) Append(bit bool) {
	index := self.Len() / 8       // 看应该对那个字节动手
	if index > len(self.data)-1 { // 不够长 把里面的append一下
		self.data = append(self.data, 0) // 填0
	}
	if bit { // fixme
		self.data[index] = setbit(self.data[index], self.Len()%8)
	} else {
		self.data[index] = clearbit(self.data[index], self.Len()%8)
	}
	self.bits++
}

func (self *BitArray) Extend(bits []bool) {
	for _, v := range bits {
		self.Append(v)
	}
}
func (self *BitArray) AppendByte(data byte) {
	for i := 0; i < 8; i++ {
		self.Append(getbit(data, i))
	}
}

func (self *BitArray) ExtendBytes(data []byte) {
	for _, v := range data {
		self.AppendByte(v)
	}
}

// todo index check
// index是bit的索引
func (self *BitArray) GetBit(index int) bool {
	return getbit(self.data[index/8], index%8)
}

// index是bit的索引
func (self *BitArray) SetBit(index int) {
	mod := index / 8
	self.data[mod] = setbit(self.data[mod], index%8)
}

// index是bit的索引
func (self *BitArray) ClearBit(index int) {
	mod := index / 8
	self.data[mod] = clearbit(self.data[mod], index%8)
}

// 含1的个数
func (self *BitArray) Count() int {
	var ret int = 0
	for i := 0; i < self.Len(); i++ {
		if self.GetBit(i) {
			ret++
		}
	}
	return ret
}

// return new array
func (self *BitArray) Add(array *BitArray) *BitArray {
	ret := NewBitArray(self.endian)
	for i := 0; i < self.Len(); i++ {
		ret.Append(self.GetBit(i))
	}
	for i := 0; i < array.Len(); i++ {
		ret.Append(array.GetBit(i))
	}
	return ret
}

func (self *BitArray) Bytes() []byte {
	mod := self.Len() % 8
	var l int
	if mod > 0 {
		l = self.Len()/8 + 1
	} else {
		l = self.Len() / 8
	}
	return self.data[:l]
}

func (self *BitArray) String() string {
	var b strings.Builder
	for i := 0; i < self.Len(); i++ {
		if self.GetBit(i) {
			b.WriteString("1")
		} else {
			b.WriteString("0")
		}
	}
	return b.String()
}
