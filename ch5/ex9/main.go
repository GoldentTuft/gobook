package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	x := "foo"
	return strings.ReplaceAll(s, "$"+x, f(x))
}

func hoge(string) string {
	return "hoge"
}

func main() {
	fmt.Println(expand("あいう$fooえお$fooかき", hoge))
}
