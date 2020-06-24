package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort はvalues内の値をその中でソートします。
func Sort(values []int) {
	var root *tree
	// 木を作ってく
	for _, v := range values {
		root = add(root, v)
	}
	// 木からスライスに戻す
	appendValues(values[:0], root)
}

// appendValues はtの要素をvaluesの正しい順序に追加し、
// 結果のスライスを返します。
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// return &tree{value: value}と同じ
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	sample := []int{5, 8, 3, 2}
	fmt.Println(sample)
	Sort(sample)
	fmt.Println(sample)
}
