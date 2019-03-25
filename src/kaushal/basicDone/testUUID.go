package main

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

func main() {

	id := uuid.New().String()

	fmt.Println("Hello: ", id)
}
