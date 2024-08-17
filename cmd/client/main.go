package main

import (
	"net"
	"bufio"
	"fmt"
)

func main() {
	// connect to port 8080 server
	// write a message
	// i may need to connect to a bufio reader to write a message

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("There was an error connecting to the server: %v", err)
	}

	conn.Write([]byte("Hello Word\n"))
	res, e := bufio.NewReader(conn).ReadString('\n')

	if e != nil {
		fmt.Printf("There was an error connecting to the server: %v\n", err)
	} else {
		fmt.Printf("Server Response: %v", res)
	}
}