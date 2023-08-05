package main

import "fmt"

func main() {
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Counting: ")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
