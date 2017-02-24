package main

import "fmt"

type Employee struct {
	Id   int
	Age  int
	Name string
}

func printEmpoloyee(e Employee) {
	fmt.Println("Id: ", e.Id)
	fmt.Println("Age: ", e.Age)
	fmt.Println("Name: ", e.Name)
}

func duplicateEmpoloyee(e *Employee) *Employee { // struct pointer as argument
	return &Employee{e.Id, e.Age, e.Name}
}

func main() {
	var tom Employee
	var jerry Employee

	// assign to a field
	tom.Id = 2
	tom.Age = 40
	tom.Name = "Tom"

	// written using a struct literal
	jerry = Employee{1, 30, "Jerry"}

	fmt.Println(jerry)
	printEmpoloyee(tom)

	var emp *Employee = duplicateEmpoloyee(&tom)
	tom.Age = 20
	fmt.Println(tom)
	fmt.Println(emp)
}
