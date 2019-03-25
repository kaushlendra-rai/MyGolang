package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doFoo(c chan int, factor int) {

	// Ensure that Task completion by Thread is indicated to the WaitGroup
	defer wg.Done()
	c <- factor * 5

}

func main() {

	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doFoo(c, i)
	}

	// We need to wait untill all the Threads complete their processing.
	wg.Wait()

	// Close channel as an indication to Range that no more elements would be added to channel anymore.
	// Only after closing channel, we can use range over it.
	close(c)

	for cVal := range c {
		fmt.Println(cVal)
	}
}
