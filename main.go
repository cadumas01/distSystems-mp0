package main

import (
	"fmt"

	"./mp0/processes"
)

func usage() {
	fmt.Println("Enter a message \nUsage: TO FROM DATE TITLE CONTENT")
}

func main() {
	const address = "8080"
	//m := messages.FromJson("message1.json")

	processes.StartServer(address)

}
