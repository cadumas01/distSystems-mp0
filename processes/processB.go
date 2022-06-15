package processes

// server

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var bufSize = 2048

const Confirmation = "ACK"

// Starts listener are returns listener variable
func StartServer(address string) (ln net.Listener) {
	ln, err := net.Listen("tcp", address)

	if err != nil {
		panic("Error listening")
	}

	fmt.Println("server started")

	return ln
}

// Waits for client to connect and recieves message
func AcceptClient(ln net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Waits for client to connect
		conn, err := ln.Accept()

		if err != nil {
			panic("error accepting")
		}

		handleConnection(conn, wg)
	}
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, bufSize)
	_, err := bufio.NewReader(conn).Read(buf)

	if err != nil {
		panic(err)
	}

	message := string(buf)
	fmt.Print("\nServer recieved the Message: \n" + message)

	// Send confiramtion message
	conn.Write([]byte(Confirmation))
	fmt.Println("Confirmation sent, server exiting...")
}
