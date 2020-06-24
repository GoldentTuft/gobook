package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordNum := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	total := 0

	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordNum[word]++
		total++

	}

	for k, v := range wordNum {
		fmt.Printf("\t%s:\t%.3f%%\n", k, float64(v)/float64(total)*100.0)
	}

	fmt.Println(wordNum)

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
}
