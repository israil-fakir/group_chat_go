package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "sync"
)

type Client struct {
    conn     net.Conn
    username string
}

var (
    clients   = make(map[net.Conn]Client) // track connected clients
    broadcast = make(chan string)         // messages to broadcast
    mu        sync.Mutex
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer ln.Close()
    fmt.Println("Chat server started on :8080")

    // goroutine: broadcast messages to all clients
    go handleBroadcast()

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        // First message should be the username
        fmt.Fprintln(conn, "Enter your username:")
        scanner := bufio.NewScanner(conn)
        scanner.Scan()
        username := strings.TrimSpace(scanner.Text())
        if username == "" {
            username = "Anonymous"
        }

        mu.Lock()
        clients[conn] = Client{conn: conn, username: username}
        mu.Unlock()

        fmt.Printf("New client joined: %s\n", username)
        broadcast <- fmt.Sprintf("%s joined the chat", username)

        go handleClient(conn, scanner)
    }
}

func handleClient(conn net.Conn, scanner *bufio.Scanner) {
    defer func() {
        mu.Lock()
        username := clients[conn].username
        delete(clients, conn)
        mu.Unlock()
        conn.Close()
        fmt.Printf("%s left the chat\n", username)
        broadcast <- fmt.Sprintf("%s left the chat", username)
    }()

    for scanner.Scan() {
        msg := strings.TrimSpace(scanner.Text())
        if msg == "" {
            continue // ignore empty messages
        }
        broadcast <- fmt.Sprintf("%s: %s", clients[conn].username, msg)
    }
}

func handleBroadcast() {
    for msg := range broadcast {
        mu.Lock()
        for _, client := range clients {
            fmt.Fprintln(client.conn, msg)
        }
        mu.Unlock()
    }
}
