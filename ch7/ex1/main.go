package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// ByteCounter は文字数?を数えます
type ByteCounter int

// Write は
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// WordCounter は単語数を数えます
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

// LineCounter は行数を数えます
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	fmt.Println()
	var wc WordCounter
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Fprint(&wc, "言語　日本語")
	fmt.Println(wc)

	fmt.Println()
	hoge := `
	hoge piyo
	
	bar %s 2
	bazz`
	var lc LineCounter
	fmt.Fprint(&lc, hoge)
	fmt.Fprintf(&lc, hoge, name)
	fmt.Println(lc)
}
