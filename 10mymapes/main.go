package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning map in golang")

	languages := make(map[string]string)
	languages["RB"] = "Ruby"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	fmt.Println("All languages in map :", languages)
	fmt.Println("Single value : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("All languages in map :", languages)

	// loops are interesting in golang
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
