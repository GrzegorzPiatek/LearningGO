package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of ptr is", ptr)

	myNumber := 23
	var ptrToMyNumber = &myNumber
	fmt.Println("Value of ptr is", ptrToMyNumber)
	fmt.Println("Value of ptr is", *ptrToMyNumber)

	*ptrToMyNumber = *ptrToMyNumber * 2

	fmt.Println("Value of ptr is", myNumber)
}
