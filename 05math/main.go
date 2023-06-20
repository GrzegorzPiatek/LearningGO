package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcom in simple GOlang math operands")

	a, b := 3, 5

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
	fmt.Printf("%v / %v = %v\n", a, b, a/b)
	fmt.Printf("%v %% %v = %v\n", a, b, a%b)
	fmt.Printf("%v ^ %v = %v\n", a, b, math.Pow(float64(a), float64(b)))
	fmt.Printf("sqrt(%v,2) = %v\n", a, math.Pow(float64(a), 0.5))
}
