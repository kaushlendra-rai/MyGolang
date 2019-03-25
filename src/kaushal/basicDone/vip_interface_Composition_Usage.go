package main

import (
	"fmt"
)

type Speaker interface {
	speak() string
}

type Listener interface {
	listen(word string)
}

type Social interface {
	Speaker
	Listener
}

type Animal struct {
	Color   string
	IsMamal bool
}

type Cat struct {
	Animal
	LikeMilk bool
}

type Dog struct {
	Animal
	CanSwim bool
}

func (c *Cat) speak() string {
	return "Meow from Cat"
}

func (c *Cat) listen(word string) {
	fmt.Println(word, " LikeMilk: ", c.LikeMilk)
}

func (d *Dog) speak() string {
	return "Bark from Dog"
}

func (d *Dog) listen(word string) {
	fmt.Println(word, " CanSwim : ", d.CanSwim)
}

func main() {
	social := []Social{
		&Cat{
			Animal: Animal{
				Color: "White", IsMamal: true,
			},
			LikeMilk: true},
		&Dog{
			Animal: Animal{
				Color: "Brown", IsMamal: true,
			},
			CanSwim: true},
	}

	for _, v := range social {
		v.listen(v.speak())
	}
}
