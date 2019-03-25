package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Handled panic in cleanup", r)
	}
}
func sayHello(s string) {
	// Indicate to wait group that the thread is done with computation.
	// Using defer to ensure that the 'Done' gets executed even if the below code gets a panic/exception
	defer cleanup()

	for i := 0; i < 3; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%s\n", s)

		if i == 2 {
			panic("In trouble i=2")
		}
	}
}

func main() {
	go sayHello("Sonu")
	wg.Add(1)
	go sayHello("Rai")
	wg.Add(1)

	// Wait for 'Sonu' & 'Rai' to complete.
	wg.Wait()
}
