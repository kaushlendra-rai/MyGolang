package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
	dept string
}

func (u User) displayName() {
	fmt.Println("displayName: ", u.name)
}

func (u *User) displayRealName() {
	fmt.Println("displayRealName: ", u.name)
}

func main() {
	u := User{"Sonu", 35, "BIPRD"}

	// NOTE: A function using 'value' semantics for object gets a copy of object upfront right at that point.
	// That copy remains uneffected by  any change in the underneath object.
	dn := u.displayName
	u.name = "Kaushal"
	dn()

	// A function reference using 'Pointer' semantics for object's method reference acts on actual original object.
	// Any change to the underlying object gets reflected in the method.
	drn := u.displayRealName
	u.name = "Changed Sonu"
	drn()
}
