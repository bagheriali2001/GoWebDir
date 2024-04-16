package helpers

import (
	"fmt"
	"net"
	"strings"
)

// TODO: colorize the output
func DisplayAvailableAddresses(serverPort string) {
	fmt.Println("Server is running on the following addresses:")
	for _, address := range getNetwork() {
		fmt.Printf("\thttp://%s:%s\n", address, serverPort)
	}
}

func getNetwork() []string {
	var addressList []string

	// Get a list of all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Iterate over the interfaces and print their names and addresses
	for _, i := range interfaces {
		if i.Flags&net.FlagUp == 0 {
			continue
		}

		// if it doesn't have a hardware address ignore it
		// TEMPORARILY COMMENTED, Localhost has no hardware address, so it will be ignored
		// if i.HardwareAddr.String() == "" {
		// 	continue
		// }
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println("Error in Reading Address: ", err)
			continue
		}

		for _, addr := range addrs {
			if getIPVersion(addr.(*net.IPNet).IP) == "IPv6" {
				continue
			}
			addressList = append(addressList, removeSubnetMask(addr.String()))
		}
	}

	return addressList
}

// A function to determine the version of the IP address
func getIPVersion(ip net.IP) string {
	if ip.To4() != nil {
		return "IPv4"
	}
	return "IPv6"
}

// A function to remove the subnet mask from the IP address
func removeSubnetMask(ip string) string {
	index := strings.Index(ip, "/")
	if index == -1 {
		return ip
	} else {
		// return the ip address without the subnet mask
		return ip[:index]
	}
}
