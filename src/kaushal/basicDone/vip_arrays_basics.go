package main

import (
	"fmt"
)

func main() {
	names := [5]string{"sonu", "rai", "kaushal", "mishti", "kshitij"}

	// If you have to change the values of the array, set it by index.
	// If you have to access the value only for read, access it by value
	// One should never mix the value and index flow
	for i, v := range names {
		names[1] = "Neha"
		v = "Manju"
		if i == 1 {
			fmt.Println(names[i])
			fmt.Println(v)
		}
	}

	fmt.Println(names[1])
}
