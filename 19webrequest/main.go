package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("Web request learning")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() // callers responsibilty to close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println("data from web page is", content)
}
