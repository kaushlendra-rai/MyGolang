package main

import (
	"fmt"
)

func main() {

	s := "*"

	for i, v := range "Ram" {
		fmt.Println(i, v)

		fmt.Printf("%v : %c \n", i, v)

		s += fmt.Sprintf("%c", v)
	}

	fmt.Println(s)
}
