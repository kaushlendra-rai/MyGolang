package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

type admin struct {
	//person user // NOT embedded user
	user  // Embedded user
	level string
}

type Notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("Sending email to user: %s at email: %s \n", u.name, u.email)
}

// @@@@@@@@@@@@@@@@@@@@@@  UNCOMMENT this method and test again @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//func (u *admin) notify() {
//	fmt.Printf("********* Sending email to admin: %s at email: %s \n", u.name, u.email)
//}

func main() {

	//ad := admin{
	//	person:  user{name: "Sonu", email: "sonu@gmail.com"},
	//	level: "super",
	//}

	ad := admin{
		user:  user{name: "Sonu", email: "sonu@gmail.com"},
		level: "super",
	}

	ad.user.notify()
	ad.notify()

	// If we comment-out the method notify for admin: 'func (u *admin) notify() {', the notify() for *user would be used by &admin call due to inner-type promotion (embedding of struct)
	sendNotification(&ad)

	u := user{"Kaushal", "kaushal@sas.com"}

	// We need to send address of the object as the receiver to the notify() method of Notifier method is a pointer.
	// If we remove the pointer from the 'notify' method, we can remove the pass by address '&'  from below all.
	sendNotification(&u)
}

func sendNotification(u Notifier) {
	u.notify()
}

func testUser(u user) {
	fmt.Println("Test User", u.name, u.email)
}
