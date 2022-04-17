package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in golang ")

	loginCount := 10

	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "Heavy User"
	} else {
		result = "something else"
	}
	fmt.Println(result)

	// even odd
	num := 9
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("Odd")
	}

	// other syntax
	if num := 3; num < 10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Not less than 10")
	}
}
