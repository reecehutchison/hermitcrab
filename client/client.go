package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	b := make([]byte, 128)
	conn.Write([]byte("Hi Reece"))
	n, _ := conn.Read(b)
	fmt.Printf("%s\n", b[:n])
	conn.Write([]byte("Bye Reece"))
	n, _ = conn.Read(b)
	fmt.Printf("%s\n", b[:n])
}