package kana

import (
	"fmt"

	"github.com/GoldentTuft/gobook/ch10/study/import/hoge"
)

func print() {
	fmt.Print("ほげ")
}
func init() {
	hoge.RegisterFormat("kana", print)
}
