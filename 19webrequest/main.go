package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("Welcom to web request GoLang lesson!")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Response type: %T\n", response)

	datebytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(datebytes)

	fmt.Println(content)
}
