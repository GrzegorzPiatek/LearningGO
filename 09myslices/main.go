package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to lessons about slices")

	var fruitList = []string{}

	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Strawberry")
	fruitList = append(fruitList, "Mango", "Strawberry")
	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fmt.Println(append(fruitList[1:3]))

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 456
	highScore[3] = 567
	// highScore[4] = 58867
	highScore = append(highScore, 555, 565, 999)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"java", "javascript", "go", "ruby"}

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
