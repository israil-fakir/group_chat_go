package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // Ask for username locally
    fmt.Print("Enter your username: ")
    inputScanner := bufio.NewScanner(os.Stdin)
    inputScanner.Scan()
    username := inputScanner.Text()

    // Send username first
    fmt.Fprintln(conn, username)

    // Goroutine: print messages from server
    go func() {
        serverScanner := bufio.NewScanner(conn)
        for serverScanner.Scan() {
            fmt.Println(serverScanner.Text())
        }
    }()

    // Main loop: read user input and send to server
    for inputScanner.Scan() {
        fmt.Fprintln(conn, inputScanner.Text())
    }
}
