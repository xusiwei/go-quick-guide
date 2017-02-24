package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func values() (int, int, int) {
	return 1, 2, 3
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {

	fmt.Println(plus(1, 2))

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	a, b, c := values()
	fmt.Println(a, b, c)

	// call anonymous directly
	func() {
		fmt.Println("Anonymous function called")
	}()

	// func as argument
	// /*
	foreach := func(s []int, f func(int, int)) {
		for i, v := range s {
			f(i, v)
		}
	} // */
	foreach([]int{2, 3, 5, 7}, func(i, v int) { fmt.Println(i, v) })

	// func F() {} // error, can not nest named functions

	// closure
	// /*
	makeCounter := func() (func() uint) {
		count := uint(0)
		return func() uint {
			count += 1
			return count
		}
	} // */
	c1 := makeCounter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	c2 := makeCounter()
	fmt.Println(c2())
	fmt.Println(c2())
}

// /*
func foreach(s []int, f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

func makeCounter() func() uint {
	count := uint(0)
	return func() uint {
		count += 1
		return count
	}
}
// */

