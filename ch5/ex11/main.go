package main

import (
	"fmt"
	"log"
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

var circledPrereqs = map[string][]string{
	"algorithms":           {"data structures"},
	"calculus":             {"linear algebra"},
	"linear algebra":       {"calculus"},        // circle
	"intro to programming": {"data structures"}, // another circle

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
	for i := 0; i < 10000; i++ {
		ts, err := topoSort(prereqs)
		if err != nil {
			log.Fatal(err)
		}
		ok := isTopologicalOrdered(ts)
		if ok == false {
			log.Fatal(ok)
		}
	}

	for i := 0; i < 10000; i++ {
		_, err := topoSort(circledPrereqs)
		if err == nil {
			log.Fatal(err)
		}
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var result []string
	seen := make(map[string]bool)
	route := make(map[string]bool)
	var visitAll func(course string) error

	visitAll = func(course string) error {
		if seen[course] == false {
			for _, c := range m[course] {
				if route[c] {
					return fmt.Errorf("Circle detected: %s => %s",
						c, course)
				}
				route[c] = true
				err := visitAll(c)
				route[c] = false
				if err != nil {
					return err
				}
			}
			seen[course] = true
			result = append(result, course)
		}
		return nil
	}

	for course := range m {
		route[course] = true
		err := visitAll(course)
		if err != nil {
			return nil, err
		}
		route = make(map[string]bool)
	}

	return result, nil
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
