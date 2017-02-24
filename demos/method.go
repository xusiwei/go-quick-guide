package main

import "fmt"

type Rect struct {
	Width, Height int
}

// area method has a *receiver type* of *Rect
func (r *Rect) area() int {
	return r.Width * r.Height
}

// methods can be defined for either pointer or value receiver types.
func (r Rect) perim() int {
	return 2*(r.Width + r.Height)
}

func (r *Rect) setWidth(w int) {
	r.Width = w
}

func main() {
	r := Rect{10, 5}
	fmt.Println(r)
	fmt.Println("area:", r.area())
	fmt.Println("prim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("prim:", rp.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	//  you may want to use pointer receiver type to aviod coping on method calls
	//  or to allow the method to mutate the receiving struct.

	f := r.setWidth // method value
	f(5)
	fmt.Println(r)
}
