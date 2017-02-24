package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	arrayArgs(a)
	fmt.Println("after arrayArgs:", a)

	s := a[:]
	sliceArgs(s)
	fmt.Println("after sliceArgs:", s)
}

// pass array copy to a function
func arrayArgs(a [5]int) {
	a[1] = 88
	fmt.Println("in arrayArgs, a:", a)
}

func sliceArgs(s []int) {
	s[1] = -2
	fmt.Println("in sliceArgs, s:", s)
}

