package main

import (
	"crypto/sha256"
	"fmt"

	"../../ch2/popcount"
)

func countDiffBit(x, y [32]byte) int {
	res := 0
	for i := 0; i < len(x) && i < len(y); i++ {
		res += popcount.PopCount(uint64(x[i] ^ y[i]))
	}
	return res
}

func fillBit() [32]byte {
	var res [32]byte
	for i := range res {
		res[i] = 255
	}
	return res
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var zero [32]byte
	fill := fillBit()
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(countDiffBit(c1, c2))
	fmt.Println(countDiffBit(zero, fill))
	fmt.Println(countDiffBit(zero, zero))
	fmt.Println(countDiffBit(fill, fill))
}
