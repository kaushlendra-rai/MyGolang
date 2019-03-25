package main

import (
	"fmt"
)

type IPAddr [4]byte

// A Stringer is a type that can describe itself as a string.
// The String() method could be implemented in case of stringer interface
func (ip IPAddr) String() string {
	ret := fmt.Sprintf("%v", ip[0])

	for i := 1; i < len(ip); i = i + 1 {
		ret += fmt.Sprintf(".%v", ip[i])
	}

	return ret
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
