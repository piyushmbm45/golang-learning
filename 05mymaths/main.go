package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in golang")
	// var numberOne int = 2
	// var numberTwo float64 = 4

	// fmt.Println("The Sum is: ", numberOne+numberTwo)

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("random num: ", myRandomNum)
}
