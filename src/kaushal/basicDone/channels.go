package main

import (
	"fmt"
)

func multiply(c chan int, b int) {
	c <- b * 5
}

func main() {
	basic()

	basic1()

	medium()
}

func medium() {
	work := make(chan int)
	//done := make(chan int)

	go func() {
		for {
			num, more := <-work

			if more {
				fmt.Println(num)
			} else {
				fmt.Println("Done")
			}
		}
	}()

	for i := 0; i < 5; i++ {
		work <- i
	}

	//done <- 0

	close(work)

	fmt.Println("Done with Channels")
}

func basic1() {
	channel := make(chan int, 10)

	go fibonaci(cap(channel), channel)

	for v := range channel {
		fmt.Println(v)
	}
}

func fibonaci(n int, channel chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		channel <- x
		x, y = y, x+y
	}

	close(channel)
}

func basic() {
	channel := make(chan int)

	go multiply(channel, 3)
	go multiply(channel, 5)

	a, b := <-channel, <-channel

	fmt.Println(a, b)
}
