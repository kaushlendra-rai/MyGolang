package main

import (
	"fmt"
)

func getClosure() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	fun := getClosure()

	fmt.Println(fun())
	fmt.Println(fun())

	newFun := getClosure()
	fmt.Println(newFun())
}
