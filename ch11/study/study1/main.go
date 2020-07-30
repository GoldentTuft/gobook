package main

import "fmt"

func main() {
	for n := 0; n < 10; n++ {
		fmt.Printf("%d: \n", n)
		fmt.Printf("\t")
		for i := 0; i < (n+1)/2; i++ {
			fmt.Print(i)
		}
		fmt.Printf("\n\t")
		for i := 0; i < (n+1)/2; i++ {
			fmt.Print(n - 1 - i)
		}
		fmt.Println("")
	}

	for n := 0; n < 10; n++ {
		fmt.Printf("%d/2=%d, (%d+1)/2=%d\n", n, n/2, n, (n+1)/2)
	}
}
