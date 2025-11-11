package main

import (
    "io"
    "net"
)

func handleClient(clientConn net.Conn) {
    defer clientConn.Close()

    remoteConn, _ := net.Dial("tcp", "localhost:8000")
    defer remoteConn.Close()

    go func() {
        io.Copy(remoteConn, clientConn)
    }()

    io.Copy(clientConn, remoteConn)
}

func main() {
    listener, _ := net.Listen("tcp", ":8001")
    defer listener.Close()

    for {
        clientConn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(clientConn)
    }
}
