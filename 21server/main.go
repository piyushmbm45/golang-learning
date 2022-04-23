package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request in go lang")
	fmt.Println("Please start backend server")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/api/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code is ", response.StatusCode)
	fmt.Println("Content Length is", response.ContentLength)
	var responseStrings strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseStrings.Write(content)
	fmt.Println("Byte Count", byteCount)
	fmt.Println("body ", responseStrings.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
}
