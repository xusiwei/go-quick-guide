package main

import "fmt"

func main() {
	count := 5
	for count > 0 {
		fmt.Println(count)
		count = count - 1
	}

	for i := 0; i < 5; i++ {
		// if i%2 { // non-bool i % 2 (type int) used as if condition
		if i%2 == 0 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
	}

	count = 1
	for { // always loop
		fmt.Println("loop")
		if count == 3 {
			break
		}
		count = count + 1
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
