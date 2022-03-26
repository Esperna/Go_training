package main

import (
	"bufio"
	"fmt"
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
	_, err := io.WriteString(c, respMsg(220))
	if err != nil {
		log.Printf("%s", err)
		return
	}
	reader := bufio.NewReader(c)
	for {
		commands := map[string]func(net.Conn, []string) error{
			"USER": user,
			"PASS": pass,
			"QUIT": quit,
			"PASV": pasv,
			"SYST": syst,
			"PORT": port,
			"FEAT": feat,
			"LIST": list,
			"RETR": retr,
			"CWD":  cwd,
			"STOR": stor,
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("%s\n", err)
			}
			break
		}
		msg := strings.Split(line, " ")
		var name string
		if len(msg) == 1 {
			if _, err := fmt.Sscanf(msg[0], "%s\n", &name); err != nil {
				log.Printf("Sscanf failed: %s", err)
			}
		} else {
			name = msg[0]
		}
		log.Printf("msg: %s len: %d\n", name, len(msg))
		if err := commands[name](c, msg); err != nil {
			log.Printf("%v\n", err)
			break
		}
	}
}
