package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	defer udpConn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := udpConn.ReadFromUDP(buf)
		fmt.Println("Received: ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("ERROR:", err)
		}
	}

}
