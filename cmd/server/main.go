package main

import (
	"bufio"
	"fmt"
	"net"
)

// handleConnection proccesses incoming TCP connections.
// It reads them, prints them out, and sends a response back.
func handleConnection(conn net.Conn) {
	// Here to ensure file is closed when function ends
	defer conn.Close()

	// Create a new reader to read from the TCP connection
	r := bufio.NewReader(conn)

	// Declares variables to handle continous reading until the error is returned
	var message string
	var err error

	// Continuously read from the connection until an error occurs
	for ; err == nil; {
		// Read a line from the connection
		message, err = r.ReadString('\n')
		if err != nil {
			fmt.Printf("Error receiving messages: %v\n", err)
			return
		}

		fmt.Printf("Received: %s\n", message)

		response := "Message: " + message

		// Sends response of the client's message back to the client as response
		_, err := conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			// Early return in case of error
			return
		}
	}
}

func main() {

	// Creates server to 8080 port on localhost
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("There was an error listening to the server: %v\n", err)
		panic(err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("There was an error connecting to the server: %v\n", err)
			continue
		}

		// Checks to ensure connection to the correct port
		address := conn.RemoteAddr()
		fmt.Printf("Address: %s\n", address.String())

		go handleConnection(conn)
	} 
}