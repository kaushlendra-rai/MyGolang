package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [5]int{1, 2, 3, 5, 7}
	fmt.Println(primes)

	var slice []int = primes[1:4]
	fmt.Println(slice)

	primes[3] = 111

	fmt.Println(slice)

	slice = append(slice, 11, 13, 17, 19)

	fmt.Println("*********************** ", slice)
	// ===============================================================

	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	board[0][0] = "X"
	board[1][0] = "Y"
	board[2][2] = "X"

	fmt.Println(board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//========================================================================

	for i, v := range primes {
		fmt.Printf("Index %d = %d,  ", i, v)
	}

	// You can skip the index or value by assigning to _
	for _, v := range primes {
		fmt.Printf(" Value: %d,  ", v)
	}

}

func Pic(dx, dy int) [][]uint8 {
	var pic = make([]([]uint8), dy)

	for y, _ := range pic {
		pic[y] = make([]uint8, dx)

		for x, _ := range pic[y] {
			pic[y][x] = uint8(x & y)
		}
	}

	return pic
}
