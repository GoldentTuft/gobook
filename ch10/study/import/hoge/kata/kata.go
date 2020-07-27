package kata

import (
	"fmt"

	// "github.com/GoldentTuft/gobook/tree/master/ch10/study/import/hoge"
	"../../hoge"
)

func print() {
	fmt.Print("ホゲ")
}
func init() {
	hoge.RegisterFormat("kata", print)
}
