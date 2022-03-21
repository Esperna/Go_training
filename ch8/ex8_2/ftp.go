package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

var responses = map[int]string{
	150: "File status okay; about to open data connection.",
	200: "Command okay.",
	202: "Command not implemented, superfluous at this site.",
	215: "NAME system type. Where NAME is an official system name from the list in the Assigned Numbers document.",
	220: "Service ready for new user.",
	221: "Service closing control connection.",
	226: "Closing data connection. Requested file action successful (for example, file transfer or file abort).",
	227: "Entering Passive Mode.",
	230: "User logged in, proceed",
	331: "User name okay, need password.",
	501: "Syntax error in parameters or arguments.",
	530: "Not Logged in.",
}

type dataPort struct {
	h1, h2, h3, h4 int
	p1, p2         int
}

func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}
	port := d.p1<<8 + d.p2
	return fmt.Sprintf("%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, port)
}

func respMsg(code int) string {
	return strconv.Itoa(code) + " " + responses[code] + "\n"
}

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

var dp dataPort //exclusive Control is needed

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
		if len(msg) == 1 {
			msg = strings.Split(msg[0], "\r")
			msg = strings.Split(msg[0], "\n")
		}
		log.Printf("msg: %s len: %d\n", msg[0], len(msg))
		name := msg[0]
		if err := commands[name](c, msg); err != nil {
			log.Printf("%v\n", err)
			break
		}
	}
}
