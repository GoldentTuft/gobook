package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	ts := topoSort(prereqs)
	for i, course := range ts {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	fmt.Println(isTopologicalOrdered(ts))
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// fmt.Printf("***%s\n", item)
				visitAll(m[item])
				order = append(order, item)
				// fmt.Println(item)
				// for _, v := range order {
				// 	fmt.Printf("  %s\n", v)
				// }

			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func isTopologicalOrdered(ts []string) bool {
	nodes := make(map[string]int)

	for i, course := range ts {
		nodes[course] = i
	}

	for course, i := range nodes {
		for _, prereq := range prereqs[course] {
			// fmt.Printf("%s:%d, %s:%d\n", course, i, prereq, nodes[prereq])
			// course:調べるコース
			// i:調べるコースの順位
			// nodes[prereq]:調べるコースの依存先の順位
			// 依存先の順位のほうが小さいはずなので、条件式はfalseになりreturn falseされない
			if i < nodes[prereq] {
				return false
			}
		}
	}
	return true
}
