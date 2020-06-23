package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("H", "A")
	fmt.Println(graph)
	fmt.Println(hasEdge("A", "H"))
	fmt.Println(hasEdge("A", "Q"))
	fmt.Println(hasEdge("X", "y"))
	fmt.Println(hasEdge("A", "C"))
}
