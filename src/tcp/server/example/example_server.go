package main

import (
    "fmt"
    "net"
    "bufio"
    "strings"
)

func main() {
    fmt.Println("About to start the server...")

    ln, listenErr := net.Listen("tcp", ":33333")
    if listenErr != nil {
        fmt.Println("Unable to start the server.")
        return
    }

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
}

