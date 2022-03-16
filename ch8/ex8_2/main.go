package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	fmt.Printf("Accept %v\n", c.RemoteAddr())
	defer c.Close()
	for {
		_, err := io.WriteString(c, "220 Service ready for new user.\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
