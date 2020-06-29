package main

import (
	"fmt"
	"strings"
)

func join(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}

func main() {
	fmt.Println(join(",", "hoge", "piyo"))
	fmt.Println(join(","))
	fmt.Println(join(",", "hoge"))
}
