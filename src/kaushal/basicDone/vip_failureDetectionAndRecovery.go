package main

import (
	"fmt"
	log "kaushal/basicDone/logger"
	//"log"
	"os"
	"os/signal"
	"time"
)

// device allows us to mock the device to which we intend to write teh logs to
type device struct {
	problem bool
}

func (d *device) Write(p []byte) (n int, err error) {

	// Simulate a blocking issue if a problem exists
	for d.problem {
		time.Sleep(time.Second)
	}

	fmt.Println(string(p))
	return len(p), nil
}

func main() {

	// Number of Goroutines that will be writing logs
	const grt = 10

	// Create a Logger with enough buffer capacity
	// For each gouroutine that will be logging in parallel.
	var d device
	//l := log.New(&d, "prefix", 0)
	l := log.New(&d, 100)

	// Generate goroutines each writing to our custom logs
	for i := 0; i < grt; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking. Capture
	// interrupt signals to toggle device issues. Use <Ctrl> Z to kill the program
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan

		// I appriciate we have a data race here with the Write method.
		// Lets keep it simple for now to show the mechanics
		d.problem = !d.problem
	}
}
