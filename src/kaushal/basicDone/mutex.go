package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	data map[string]int
	mux  sync.Mutex
}

func (counter *SafeCounter) increment(key string) int {

	// If we do not use Locks, we get error ::
	// "fatal error: concurrent map writes"
	counter.mux.Lock()
	counter.data[key]++
	counter.mux.Unlock()

	return counter.data[key]
}

func (counter *SafeCounter) value(key string) int {

	// If I do not use the Locking while Read, I get below Error::
	// fatal error: concurrent map read and map write

	// It MUST be noted that when using 'defer' for Unlock API, it worked sometime and failed randomly as well
	// Hence, using explicit  Unlock prior to returning.

	counter.mux.Lock()
	val := counter.data[key]
	defer counter.mux.Unlock()

	return val
}

func main() {
	c := SafeCounter{data: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.increment("abc")
	}

	time.Sleep(time.Second)

	fmt.Println(c.value("abc"))

	fmt.Println("Done")
}
