package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")
var maxPort = flag.Int("maxPort", 100, "max port to scan")

func main() {
	flag.Parse()
	var wg sync.WaitGroup

	for port := 0; port < *maxPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			var err error
			var conn net.Conn

			conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}

			fmt.Println("Port", port, "open")

			err = conn.Close()
			if err != nil {
				fmt.Println("Error closing connection on port", port)
			}
		}(port)
	}
	wg.Wait()
}
