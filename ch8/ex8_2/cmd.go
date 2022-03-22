package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

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

var (
	dp dataPort //exclusive Control is needed
)

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

	dataConn, err := net.Dial("tcp", dp.toAddress())
	defer dataConn.Close()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if !strings.HasPrefix(m, "./") {
		m = "./" + m
	}
	files, err := ioutil.ReadDir(m)
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
	var m string
	if len(msg) <= 1 {
		if _, err := io.WriteString(c, respMsg(501)); err != nil {
			return fmt.Errorf("%s", err)
		}
		return fmt.Errorf("Invalid argument")
	}
	if _, err := fmt.Sscanf(msg[1], "%s\n", &m); err != nil {
		return fmt.Errorf("Sscanf failed: %s", err)
	}

	dataConn, err := net.Dial("tcp", dp.toAddress())
	defer dataConn.Close()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if !strings.HasPrefix(m, "./") {
		m = "./" + m
	}

	file, err := os.Open(m)
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
	var m string
	if len(msg) <= 1 {
		if _, err := io.WriteString(c, respMsg(501)); err != nil {
			return fmt.Errorf("%s", err)
		}
		return fmt.Errorf("Invalid argument")
	}

	if _, err := fmt.Sscanf(msg[1], "%s\n", &m); err != nil {
		return fmt.Errorf("Sscanf failed: %s", err)
	}

	if !strings.HasPrefix(m, "./") {
		m = "./" + m
	}
	log.Printf("%s\n", m)
	if err := os.Chdir(m); err != nil {
		return fmt.Errorf("%s", err)
	}

	if _, err := io.WriteString(c, respMsg(250)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func stor(c net.Conn, _ []string) error {
	if _, err := io.WriteString(c, respMsg(202)); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
