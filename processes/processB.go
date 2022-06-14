package processes

// server

import (
	"bufio"
	"fmt"
	"net"
)

var bufSize = 2048

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
func AcceptClient(ln net.Listener) {
	for {
		// Waits for client to connect
		conn, err := ln.Accept()

		if err != nil {
			panic("error accepting")
		}

		go handleConnection(conn)

		//		defer ln.Close()
	}
}

func handleConnection(conn net.Conn) {

	fmt.Println("Client connected from " + conn.RemoteAddr().String())
	//	defer ln.Close()

	buf := make([]byte, bufSize)

	// fmt.Println("before the loop")

	// fmt.Println("top of loop")
	_, err := bufio.NewReader(conn).Read(buf)
	//fmt.Println("n  =" + strconv.Itoa(n))

	if err != nil {
		panic(err)
	}

	message := string(buf)
	fmt.Print("\nServer recieved the Message: \n" + message)

	// Send confiramtion message
	conn.Write([]byte("MESSAGE RECEIVED"))

	//conn.Write([]byte("Just read a message:" + message))

	//conn.Write([]byte("this is a server message"))

	//fmt.Println("Post error check")

}
