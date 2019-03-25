package mainTemp

import (
	"fmt"
)

func doMath(a int, b int) int {
	return (a + b)
}

func doMath2(a, b int) int {
	return (a + b)
}

// Multiple return values
func swap(a, b int) (int, int) {
	return b, a
}

// Named returns
func namedReturns(x, y int) (a, b int) {
	a = x * 2
	b = y * 2

	return
}

func main() {
	sum := doMath(2, 3)
	fmt.Println(sum)

	sum = doMath2(3, 3)
	fmt.Println(sum)

	a, b := swap(6, 8)
	fmt.Println(a, b)

	c, d := namedReturns(11, 16)
	fmt.Println(c, d)
}
