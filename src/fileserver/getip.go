package main

import (
	"fmt"
	"net"
)

func getInternal() {
	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	os.Stderr.WriteString("Oops:" + err.Error())
	// 	os.Exit(1)
	// }
	// for _, a := range addrs {
	// 	if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			os.Stdout.WriteString("\t" + ipnet.IP.String() + "\n")
	// 		}
	// 	}
	// }

	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	for i, inter := range interfaces {
		fmt.Printf("\t(%d)%s :\n", i, inter.Name)
		addrs, _ := inter.Addrs()
		for indx, addr := range addrs {
			fmt.Printf("\tAddr[%d] = %s\n", indx, addr)
		}
	}
}
