package main

import "fmt"

func main() {
	k := 2        // declare and assigned
	var i int     // declare but not assigned, 0 by default
	var j int = 1 // declare and assigned

	// declare more
	x, y := 320, 240
	var a, b, c int = 3, 4, 5

	// boolean
	var isprime = true // infer type by value

	fmt.Println(i, j, k)
	fmt.Println(a, b, c)
	fmt.Println(x, y)
	fmt.Println(isprime)

	var m uint32 = 0x12345678
	var n uint64 = 0x1234567812345678
	var r float32 = 5.0
	var pi float64 = 3.141592653525
	vec1 := 3 + 4i // complex

	var name string = "Brian W. Kernighan"

	fmt.Println(m, n)
	fmt.Println(r, pi)
	fmt.Println(vec1)
	fmt.Println(name)

	iptr := &i // pointer

	printType := func(i interface{}) {
		switch t := i.(type) { // type switch
		default:
			fmt.Printf("%v:\t%T\n", i, t)
		}
	}
	printType(iptr)
	printType(vec1)
	printType(3.14)
	printType(1234)
	printType(true)
	printType("abc")
}
