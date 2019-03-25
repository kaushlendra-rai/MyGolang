package main

import (
	"fmt"
	"rai/string"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(string.Reverse("Hello Kaushal"))

	a, b := 0, 0
	i := 5
	a, b = b+i, i

	fmt.Println(a)
	fmt.Println(b)
}
