package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 3090, "port")
	host = flag.String("host", "localhost", "host")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))

	if err != nil {
		log.Fatal(conn)
	}

	done := make(chan struct{})
	go func() {
		SendMessage(os.Stdout, conn)
		done <- struct{}{}
	}()

	SendMessage(conn, os.Stdin)
}

func SendMessage(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)

	if err != nil {
		log.Fatal(err)
	}
}
