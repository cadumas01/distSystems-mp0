package processes

//client

import (
	"bufio"
	"bytes"
	"fmt"
	"mp0/messages"
	"net"
	"strings"
)

// argument:
func SendMessage(address string) {
	//Connect to port
	conn, err := net.Dial("tcp", address)

	if err != nil {

		fmt.Println("SendMessage error")
		SendMessage(address)
		panic(err)
	}

	fmt.Println("Successfully connected")

	// load message and send it on channel
	m := messages.FromJson("message1.json")

	//conn.Write([]byte("this is a client message"))
	//conn.Write([]byte(m.String()))

	// Send Message
	fmt.Println("Sending a message...")

	_, err = conn.Write([]byte(m.String()))

	// 	fmt.Println("Just wrote n = " + strconv.Itoa(n))
	if err != nil {
		panic("Error writing message")
	}

	// fmt.Println("About to wait for ACK")
	waitForACK(conn)
}

func waitForACK(conn net.Conn) {

	buf := make([]byte, 64)

	// fmt.Println("Inside waitForACK")
	for s := ""; s != "MESSAGE RECEIVED"; {
		s = read(conn, buf)
		//	fmt.Println("s = " + s)
		// fmt.Println(strconv.Itoa(len(s)) + " | " + strconv.Itoa(len("MESSAGE RECEIVED")))
	}

	fmt.Println("Broke out of loop, message receieved")

}

func read(conn net.Conn, buf []byte) string {
	_, err := bufio.NewReader(conn).Read(buf)

	if err != nil {
		fmt.Println("Error reading")
	}

	// fmt.Println("Just read n = " + strconv.Itoa(n))

	// trim unused bytes in buffer
	buf = bytes.Trim(buf, "\x00")
	// fmt.Println(buf)
	s := strings.TrimSpace(string(buf))
	// fmt.Println(s)

	return s

}
