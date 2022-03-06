// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	go netcat("localhost:8010")
	go netcat("localhost:8020")
	go netcat("localhost:8030")
	conn1, err := net.Dial("tcp", "localhost:8010")
	if err != nil {
		log.Fatal(err)
	}
	conn2, err := net.Dial("tcp", "localhost:8020")
	if err != nil {
		log.Fatal(err)
	}
	conn3, err := net.Dial("tcp", "localhost:8030")
	if err != nil {
		log.Fatal(err)
	}
	defer conn1.Close()
	defer conn2.Close()
	defer conn3.Close()

	for {

	}
	// concat(conn1, conn2, conn3)
	// mustCopy(os.Stdout, conn1)
	// mustCopy(os.Stdout, conn2)
	// mustCopy(os.Stdout, conn3)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func concat(conn1, conn2, conn3 io.Reader) {
	rd := io.MultiReader(conn1, conn2, conn3)
	b, _ := ioutil.ReadAll(rd)
	os.Stdout.Write(b)

}

func netcat(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

//!-
