package main

import "fmt"

type entry struct {
	value int
	ready chan struct{}
}

func main() {
	var e, f entry

	e = entry{ready: make(chan struct{})}

	fmt.Println(e, f)
}
