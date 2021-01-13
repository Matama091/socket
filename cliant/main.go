package main

import (
	"fmt"
	"io"
	"net"
)



func main() {
	conn, err := net.Dial("tcp", "localhost:4501")
	if err != nil {
		fmt.Println("Listen Error")
	}

	// 最後に実行
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("Error when TCP Listen closing...", err)
		}
	}()

	// Write
	for {
		_, err := conn.Write([]byte("ping"))
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
}





