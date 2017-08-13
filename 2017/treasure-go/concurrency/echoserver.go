package main

import (
	"io"
	"log"
	"net"
)

// use `nc localhost 8888` for debug
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	io.Copy(c, c)
	c.Close()
}
