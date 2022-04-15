package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome from time package")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-06-2006 15:04:05 Monday"))

	createdDate := time.Date(2021, time.July, 2, 1, 45, 3, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
