package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang!")

	content := "This needs to go in a file.\nA test"

	file, err := os.Create("./myGoFile.txt")
	panicError(err)

	length, err := io.WriteString(file, content)
	panicError(err)

	fmt.Println("len of file is: ", length)
	defer file.Close()

	fmt.Print(readFile("./myGoFile.txt"))
}

func readFile(filename string) string {
	databyte, err := ioutil.ReadFile(filename)
	panicError(err)
	return string(databyte)
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
