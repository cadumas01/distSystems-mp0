package main

import (
	"mp0/processes"
)

func main() {
	const address = ":8080"

	ln := processes.StartServer(address)

	go processes.AcceptClient(ln)

	processes.SendMessage("localhost" + address)
}
