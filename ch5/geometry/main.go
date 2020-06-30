package main

//package geometry

import (
	"fmt"
	"math"
)

// Point は座標
type Point struct{ X, Y float64 }

// Path は点を直線で結びつける道のりです。
type Path []Point

// Distance は
// 昔ながらの関数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance は
// 同じだが、Point型のメソッドとして
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance はpathに沿って進んだ距離を返します。
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 5
	fmt.Println(p.Distance(q))  // 5

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
