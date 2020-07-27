package main

import (
	"fmt"
	"os"

	"./hoge"
	_ "./hoge/kana"
	_ "./hoge/kata"
)

func main() {
	err := hoge.Print("kata")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
