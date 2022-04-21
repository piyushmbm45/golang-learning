package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courseName=reactjs&user=piyush&role=admin"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	// parsing

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qParamas := result.Query()
	fmt.Println(qParamas["courseName"])

	for _, val := range qParamas {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/learn",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
