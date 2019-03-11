package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return fmt.Sprintf("%q", string(ip[:len(ip)]))
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

// TODO: MUST be
// Misalnya, IPAddr{1, 2, 3, 4} mengeluarkan "1.2.3.4".
