# MP0 
A simple tcp server-client message program

## Usage

Run: ``go run main.go``

## Overview
processA is the client
processB is the server

1. TCP server is initalized, a net.Listener is returned
2. A goroutine is started for the server to accept client connections to that net.Listener
3. The client connects to the listener 
4. The server accepts the client
5. The client sends a message (via CLI or json file)
6. The server recieves the message, sends confirmation and exits
7. The client recieves confirmation and exits

## Resources / References
- https://pkg.go.dev/net
- https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/
- https://freshman.tech/snippets/go/read-console-input/
- https://tutorialedge.net/golang/parsing-json-with-golang/
