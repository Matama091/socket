package main

import (
	"fmt"
	"net"
)

func main() {
	// TCP Listen
	listen, err := net.Listen("tcp", "localhost:4501")
	if err != nil {
		fmt.Println("Listen Error")
	}

	// Accept
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("accept error")
	}

	defer func() {
		if err := listen.Close(); err != nil {
			fmt.Println("Error when TCP Listen closing...", err)
		}
	}()

	// Read
	for {
		buffer := make([]byte, 1500)
		_, err := conn.Read(buffer)
		if err != nil {
		}
		// fmt.Printf("hoge: %d\n", hoge)
		// Display packets
		fmt.Printf("%s", buffer)
		conn.Write([]byte("pong"))
	}
}
