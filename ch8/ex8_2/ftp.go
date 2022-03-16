package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
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
	defer c.Close()
	log.Printf("Accept %v\n", c.RemoteAddr())
	_, err := io.WriteString(c, "220 Service ready for new user.\n")
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	reader := bufio.NewReader(c)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("%v\n", err)
			}
			break
		}
		log.Printf("line: %s\n", line)
		str := strings.Split(line, " ")
		if len(str) > 2 {
			_, err := io.WriteString(c, "202 Command not implemented, superfluous at this site.\n")
			if err != nil {
				return
			}
		}
		cmd := str[0]
		switch cmd {
		case "USER":
			_, err := io.WriteString(c, "331 User name okay, need password.\n")
			if err != nil {
				log.Printf("%v\n", err)
				break
			}
		case "PASS":
			_, err := io.WriteString(c, "230 User logged in, proceed.\n")
			if err != nil {
				log.Printf("%v\n", err)
				break
			}
		default:
			_, err := io.WriteString(c, "202 Command not implemented, superfluous at this site.\n")
			if err != nil {
				break
			}
		}
	}
}
