package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	domain := "google.com"
	if len(os.Args) >= 2 {
		domain = os.Args[1]
	}

	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf(domain+" IN A %s\n", ip.String())
	}
}
