package main

import (
    "fmt"
    "net"
    "bufio"
    "strings"
    "github.com/michaelbpaulson/tcp"
)

func main() {
    fmt.Println("About to start the server...")

    ln, listenErr := net.Listen("tcp", ":33333")
    if listenErr != nil {
        fmt.Println("Unable to start the server.")
        return
    }

    outChannel := tcp.FrameConnections(ln)

    // Every message of Responder will come through here
    fmt.Println("About to range over channel")
    for message := range outChannel {
        fmt.Printf("Received message %v\n", message)
        conn.Write([]byte(newMessage + "\n"))
    }

    /*
    fmt.Println("Now Listening...")
    conn, acceptErr := ln.Accept()
    if acceptErr != nil {
        fmt.Println("Unable to accept connection.")
    }

    fmt.Println("Accept() has finished")
    for {
        fmt.Println("Waiting for message...")
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println("Received: %s", string(message))
        newMessage := strings.ToUpper(message)
        conn.Write([]byte(newMessage + "\n"))
    }
    */
}

