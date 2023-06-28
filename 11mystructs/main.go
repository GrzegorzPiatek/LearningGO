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
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
