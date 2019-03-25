package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

func (v Vertex) area() int {
	return v.x * v.y
}

func area2(v *Vertex) int {
	return v.x * v.y
}

// Pointer receivers
func (v *Vertex) scale(factor int) {
	v.x = v.x * factor
	v.y = v.y * factor
}

func scale2(v *Vertex, factor int) {
	v.x = v.x * factor
	v.y = v.y * factor
}

func main() {
	v := Vertex{10, 20}

	fmt.Println(v.area())

	v1 := &v
	v1.scale(2)
	fmt.Println(v1.area())

	p := &v
	fmt.Println("Test2")
	scale2(p, 2)
	fmt.Println(area2(p))
}
