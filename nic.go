// Title: Script to print network information

// Author: Chu Dumebi Umezinne

package main

// Import required packages
import (
	"fmt"
	"net"
)

// Utilizes the builtin net package to gather network information and prints to screen
func main() {
	nics, _ := net.Interfaces() // lists all interfaces
	for _, nic := range nics {  // iterates through interfaces to print and get IPs
		networkAddress, _ := nic.Addrs() // gets the IPv4/6 address for each interface
		fmt.Println("Index:", nic.Index)
		fmt.Println("MTU:", nic.MTU)
		fmt.Println("Name:", nic.Name)
		fmt.Println("MAC:", nic.HardwareAddr)
		fmt.Println("IPv4/6:", networkAddress, "\n")
	}
}
