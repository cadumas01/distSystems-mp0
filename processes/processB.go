package processes

// server

import (
	"bufio"
	"fmt"
	"net"
)

var bufSize = 2048

const Confirmation = "ACK"

// Starts listener are returns listener variable
func StartServer(address string) {
	ln, err := net.Listen("tcp", address)

	if err != nil {
		panic("Error listening")
	}

	fmt.Println("server started")

	acceptClient(ln)
}

// Waits for client to connect and recieves message
func acceptClient(ln net.Listener) {

	// Waits for client to connect
	conn, err := ln.Accept()

	if err != nil {
		panic("error accepting")
	}

	handleConnection(conn)

}

func handleConnection(conn net.Conn) {

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

	return
}
