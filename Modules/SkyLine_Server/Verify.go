package SkyLine_Network

import (
	"fmt"
	"net"
)

var PortList = []string{
	"5598",
	"5597",
	"5596",
	"5595",
	"5594",
	"5593",
	"5592",
	"5591",
}

func VerifyConnection() string {
	for _, k := range PortList {
		fmt.Println("Checking...")
		host := net.JoinHostPort("127.0.0.1", k)
		fmt.Println("Checking host address ", host)
		conn, err := net.Dial("tcp", host)
		if err != nil {
			return k
		} else {
			conn.Close()
		}
	}
	return "empty"
}
