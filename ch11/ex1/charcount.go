// https://github.com/YoshikiShibata/gpl/blob/master/ch11/ex01/main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts, utflen, invalid, err := charcount(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func charcount(r io.Reader) (map[rune]int, []int, int, error) {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(r)
	for {
		for {
			r, n, err := in.ReadRune() // rune, nbytes, errorを返す
			if err == io.EOF {
				return counts, utflen[:], invalid, nil
			}
			if err != nil {
				return nil, nil, 0, err
			}
			if r == unicode.ReplacementChar && n == 1 {
				invalid++
				continue
			}
			counts[r]++
			utflen[n]++
		}
	}
}
