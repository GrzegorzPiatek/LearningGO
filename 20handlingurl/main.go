package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asd123"

func main() {
	fmt.Println("Welcom to handling URLs in go lang!")

	fmt.Println(myurl)

	// parsing
	parsedUrl, _ := url.Parse(myurl)
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("RawQuery:", parsedUrl.RawQuery)

	qparams := parsedUrl.Query()

	fmt.Printf("Type of query params: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(partsOfUrl)
	fmt.Println(anotherUrl)

}
