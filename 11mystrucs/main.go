package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")
	// no inheritance in golang ; NO super, or parent
	piyush := User{"Piyush", "p@gmail.com", true, 16}
	fmt.Println(piyush)
	fmt.Printf("Piyush Details Are: %+v\n", piyush)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
