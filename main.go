package main

import (
	"fmt"

	"github.com/RomiChan/bitarray/bitarray"
)

func main() {
	array := bitarray.NewBitArray(nil, bitarray.LSB)
	array.ExtendBytes([]byte{0, 0, 123, 234})
	fmt.Println(array.String())
	array.Append(true)
	array.Append(false)
	fmt.Println(array.String())
	for i := 0; i < array.Len(); i++ {
		fmt.Println(array.GetBit(i))
	}
	fmt.Println(array.Count())
	fmt.Println(array.Bytes())
	array = bitarray.NewBitArray(nil, bitarray.MSB)
	array.ExtendBytes([]byte{0, 0, 123, 234})
	fmt.Println(array.String())
	array.Append(true)
	array.Append(false)
	fmt.Println(array.String())
	for i := 0; i < array.Len(); i++ {
		fmt.Println(array.GetBit(i))
	}
	fmt.Println(array.Count())
	fmt.Println(array.Bytes())
}
