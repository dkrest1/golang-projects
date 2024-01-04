package main

import (
	"fmt"
	"net"
)

func main() {
	nameserver, _ := net.LookupNS("google.com")

	for  _, ns := range nameserver {
		fmt.Println(ns)
	}
}