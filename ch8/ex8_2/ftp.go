package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
		go handleControlConn(conn)
	}
}

func handleControlConn(c net.Conn) {
	defer c.Close()
	log.Printf("Accept %v\n", c.RemoteAddr())
	_, err := io.WriteString(c, respMsg(220))
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	reader := bufio.NewReader(c)
	var dp dataPort
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
		if len(str) == 1 {
			str = strings.Split(str[0], "\r")
			str = strings.Split(str[0], "\n")
		}
		log.Printf("str: %v len: %d\n", str, len(str))
		cmd := str[0]
		switch cmd {
		case "USER":
			if _, err := io.WriteString(c, respMsg(331)); err != nil {
				log.Printf("%v\n", err)
				break
			}
		case "PASS":
			if _, err := io.WriteString(c, respMsg(230)); err != nil {
				log.Printf("%v\n", err)
				break
			}
		case "QUIT":
			if _, err := io.WriteString(c, respMsg(221)); err != nil {
				log.Printf("%v\n", err)
			}
		case "PASV":
			A := strings.Split("127.0.0.1", ".")
			a := []int{2, 0} //Port:20
			msg := fmt.Sprintf("%s %s,%s,%s,%s,%d,%d\n", responses[227], A[0], A[1], A[2], A[3], a[0], a[1])
			log.Printf("%v\n", msg)
			if _, err := io.WriteString(c, msg); err != nil {
				log.Printf("%v\n", err)
				break
			}
		case "SYST":
			if _, err := io.WriteString(c, respMsg(215)); err != nil {
				log.Printf("%v\n", err)
			}
			break
		case "PORT":
			if _, err := fmt.Sscanf(str[1], "%d,%d,%d,%d,%d,%d\n", &dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2); err != nil {
				log.Printf("Sscanf failed: %v\n", err)
				break
			}
			log.Printf("%v\n", dp)
			if _, err := io.WriteString(c, respMsg(200)); err != nil {
				log.Printf("%v\n", err)
				break
			}
		case "FEAT":
			if _, err := io.WriteString(c, respMsg(202)); err != nil {
				log.Printf("%v\n", err)
				break
			}
		case "LIST":
			files, err := ioutil.ReadDir("./")
			if err != nil {
				log.Printf("%s\n", err)
				break
			}
			if _, err := io.WriteString(c, respMsg(150)); err != nil {
				log.Printf("%s\n", err)
				break
			}
			dataConn, err := net.Dial("tcp", dp.toAddress())
			if err != nil {
				log.Printf("%s\n", err)
				break

			}
			for _, file := range files {
				if _, err := fmt.Fprintf(dataConn, "%s\r\n", file.Name()); err != nil {
					log.Printf("%s\n", err)
				}
			}
			if _, err := io.WriteString(c, respMsg(226)); err != nil {
				log.Printf("%s\n", err)
				dataConn.Close()
				break
			}
			dataConn.Close()
			break
		default:
			// if _, err := io.WriteString(c, respMsg(202)); err != nil {
			// 	log.Printf("%v\n", err)
			// 	break
			// }
		}
	}
}
