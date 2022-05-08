package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incommingClients   = make(chan Client)
	leavingClients     = make(chan Client)
	allClientsMessages = make(chan string)
)

var (
	host = flag.String("host", "localhost", "host")
	port = flag.Int("port", 3090, "port")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	newClientMessages := make(chan string)

	// Send messages from newClientMessages to the connection
	go MessageWrite(conn, newClientMessages)

	clientName := conn.RemoteAddr().String()
	newClientMessages <- fmt.Sprintf("Welcome to the server, %s", clientName)
	allClientsMessages <- fmt.Sprintf("%s has joined", clientName)

	incommingClients <- newClientMessages

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		allClientsMessages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- newClientMessages
	allClientsMessages <- fmt.Sprintf("%s has left", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprint(conn, message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-allClientsMessages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incommingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	var err error
	var listener net.Listener
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	go Broadcast()
	for {
		var conn net.Conn
		conn, err = listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go HandleConnection(conn)
	}
}
