package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	length, breath float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func (r Rect) area() float64 {
	return r.breath * r.length
}

func (r Rect) perimeter() float64 {
	return 2 * (r.breath + r.length)
}

func printMe(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {

	c := Circle{2}

	printMe(c)

	r := Rect{3, 4}

	printMe(r)
}
