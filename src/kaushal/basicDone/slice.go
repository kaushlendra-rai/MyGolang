package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)

	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)

	s = append(s, "d")

	fmt.Println(s)

	str := []string{"1", "2", "3"}
	fmt.Println(str)

	subSlice := str[2:]
	fmt.Println("subSlice :", subSlice)

	twoD := make([][]int, 3)

	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, i+1)

		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	// =====================================
	// Append one slice to another

	s = append(s, str...)
	fmt.Println(s)
}
