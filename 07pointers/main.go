package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// initial value of pointer
	var one *int
	fmt.Println(one)

	myNumber := 2
	var ptr = &myNumber
	// memory address
	fmt.Println("Value of actual pointer ", ptr)
	// memory value
	fmt.Println("Value is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is", myNumber)
}
