package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to oue home page")
	fmt.Println("Rate our page between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	// reading till one line
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	numberRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number rating", numberRating+1)
	}
}
