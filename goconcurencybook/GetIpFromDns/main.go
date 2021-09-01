package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// ips, _ := net.LookupIP("ec2-3-213-82-59.compute-1.amazonaws.com")
	ipsStr := []string{}
	ips, _ := net.LookupIP("lightspeed.tp-eng-usva-2.nite.tradestation.io")
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			ipsStr = append(ipsStr, ipv4.String())
			fmt.Println("IPv4: ", ipv4)
		}
	}
	fmt.Println(ipsStr)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s", hostname)
}
