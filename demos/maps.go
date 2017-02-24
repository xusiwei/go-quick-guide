package main

import "fmt"

func main() {
	// create from map literal
	n := map[string]int {"foo": 1, "bar": 2}
	fmt.Println("n:", n)

	// use make create map
	m := make(map[string]int)
	// m := map[string]int{}

	// insert new key, value
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println("m:", m)

	// lookup key
	v1 := m["one"]
	fmt.Println("m[\"one\"]:", v1)

	fmt.Println("len(m):", len(m))

	// remove key
	delete(m, "one")
	fmt.Println("after delete, m:", m)

	// iterate
	for k, v := range m {
		fmt.Printf("%s => %d\n", k, v)
	}

	// second return value indicates:
	//  the key was present in map or not
	_, found := m["one"]
	fmt.Println("found:", found)
}
