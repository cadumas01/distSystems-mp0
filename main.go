package main

import (
	"fmt"

	"github.com/cadumas01/distSystems-mp0/messages"
)

func usage() {
	fmt.Println("Enter a message \nUsage: TO FROM DATE TITLE CONTENT")
}

func main() {
	m := messages.FromJson("message1.json")
	fmt.Println(m.Content, "3")

}
