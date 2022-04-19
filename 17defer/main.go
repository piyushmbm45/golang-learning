package main

import "fmt"

func main() {
	// follow lifo
	defer fmt.Println("this is from defer")
	defer fmt.Println("this is from defer-1")
	defer fmt.Println("this is from defer-2")
	fmt.Println("Welcome to defer")
	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}
}
