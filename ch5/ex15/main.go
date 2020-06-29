package main

import (
	"fmt"
)

func max1(values ...int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("No arguments")
	}
	res := values[0]
	for _, v := range values {
		if res < v {
			res = v
		}
	}
	return res, nil
}

func min1(values ...int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("No arguments")
	}
	res := values[0]
	for _, v := range values {
		if res > v {
			res = v
		}
	}
	return res, nil
}

func max2(a int, values ...int) int {
	res := a
	for _, v := range values {
		if res < v {
			res = v
		}
	}
	return res
}

func min2(a int, values ...int) int {
	res := a
	for _, v := range values {
		if res > v {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(max1(5, -5, 10, 3))
	fmt.Println(max1())
	fmt.Println(min1(5, -5, 10, 3))
	fmt.Println(min1())
	fmt.Println()
	fmt.Println(max2(5, -5, 10, 3))
	fmt.Println(max2(0))
	fmt.Println(min2(5, -5, 10, 3))
	fmt.Println(min2(0))
}
