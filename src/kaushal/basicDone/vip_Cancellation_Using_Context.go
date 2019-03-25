package main

import (
	"context"
	"fmt"
	"time"
)

type data string

func main() {
	duration := 150 * time.Millisecond

	// Create a context which is both manually cancellable and will signal
	// after a certain period of time
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan data, 1)

	// Ask teh goroutine to do some work for us
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Done the work")

		ch <- data("Done Work")
	}()

	select {
	case out := <-ch:
		fmt.Print("Output Received: ", out)
	case <-ctx.Done():
		// When the above line is executed, the actual timer beings and send a signal after timeout defined above.
		// In the above case, if we set the timeout to say '10', we will get a timeout as the goroutine sleeps for 50 ms to do work.
		fmt.Print("Timeout ")
	}
}
