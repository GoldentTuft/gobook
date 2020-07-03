package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

// CounterWriter は
type CounterWriter struct {
	writer io.Writer
	count  int64
}

func (w *CounterWriter) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	w.count += int64(n)
	return n, err
}

// CountingWriter は
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw CounterWriter
	cw.writer = w
	return &cw, &cw.count
}

func main() {
	writer, counter := CountingWriter(new(WordCounter))
	fmt.Fprintf(writer, "hello, %s", "hoge")
	fmt.Fprint(writer, "言語　日本語")
	fmt.Println(*counter)

}
