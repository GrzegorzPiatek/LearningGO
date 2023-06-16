package main

import "fmt"

const loginToken string = "asdfoihbg"

func main() {
	var username string = "gregory"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.01230123012301230123
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var otherOne string
	fmt.Println(otherOne)
	fmt.Printf("Variable is of type: %T \n", otherOne)

	var website = "somewebpage.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	numberOfusers := 30000
	fmt.Println(numberOfusers)
	fmt.Printf("Variable is of type: %T \n", numberOfusers)

	fmt.Println(loginToken)
	fmt.Printf("Variable is of type: %T \n", loginToken)
}
