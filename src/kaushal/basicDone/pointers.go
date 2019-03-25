package main

import (
	"fmt"
)

func main() {
	i, j := 10, 15

	p := &i

	fmt.Println("Pointer : ", *p)

	*p = 21

	fmt.Println("Updated Pointer i : ", i)

	p = &j
	*p = *p / 5

	fmt.Println("Updated Pointer j : ", j)
}
