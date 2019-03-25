package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayHello(s string) {
	// Indicate to wait group that the thread is done with computation.
	// Using defer to ensure that the 'Done' gets executed even if the below code gets a panic/exception// Indicate to wait group that the thread is done with computation.
	wg.Done()
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%s\n", s)
	}
}

func main() {
	go sayHello("Sonu")
	wg.Add(1)
	go sayHello("Rai")
	wg.Add(1)

	// Wait for 'Sonu' & 'Rai' to complete.
	wg.Wait()

	go sayHello("Done")
	wg.Add(1)
	wg.Wait()
}
