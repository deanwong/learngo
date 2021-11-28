package main

import (
	"fmt"

	"github.com/deanwong/learngo/geometry"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

type coloredTriangle struct {
	triangle
	color string
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 3
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
	t.doubleSize()
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())

	colorT := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", colorT.size)
	fmt.Println("Perimeter", colorT.triangle.perimeter())

	gt := geometry.Triangle{}
	gt.SetSize(3)
	fmt.Println("Perimeter", gt.Perimeter())
}
