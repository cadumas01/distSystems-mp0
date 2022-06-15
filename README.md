# MP0 
A simple tcp server-client message program

## Usage

1. Run: ``go run main.go -B`` to start the server
2. Open a second shell, run ``go run main.go -A`` to start the client. Interact with the client's prompts to send a message to the server.

## Overview
processA is the client

processB is the server

1. Server is started and initializes a tcp listener
2. The client connects to the listener 
3. The server accepts the client
4. The client sends a message (via CLI or json file)
5. The server recieves the message, sends confirmation and exits
6. The client recieves confirmation and exits

## Resources / References
- https://pkg.go.dev/net
- https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/
- https://freshman.tech/snippets/go/read-console-input/
- https://tutorialedge.net/golang/parsing-json-with-golang/
