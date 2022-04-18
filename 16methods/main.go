package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")

	// no inheritance in golang ; NO super, or parent
	piyush := User{"Piyush", "p@gmail.com", true, 16}
	fmt.Println(piyush)
	fmt.Printf("Piyush Details Are: %+v\n", piyush)
	piyush.GetStatus()
	piyush.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.com"
	fmt.Println("Email of this user is ", u.Email)
}
