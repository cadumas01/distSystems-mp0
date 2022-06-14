package main

import (
	"fmt"
	"mp0/processes"
)

func usage() {
	fmt.Println("Enter a message \nUsage: TO FROM DATE TITLE CONTENT")
}

func main() {
	// fmt.Println("line 13")
	const address = ":8080"

	//channel := make(chan messages.Message)

	ln := processes.StartServer(address)

	go processes.AcceptClient(ln)

	// fmt.Println("line 21")
	// load message into channel and send it
	//channel <- *m
	processes.SendMessage("localhost" + address)

	// fmt.Println("Line 27")
}
