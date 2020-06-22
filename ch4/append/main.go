package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 拡大する余地がある。スライスを拡張する。
		z = x[:zlen]
	} else {
		// 十分な領域がない。新たな配列を割り当てる。
		// 計算量を線形に均すために倍に拡大する。
		zcap := zlen
		// このifがわからない。len(x) != 0これと同じ?
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 組み込み関数
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
