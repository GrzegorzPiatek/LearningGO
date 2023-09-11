package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // usually pointers
var mut sync.Mutex    //usually pointers

func main() {

	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://lco.dev",
		"https://github.com",
		"https://asasdasdaqwdasd.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

//func greeter(s string) {
//	for i := 0; i < 6; i++ {
//		time.Sleep(2 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Oops, error for %s\n", endpoint)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d code for website for %s \n", result.StatusCode, endpoint)
	}
}
