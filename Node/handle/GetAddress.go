package handle

import (
	"fmt"
	"net"
)

func GetAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		for _, addr := range addrs {
			ipAddr, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				fmt.Println("Error parsing IP:", err)
				continue
			}

			// Skip loopback and non-IPv4 addresses
			if ipAddr.IsLoopback() || ipAddr.To4() == nil {
				continue
			}

			return string(ipAddr)
			// fmt.Fprintf(w, "Interface: %s, IP: %s\n", iface.Name, ipAddr)
		}
	}
	return ""
}
