package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Darwin")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Possibly Windows", runtime.GOOS)
	}

	switch day := time.Now().Weekday(); day {
	case time.Friday, time.Saturday, time.Sunday:
		fmt.Println("Happy weekend")
	default:
		fmt.Println("Work on weekdays")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
