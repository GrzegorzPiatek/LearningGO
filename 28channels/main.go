package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang!")

	myChannel := make(chan int, 2)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	// recv only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		if isChannelOpen {
			fmt.Println(val)
		}

		wg.Done()
	}(myChannel, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 7

		close(ch)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
