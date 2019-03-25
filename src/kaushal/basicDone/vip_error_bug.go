package main

import (
	"log"
)

type customError struct{}

// Error implements the 'error' interface
func (ce *customError) Error() string {
	return "Finding the BUG"
}

// We should never return the concrete 'error' type because it would never be nil.
// The below concrete value holds a pointer to the type 'customError' and a value 'nil'.
// Hence when this nil is checked in the main method against 'nil', it would fail as not nil (due to reference to the type of error this object holds
//func fail() ([]byte, *customError) {
func fail() ([]byte, error) {
	return nil, nil
}

func main() {
	var err error

	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail")
	}

	log.Println("No Problem")
}
