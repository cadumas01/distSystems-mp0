package main

import (
	"fmt"
	"test/mp0/messages"
)

func usage() {
	fmt.Println("Enter a message \nUsage: TO FROM DATE TITLE CONTENT")
}

func main() {
	m := messages.FromJson("message1.json")
	fmt.Println(m.Content, "3")

}
