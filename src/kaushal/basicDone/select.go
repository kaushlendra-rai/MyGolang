package main

import (
	"fmt"
	"time"
)

func main() {
	basic()

	medium()
}

func medium() {
	work := make(chan int)
	done := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-work)
		}

		done <- 0
	}()

	fibonacci(work, done)
}

func fibonacci(work, done chan int) {
	x, y := 0, 1

	for {
		select {
		case work <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}

func basic() {
	a := make(chan string)
	b := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		a <- "sonu"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		b <- "rai"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-a:
			fmt.Println(msg1)

		case msg2 := <-b:
			fmt.Println(msg2)

		}
	}
}
