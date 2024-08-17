package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Defers connection until EOF
	defer conn.Close()

	// Creates new reader to read from TCP connection
	r := bufio.NewReader(conn)

	// Initalize v to empty string
	// Initalize e to empty error
	var v string
	var e error

	// As long as e is nil run code
	for ; e == nil; {
		// Reads for strings until it gets to new line
		v,e = r.ReadString('\n')
		// Checks if there is an error
		if e != nil {
			fmt.Printf("Error receiving messages: %v\n", e)
			// Early return in case of error
			return
		}
		fmt.Printf("Received %s\n", v)

		response := "Message: " + v

		_, err := conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			// Early return in case of error
			return
		}
	}
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("There was an error listening to the server")
		panic(err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("There was an error connecting to the server")
			continue
		}

		address := conn.RemoteAddr()
		fmt.Printf("Address: %s\n", address.String())

		go handleConnection(conn)
	} 
}