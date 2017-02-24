package main

import "fmt"
import "math"

type Shape interface {
	Area() float64
}

func Area(s Shape) float64 {
	return s.Area()
}

// main
func main() {
	r := Rect{Point{0, 0}, Point{3, 4}}
	c := Circle{Point{2, 2}, 1}

	// Rect satisfies Shape interface, so we can pass r to Area
	// but no type dependency between Rect and Shape
	fmt.Println(Area(r))

	// *Shape satisfies Shape interface
	fmt.Println(Area(&c))
}

type Point struct {
	X float64
	Y float64
}

type Rect struct {
	LeftTop  Point
	RightBot Point
}

func (r Rect) Area() float64 {
	width := r.RightBot.X - r.LeftTop.X
	height := r.RightBot.Y - r.LeftTop.Y
	if width < 0 || height < 0 {
		panic(fmt.Sprintf("invalid Rect %v", r))
	}
	return width * height
}

type Circle struct {
	Center Point
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}


