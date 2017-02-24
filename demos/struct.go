package main

import "fmt"

type Employee struct {
	id   int
	age  int
	name string
}

func printEmpoloyee(e Employee) {
//func printEmpoloyee(e *Employee) { // struct pointer as argument
	fmt.Println("id: ", e.id)
	fmt.Println("age: ", e.age)
	fmt.Println("name: ", e.name)
}

func main() {
	var tom Employee
	var jerry Employee

	// assign to a field
	tom.id = 2
	tom.age = 21
	tom.name = "Tom"

	// written using a struct literal
	jerry = Employee{1, 20, "Jerry"}

	fmt.Println(jerry)
	printEmpoloyee(tom)
	//printEmpoloyee(&tom)

	var emp *Employee = &tom
	fmt.Println(emp)
}
