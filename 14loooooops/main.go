package main

import "fmt"

func main() {
	fmt.Println("Try some loops")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	for i := range days {
		fmt.Println(i+1, days[i])
	}

	for i, day := range days {
		fmt.Printf("day of index %v is %v \n", i, day)
	}

	rougueValue := 0

	for rougueValue < 10 {
		rougueValue++

		if rougueValue == 2 {
			continue
		}
		fmt.Println("Value is: ", rougueValue)

		if rougueValue == 4 {
			goto lco
		}
		fmt.Println("Value is: ", rougueValue)
	}

lco:
	fmt.Println("Jumping to this")
}
