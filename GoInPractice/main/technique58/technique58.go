package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hostname:", name)
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, a := range addrs {
		fmt.Println(a)
	}
}
