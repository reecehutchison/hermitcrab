package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("hi")
}

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}