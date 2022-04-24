package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request in go lang")
	fmt.Println("Please start backend server")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostForm()
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

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/api/post"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename": "golang",
		"price": 0
	}
`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("json data", string(content))
}

func PerformPostForm() {
	const myUrl = "http://localhost:8000/api/post-form/"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Piyush")
	data.Add("lastname", "Jain")
	data.Add("email", "piyush@go.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("body", string(content))
}
