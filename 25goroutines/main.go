package main

import (
	"fmt"
)

func main() {
	go greeter("hello")
	greeter("welcome")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
