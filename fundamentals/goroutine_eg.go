package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipaddr IPAddr) String() string {
	var s string = ""
	for _, octet := range ipaddr {
		s = fmt.Sprintf("%v%v%v", s, ".", octet)
	}
	return fmt.Sprintf("%v", s[1:])
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
