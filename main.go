package main

import (
	"fmt"
	"mp0/processes"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("For client run: go run main.go -A\nFor server run: go run main.go -B")
}

func main() {
	const address = ":8080"

	fmt.Println(os.Args[1])

	args := os.Args

	fmt.Println(strconv.Itoa(len(args)))

	if len(args) != 2 {
		usage()
	} else if args[1] == "-B" {
		processes.StartServer(address)
	} else if args[1] == "-A" {
		processes.SendMessage("localhost" + address)
	} else {
		usage()
	}

}
