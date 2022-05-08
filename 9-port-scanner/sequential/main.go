package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {
		var err error
		var conn net.Conn
		conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			fmt.Println("Port", i, "closed")
			continue
		}

		fmt.Println("Port", i, "open")
		err = conn.Close()
		if err != nil {
			fmt.Println("Error closing connection on port", i)
		}
	}
}
