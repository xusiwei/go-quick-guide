package main

import "fmt"

func main() {
	// create slice from slice-literal
	t := []string {"s", "l", "i", "c", "e"}
	fmt.Println("t:", t)

	a := [5]string {"a", "r", "r", "a", "y"}
	fmt.Println("\na:", a)
	sa := a[2:4]
	fmt.Println("sa:", sa)
	sa[0] = "R"
	sa[1] = "A"
//	a[4:] = []string{"YY", "z"} // don't support
	fmt.Println("sa:", sa)
	fmt.Println("a:", a)   // also update
	fmt.Println("len(sa):", len(sa))
	fmt.Println("cap(sa):", cap(sa))

	// create a slice using `make`
	s := make([]string, 3)
	fmt.Println("\ninit, s:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("after set, s:", s)
	fmt.Println("s[2]: ", s[2])
//	fmt.Println("s[3]: ", s[3]) // panic: runtime error: index out of range

	// slice size not fixed, unlike array
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("after append, s:", s)

	c := make([]string, len(s))
	copy(c, s) // built-in function copy
	fmt.Println("\ncopy c <-- s:", c)


	// slice operatoration [:]
	fmt.Println("\ns[2:5]:", s[2:5])
	fmt.Println("s[2:]:", s[2:])
	fmt.Println("s[:5]:", s[:5])
	fmt.Println("s[:]:", s[:])
//	fmt.Println("s[:-1]:", s[:-1]) // don't support


	// slice can be multi-dimensional:
	//  the length of inner slices can vary,
	//  unlike multi-dimensional array.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("\n2d slice:", twoD)
}

// refers: http://blog.golang.org/go-slices-usage-and-internals
