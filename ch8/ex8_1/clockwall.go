// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Clock struct {
	index int
	time  Time
	addr  string
}

type Time struct {
	zone  string
	value string
}

func main() {
	clocks, numOfClock := parse(os.Args)
	clk := make(chan Clock, numOfClock)
	for _, clock := range clocks {
		go readWorldTime(clock, clk)
	}

	for {
		m := make(map[int]Time)
		for i := 0; i < numOfClock; i++ {
			clock := <-clk
			m[clock.index] = clock.time
		}
		timelist := ""
		for i := 0; i < len(m); timelist += " " {
			timelist += m[i].zone + " " + m[i].value
			i++
		}
		fmt.Printf("\r%s", timelist)
	}
}

func parse(args []string) ([]Clock, int) {
	var clocks []Clock
	for i := 1; i < len(args); i++ {
		var clock Clock
		slices := strings.Split(args[i], "=")
		clock.index = i - 1
		clock.time.zone = slices[0]
		clock.addr = slices[1]
		clocks = append(clocks, clock)
	}
	numOfClock := len(args) - 1
	return clocks, numOfClock
}

func readWorldTime(clock Clock, clk chan<- Clock) {
	conn, err := net.Dial("tcp", clock.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		var c Clock
		c.index = clock.index
		c.time.zone = clock.time.zone
		c.time.value = scanner.Text()
		clk <- c
	}
}

//!-
