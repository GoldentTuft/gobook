package popcount

import (
	"log"
	"os"
)

// pc[i] はiのポピュレーションカウントです。
var pc [256]byte

var cwd string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		// fmt.Printf("pc[%d]:%v, i/2:%v, i&1:%b\n", i, pc[i], i/2, byte(i&1))
	}
}

// PopCount はxのポピュレーションカウント(1が設定されているビット数)を返します。
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
