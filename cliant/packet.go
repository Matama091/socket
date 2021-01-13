package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"encoding/binary"
)

type Packet struct {
	Version        uint8
	Type           uint8
	Flag           uint8
	Count          uint8
	TransactionID  uint8
}



func main() {

	p := Packet{
		Version: 1,
		Type: 1,
		Flag: 0,
		Count: 1,
		TransactionID: 110,
	}

	// TCP Listen
	conn, err := net.Dial("tcp", "localhost:4501")
	if err != nil {
		fmt.Println("Listen Error")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("Error when TCP Listen closing...", err)
		}
	}()

	// Write
	b := make([]byte, 0, 5)
	buf := bytes.NewBuffer(b)

	if err := binary.Write(buf, binary.BigEndian, p.Version); err != nil {
		fmt.Println("Error Version...", err)
	}
	if err := binary.Write(buf, binary.BigEndian, p.Type); err != nil {
		fmt.Println("Error Type...", err)
	}
	if err := binary.Write(buf, binary.BigEndian, p.Flag); err != nil {
		fmt.Println("Error Flag...", err)
	}
	if err := binary.Write(buf, binary.BigEndian, p.Count); err != nil {
		fmt.Println("Error Count...", err)
	}
	if err := binary.Write(buf, binary.BigEndian, p.TransactionID); err != nil {
		fmt.Println("Error TransactionID...", err)
	}

	fmt.Printf(" % x\n", buf.Bytes())

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		if err == io.EOF {
			return
		}
		return
	}
	buffer := make([]byte, 1500)
	_, err = conn.Read(buffer)
	fmt.Printf("%s", buffer)
	// fmt.Printf("hoge: %d\n", hoge)
	// Display packets

}
