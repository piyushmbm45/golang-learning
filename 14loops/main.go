package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println(index)
	// 	fmt.Println(day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto try
		}
		if rougueValue == 5 {
			rougueValue++
			// break
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

try:
	fmt.Println("this is from goto statement")
}
