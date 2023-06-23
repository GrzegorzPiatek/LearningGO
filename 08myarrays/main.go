package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[3] = "Potato"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushrooms"}
	fmt.Println("Fruit list is: ", vegList)
	fmt.Println("Fruit list is: ", len(vegList))
	fmt.Printf("Type of fruitList is %T\n", vegList)

}
