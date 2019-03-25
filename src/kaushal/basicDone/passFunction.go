package main

import (
	"fmt"
	"math"
)

func compute(fun func(x, y float64) float64, a, b float64) float64 {
	return fun(a, b)
}

var counter int = 0

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		counter = counter + 1
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i, j, counter := 0, 1, 0

	fib := func() int {
		if counter < 2 {
			counter = counter + 1
			return counter - 1
		}

		num := i + j
		i, j = j, num

		return num
	}

	return fib
}

func main() {
	hypot := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}

	fmt.Println(hypot(4, 3))

	fmt.Println(compute(hypot, 5.0, 12.0))

	myFunc := adder()

	for i := 0; i < 5; i++ {
		fmt.Println(myFunc(i), counter)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
