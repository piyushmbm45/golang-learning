package main

import "fmt"

const Jwttoken string = "dsjbfjdsbj" // Public - as first letter of name is capital

func main() {
	var username string = "Piyush"
	fmt.Println(username)
	fmt.Printf("Varilable of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varilable of type: %T \n", isLoggedIn)

	var smallVal float64 = 3490989.984584785
	fmt.Println(smallVal)
	fmt.Printf("Varilable of type: %T \n", smallVal)

	// default values and alias
	var myInt int
	fmt.Println(myInt)
	fmt.Printf("Varilable of type: %T \n", myInt)

	var myStr string
	fmt.Println(myStr)
	fmt.Printf("Varilable of type: %T \n", myStr)

	// implicit type - don't provide type
	var myProject = "Authentication"
	fmt.Println(myProject)
	fmt.Printf("Varilable of type: %T \n", myProject)

	// no var style - valrus operator
	numOfUsers := 5000
	fmt.Println(numOfUsers)

	// public varible
	fmt.Println(Jwttoken)
	fmt.Printf("Varilable of type: %T \n", Jwttoken)

}
