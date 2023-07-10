package main

import "fmt"

func main() {
	fmt.Println("Welcom in GOlang structs")
	gregory := User{"Grzegorz", "grzegorz@go.dev", true, 23}

	fmt.Println(gregory)

	fmt.Printf("gregory details are: %+v \n", gregory)
	fmt.Printf("gregory name: %v \n", gregory.Name)
	fmt.Printf("gregory email: %v \n", gregory.Email)
	fmt.Printf("gregory age: %v \n", gregory.Age)
	fmt.Printf("gregory status: %v \n", gregory.Status)
	gregory.GetStatus()
	gregory.NewEmail()
	fmt.Printf("gregory email: %v \n", gregory.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
