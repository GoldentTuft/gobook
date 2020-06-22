package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"

	"../../ch2/popcount"
)

var useSHA384 = flag.Bool("sha384", false, "use sha384")
var useSHA512 = flag.Bool("sha512", false, "use sha512")

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
	flag.Parse()

	if len(os.Args) <= 1 {
		os.Exit(1)
	}
	x := []byte(os.Args[1])
	if *useSHA384 {
		fmt.Printf("SHA384:%x\n", sha512.Sum384(x))
	} else if *useSHA512 {
		fmt.Printf("SHA512:%x\n", sha512.Sum512(x))
	} else {
		fmt.Printf("SHA256:%x\n", sha256.Sum256(x))
	}
}
