package main

import (
	"fmt"
	"net"
	"os"
)

const port string = "8080"

func handleConnection(conn net.Conn) {
	b := make([]byte, 128)
	n, _ := conn.Read(b)
	fmt.Printf("%q\n", b[:n])
}

func main() {
	ln, err := net.Listen("tcp", ":" + port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	fmt.Println("Server is listening on port", port + "...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		go handleConnection(conn)
	}
}