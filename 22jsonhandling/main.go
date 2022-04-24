package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string `json:"coursename"`
	Price     int
	Plateform string   `json:"website"`
	Password  string   `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json handling in go lang")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"React", 200, "dev.go", "sdfbsdbb", []string{"web-dev", "react"}},
		{"Java", 201, "dev.go", "sdfbsdbb", []string{"web-dev", "jsva"}},
		{"JS", 202, "dev.go", "sdfbsdbb", []string{"web-dev", "Js"}},
		{"Go", 203, "dev.go", "sdfbsdbb", nil},
	}

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
