package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

const port string = "8080"

func handleConnection(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 128)
	for {
		n, err := conn.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		fmt.Printf("client '%s' : '%s'\n", conn.RemoteAddr().String(), b[:n])
		conn.Write(b[:n])
	}
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