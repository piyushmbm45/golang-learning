package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1, Open Now")
	case 2:
		fmt.Println("Move by 2 Spot")
	case 3:
		fmt.Println("Move by 3 Spot")
		fallthrough
	case 4:
		fmt.Println("Move by 4 Spot")
		fallthrough
	case 5:
		fmt.Println("Move by 5 Spot")
	case 6:
		fmt.Println("Move by 6 Spot")
	default:
		fmt.Println("Ohoo Some Error")
	}
}
