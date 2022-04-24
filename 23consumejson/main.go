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
	fmt.Println("Hello from json consume go lang")
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React",
		"Price": 200,
		"website": "dev.go",
		"tags": ["web-dev","react"]
	}
	`)
	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Not Valid")
	}

	// some cases where you just want to add data to key value

	var myJsonweb map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myJsonweb)
	fmt.Printf("%#v\n", myJsonweb)
	for k, v := range myJsonweb {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}
}
