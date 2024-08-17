package main

import (
	"net"
	"bufio"
	"fmt"
)

func sendAndReceive(conn net.Conn) {
	// Writes a message to the server
	conn.Write([]byte("Hello Word\n"))

	// Handles reading messages sent back to client
	response, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error connecting to the server: %v\n", err)
	} else {
		// Writes the server's response as a 'Server Response'
		fmt.Printf("Server Response: %v", response)
	}
}

func main() {
	// Connects to the server on port 8080
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("There was an error connecting to the server: %v", err)
	}

	sendAndReceive(conn)
}