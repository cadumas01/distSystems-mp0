package processes

//client

import (
	"bufio"
	"bytes"
	"fmt"
	"mp0/messages"
	"net"
)

func SendMessage(address string) {
	//Connect to port
	conn, err := net.Dial("tcp", address)

	if err != nil {
		SendMessage(address)
		panic(err)
	}

	fmt.Println("Client Successfully connected from " + conn.LocalAddr().String())

	// Dealing with Message construction
	m := messages.ConstructMessage()

	// Send Message
	fmt.Println("Sending a message...")

	// Write Message over tcp channel
	_, err = conn.Write([]byte(m.String()))

	if err != nil {
		panic("Error writing message")
	}

	// Wait for confirmation
	waitForACK(conn)
}

func waitForACK(conn net.Conn) {

	buf := make([]byte, 64)

	// fmt.Println("Inside waitForACK")
	for s := ""; s != Confirmation; {
		s = read(conn, buf)
	}

	fmt.Println("Confirmation received, client exiting...")

	return

}

func read(conn net.Conn, buf []byte) string {
	_, err := bufio.NewReader(conn).Read(buf)

	if err != nil {
		fmt.Println("Error reading")
	}

	// trim unused bytes in buffer and convert to string
	s := string(bytes.Trim(buf, "\x00"))

	return s
}
