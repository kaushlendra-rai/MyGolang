package main

import (
	"fmt"
	"time"
)

type MyError struct {
	errorCode   int
	errorString string
	time        time.Time
}

func (myError MyError) Error() string {
	return fmt.Sprintf("ErrorCode=%v  , Error:%s  occured at:%v", myError.errorCode, myError.errorString, myError.time)
}

func testMe(val int) (int, error) {
	if val > 0 {
		return val * 2, nil
	} else {
		return 0, MyError{111, "Negative number", time.Now()}
	}
}

func main() {
	fmt.Println(testMe(2))
	fmt.Println(testMe(-2))

	negNum := ErrNegativeSqrt(-2.12)
	fmt.Printf("%s", negNum.Error())
}

type ErrNegativeSqrt float32

func (negFloat ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(negFloat))
}
