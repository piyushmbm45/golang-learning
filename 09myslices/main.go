package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("Type of fruit list %T\n", fruitList)

	fruitList = append(fruitList, "Banana", "Grapes")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 342
	highScore[2] = 423
	highScore[3] = 3421
	// highScore[4] = 421

	highScore = append(highScore, 323, 333, 221)
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
}
