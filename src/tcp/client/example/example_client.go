package main

import (
    "net"
    "fmt"
    "bufio"
    "encoding/binary"
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

        // Read a line from the OS
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("Text to send: ")

        text, _ := reader.ReadString('\n')
        fmt.Fprintf(conn, text + "\n")

        message, _ := bufio.NewReader(conn).ReadString('\n')
        length := len(message)

        fmt.Printf("Writing(%d): %v\n", length, message)
        encodedMessage := make(byte[], length + 4)
        copy(encodedMessage[4:], message)

        lengthBytes = encodedMessage[0:4]
        lengthBuf = bytes.NewBuffer(lengthBytes)

        binary.Write(lengthBuf, binary.LittleEndian, length)
    }
}

