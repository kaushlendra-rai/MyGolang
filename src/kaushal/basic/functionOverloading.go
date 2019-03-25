package main

import (
	"fmt"
)

func main() {
	doTestMe(21)
	doTestMe(21, "Kaushal")
}

func doTestMe(age int) {
	fmt.Println("Age: %d", age)
}

func doTestMe(age int, name string) {
	fmt.Println("Age: %d, name: %s", age, name)
}
