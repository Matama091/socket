package main

import (
	"fmt"
	"net"
)
type Packet struct {
	Version        uint8
	Type           uint8
	Flag           uint8
	Count          uint8
	TransactionID  uint8
}

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
	
	b := make([]byte, 1500)
	_, err = conn.Read(b)
	if err != nil {
		fmt.Println("Read error")
	}

	p := Packet{
		Version: b[0],
		Type: b[1],
		Flag: b[2],
		Count: b[3],
		TransactionID: b[4],
	}
	
	// fmt.Printf("hoge: %d\n", hoge)
	// Display packets
	fmt.Printf("aaa:%s", b)
	fmt.Printf("Version: %x\n", p.Version)
	fmt.Printf("Type: %x\n", p.Type)
	fmt.Printf("Flag: %x\n", p.Flag)
	fmt.Printf("Count: %x\n", p.Count)
	fmt.Printf("TransactionID: %x\n", p.TransactionID)
	conn.Write([]byte("pong"))

}