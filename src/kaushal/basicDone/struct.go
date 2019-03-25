package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{10, 20})

	v := Vertex{2, 3}
	v.X = 5

	fmt.Println(v.X)

	p := &v

	fmt.Println(p.X)

	p.Y = 1e5
	fmt.Println(v)
}
