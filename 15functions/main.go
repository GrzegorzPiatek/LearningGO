package main

import "fmt"

func main() {
	fmt.Println("Welcom to functions in GO!")
	greeter()

	value := adder(3, 5)
	uslessVarTocheckUslessFunction := slicesAppender([]string{"one", "two"}, "three")
	fmt.Println("Result is: ", value)
	fmt.Println("UslessVarTocheckUslessFunction is: ", uslessVarTocheckUslessFunction)
	val1, val2 := proAdder(1, 2, 3, 4)
	fmt.Println("Pro adder of 1, 2, 3, 4 is: ", val1)
	fmt.Println("second from pro adder", val2)
}

func adder(a int, b int) int {
	return a + b
}

func slicesAppender(slice []string, someStr string) []string {
	return append(slice, someStr)
}

func greeter() {
	fmt.Println("Namstey from golang")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi pro result function"
}
