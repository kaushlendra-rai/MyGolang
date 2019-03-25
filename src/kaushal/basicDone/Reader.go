package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "Have a nice day"

	r := strings.NewReader(s)
	b := make([]byte, 8)

	for {
		n, e := r.Read(b)

		fmt.Printf("%q", b[:n])

		if e == io.EOF {
			break
		}
	}
}
