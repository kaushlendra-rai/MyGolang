package main

import "fmt"

type Mooing interface {
	Moo() string
}

type Grazing interface {
	EatGrass()
}

type Cow struct{}

func (c *Cow) Moo() string {
	return "moo"
}

func (c *Cow) EatGrass() {
	fmt.Println("Eating grass")
}

// Composite Interface
type Animal interface {
	Mooing
	Grazing
}

func Milk(cow Animal) {
	cow.Moo()
	cow.EatGrass()
}

func main() {
	// To access a method over Interface, we always pass address to the method
	Milk(&Cow{})
}
