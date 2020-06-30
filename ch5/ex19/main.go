package main

import "fmt"

func double(x int) (res int) {
	defer func() {
		p := recover() // ないとpanic
		fmt.Printf("defer %v\n", p)
		res = x + x
	}()
	panic("double")
}

func main() {
	fmt.Println(double(4))
}
