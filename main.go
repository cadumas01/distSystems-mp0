package main

import (
	"mp0/processes"
	"sync"
)

func main() {
	const address = ":8080"

	wg := new(sync.WaitGroup)

	wg.Add(2) // add 2 for the two goroutines

	ln := processes.StartServer(address)

	go processes.AcceptClient(ln, wg)

	go processes.SendMessage("localhost"+address, wg)

	wg.Wait()
}
