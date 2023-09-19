package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error getting network interfaces: %s\n", err)
		return
	}

	fmt.Println("Device Name\tIP Address")

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("Error getting address for interface %s: %s\n", iface.Name, err)
			continue
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if v.IP.To4() != nil { // we only want IPv4 addresses
					fmt.Printf("%s\t%s\n", iface.Name, v.IP)
				}
			}
		}
	}
}
