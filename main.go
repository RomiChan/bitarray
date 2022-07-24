package main

import (
	"fmt"
	"github.com/RomiChan/bitarray/bitarray"
)

func main() {
	array := bitarray.NewBitArray(bitarray.LITTLE)
	array.ExtendBytes([]byte{0, 1, 255, 255})
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
