// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		echo := func(conn net.Conn, shout string, delay time.Duration) {
			defer wg.Done()
			fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
			time.Sleep(delay)
			fmt.Fprintln(conn, "\t", shout)
			time.Sleep(delay)
			fmt.Fprintln(conn, "\t", strings.ToLower(shout))
		}
		go echo(conn, input.Text(), 1*time.Second)
	}
	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	c, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("conn doesn't have net.TCPConn")
	}
	c.CloseWrite()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
