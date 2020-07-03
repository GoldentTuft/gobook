package main

import (
	"fmt"
	"image/color"
	"math"
)

// Point は
type Point struct{ X, Y float64 }

// ColoredPoint は
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Distance は
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy はスケールする
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Add は
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

// Sub は
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

// Path は
type Path []Point

// TranslateBy は
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	var cp ColoredPoint
	// Point指定しなくても、Xにアクセスできる
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	// メソッドに関しても、Point指定しなくてもPointのDistanceを呼び出せる
	// is-aではないので、q.Pointと渡さなければならない
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Point.Distance(q.Point))
	fmt.Printf("%T\n", p.Distance)
	fmt.Printf("%T\n", Point.Distance)

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	perim.TranslateBy(Point{1, 1}, false)
	fmt.Println(perim)
}
