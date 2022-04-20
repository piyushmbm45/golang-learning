package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files of gplang")
	content := "this needs to go in file - myname"

	file, err := os.Create("./myfile.txt")

	errFunc(err)
	length, err := io.WriteString(file, content)
	errFunc(err)
	fmt.Println("length is", length)
	readFile("./myfile.txt")
	defer file.Close()
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	errFunc(err)
	fmt.Println("text data in file \n", string(dataByte))
}

func errFunc(err error) {
	if err != nil {
		panic(err)
	}
}
