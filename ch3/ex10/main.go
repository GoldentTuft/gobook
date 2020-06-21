package main

import (
	"bytes"
	"fmt"
)

func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma1(s[:n-3]) + "," + s[n-3:]
}

func reverseString(s string) string {
	res := ""
	for _, v := range s {
		res = string(v) + res
	}
	return res
}

func comma2(s string) string {
	var buf bytes.Buffer
	var slen = len(s)

	for i := 1; i <= slen; i++ {
		buf.WriteByte(s[slen-i])
		if i%3 == 0 && i != slen {
			buf.WriteByte(',')
		}
	}
	return reverseString(buf.String())
}

func main() {
	fmt.Println(comma1("1234567"))
	fmt.Println(comma2("1234567"))
	fmt.Println(comma1("123"))
	fmt.Println(comma2("123"))
	fmt.Println(comma1(""))
	fmt.Println(comma2(""))
}
