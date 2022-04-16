package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Guvava"
	fruitList[3] = "Mango"
	fmt.Println("The fruit list is ", fruitList)
	fmt.Println("The length of fruit list", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Loki"}
	fmt.Println("Veg List ", vegList)
}
