package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("Paper sent to employee")
	}()

	// Below line returns another channel of type timeout
	tc := time.After(100 * time.Millisecond)

	select {
	case p := <-ch:
		fmt.Println("Value from standard channel: ", p)
	case t := <-tc:
		// Timeout can be easily observed if we use 5000 instead of 500 in the above sleep random time
		fmt.Println("Timeout occured: ", t)
	}

	time.Sleep(time.Second)
	fmt.Println("___________________----------------------------------------------________________________________")
}
