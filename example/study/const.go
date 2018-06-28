package main

import "fmt"

func main() {
	const (
		IP0             = "192.168.1.0"
		IP1      string = "192.168.1.1"
		IP2, IP3        = "192.168.1.2", "192.168.1.3"
		C0              = iota
	)
	fmt.Println(IP0, IP1, IP2, IP3, C0)
}
