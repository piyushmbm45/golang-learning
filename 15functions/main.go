package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println(result)

	resultAll, str := proAdder(1, 2, 3, 4, 5)
	fmt.Println(resultAll)
	fmt.Println(str)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		// total = total + val
		total += val
	}
	return total, "Hi resultAll"
}

func greeter() {
	fmt.Println("Namaste from Function")
}
