// Dup2は標準入力から2回以上現れる行を出現回数と一緒に表示します。
// 標準入力から読み込むか、名前が指定されたファイルの一覧から読み込みます。
// https://github.com/YoshikiShibata/gpl/blob/master/ch01/ex04/dup2.go
// を参考にしてしまった。
// 後に出てくるであろう、データ構造や便利な関数を使ってなんとかするのか、
// それとも、現状の限られた道具でパズルすべきなのか困る。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, arg, counts, filenames)
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t(%s)\n", n, line, filenames[line])
		}
	}
}

func countLines(
	f *os.File,
	filename string,
	counts map[string]int,
	filenames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if filenames[line] == "" {
			filenames[line] = filename
		} else {
			if filenames[line][0:len(filename)] != filename {
				filenames[line] = filename + ", " + filenames[line]
			}
		}
	}
	// 注意: input.Err()からのエラーの可能性を無視している
}
