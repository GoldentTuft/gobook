package kana

import (
	"fmt"

	// "github.com/GoldentTuft/gobook/tree/master/ch10/study/import/hoge"
	"../../hoge"
)

func print() {
	fmt.Print("ほげ")
}
func init() {
	hoge.RegisterFormat("kana", print)
}
