package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
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
	250: "Requested file action okay, completed.",
	331: "User name okay, need password.",
	501: "Syntax error in parameters or arguments.",
	530: "Not Logged in.",
}

var (
	dp dataPort //exclusive Control is needed
)

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

func user(c net.Conn, _ []string) error {
	if _, err := io.WriteString(c, respMsg(331)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func pass(c net.Conn, msg []string) error {
	var m string
	if _, err := fmt.Sscanf(msg[1], "%s\n", &m); err != nil {
		return fmt.Errorf("Sscanf failed: %s", err)
	}
	password := "huga"
	if m == password {
		if _, err := io.WriteString(c, respMsg(230)); err != nil {
			return fmt.Errorf("%s", err)
		}
	} else {
		if _, err := io.WriteString(c, respMsg(530)); err != nil {
			return fmt.Errorf("%s", err)
		}
	}
	return nil
}

func quit(c net.Conn, _ []string) error {
	if _, err := io.WriteString(c, respMsg(221)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func pasv(c net.Conn, _ []string) error {
	A := strings.Split("127.0.0.1", ".")
	a := []int{2, 0} //Port:20
	msg := fmt.Sprintf("%s %s,%s,%s,%s,%d,%d\n", responses[227], A[0], A[1], A[2], A[3], a[0], a[1])
	if _, err := io.WriteString(c, msg); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func syst(c net.Conn, _ []string) error {
	if _, err := io.WriteString(c, respMsg(215)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func feat(c net.Conn, _ []string) error {
	if _, err := io.WriteString(c, respMsg(202)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func port(c net.Conn, msg []string) error {
	if _, err := fmt.Sscanf(msg[1], "%d,%d,%d,%d,%d,%d\n", &dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2); err != nil {
		return fmt.Errorf("Sscanf failed: %s", err)
	}
	log.Printf("%v\n", dp)
	if _, err := io.WriteString(c, respMsg(200)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func list(c net.Conn, msg []string) error {
	var m string
	if len(msg) > 1 {
		if _, err := fmt.Sscanf(msg[1], "%s\n", &m); err != nil {
			return fmt.Errorf("Sscanf failed: %s", err)
		}
	}
	if !strings.HasPrefix(m, "./") {
		m = "./" + m
	}

	files, err := ioutil.ReadDir(m)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	dataConn, err := net.Dial("tcp", dp.toAddress())
	defer dataConn.Close()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	if _, err := io.WriteString(c, respMsg(150)); err != nil {
		return fmt.Errorf("%s", err)
	}
	for _, file := range files {
		if _, err := fmt.Fprintf(dataConn, "%s\r\n", file.Name()); err != nil {
			return fmt.Errorf("%s", err)
		}
	}
	if _, err := io.WriteString(c, respMsg(226)); err != nil {
		log.Printf("%s\n", err)
	}
	return nil
}

func retr(c net.Conn, msg []string) error {
	path, err := preCmdCheckForAtLeast1Arg(len(msg), msg[1])
	if err != nil {
		if _, err := io.WriteString(c, respMsg(501)); err != nil {
			return err
		}
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	dataConn, err := net.Dial("tcp", dp.toAddress())
	defer dataConn.Close()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	if _, err := io.WriteString(c, respMsg(150)); err != nil {
		return fmt.Errorf("%s", err)
	}
	if _, err := io.Copy(dataConn, file); err != nil {
		return fmt.Errorf("%s", err)
	}

	if _, err := io.WriteString(c, respMsg(226)); err != nil {
		log.Printf("%s\n", err)
	}

	return nil
}

func cwd(c net.Conn, msg []string) error {
	path, err := preCmdCheckForAtLeast1Arg(len(msg), msg[1])
	if err != nil {
		if _, err := io.WriteString(c, respMsg(501)); err != nil {
			return err
		}
	}
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("%s", err)
	}

	if _, err := io.WriteString(c, respMsg(250)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func preCmdCheckForAtLeast1Arg(length int, msg string) (string, error) {
	if length <= 1 {
		return "", fmt.Errorf("Invalid argument")
	}
	var m string
	if _, err := fmt.Sscanf(msg, "%s\n", &m); err != nil {
		return "", fmt.Errorf("Sscanf failed: %s", err)
	}
	if !strings.HasPrefix(m, "./") {
		m = "./" + m
	}
	return m, nil
}

func stor(c net.Conn, _ []string) error {
	if _, err := io.WriteString(c, respMsg(202)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
