package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	errorCode int
	msg       string
}

func main() {

	v, e := testError(1)

	fmt.Println(v, e)

	v, e = testError(2)
	fmt.Println(v, e)

	// ================== Custom Errors
	v, e = testError2(1)

	fmt.Println(v, e)

	v, e = testError2(2)
	fmt.Println(v, e)

	//------------------ Casting Interface to Struct
	v, e = testError2(1)
	fmt.Println(v, e)
	m, ok := e.(MyError)

	fmt.Println("After casting: ", m, ok)
}

func (e MyError) Error() string {
	return fmt.Sprintf("%d - %s", e.errorCode, e.msg)
}

func testError2(num int) (int, error) {

	if num == 1 {
		return -1, MyError{1001, "********Sample Default Error"}
	}

	return 2 * num, nil
}

func testError(num int) (int, error) {

	if num == 1 {
		return -1, errors.New("Sample Default Error")
	}

	return 2 * num, nil
}
