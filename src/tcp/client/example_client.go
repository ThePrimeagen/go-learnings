package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:33333")

	if err != nil {
		fmt.Println("Found an error!")
		fmt.Println(err)

		return
	}

	for {
		fmt.Print("Text to send: ")

		// Read a line from the OS
		reader := bufio.NewReader(os.Stdin)

		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		length := len(text)

		fmt.Printf("Writing(%d): %v\n", length, text)
		encodedMessage := make([]byte, length+4)
		copy(encodedMessage[4:], text)
		fmt.Printf("encodedMessage %v\n", encodedMessage)

		lengthBytes := encodedMessage[0:4]

		binary.LittleEndian.PutUint32(lengthBytes, uint32(length))
		fmt.Printf("encodedMessage again %v\n", encodedMessage)

		n, e := conn.Write(encodedMessage)

		if e != nil {
			fmt.Printf("Writing the text message has failed with %v\n", e)
		}

		if n != length+4 {
			fmt.Printf("Expected to write %d but wrote %d\n", length+4, n)
		}
	}
}
