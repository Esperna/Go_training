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
		//line, err := reader.ReadString('\n')
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Printf("%s\n", err)
			}
			break
		}
		if isPrefix {
			log.Printf("line is too long\n")
			break
		}
		msg := strings.Split(string(line), " ")
		name := msg[0]

		log.Printf("msg: %s len: %d\n", name, len(msg))
		if err := commands[name](c, msg); err != nil {
			log.Printf("%v\n", err)
			break
		}
	}
}
