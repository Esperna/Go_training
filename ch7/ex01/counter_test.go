package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

type CounterType int

func TestCounterWrite(t *testing.T) {
	const (
		Byte CounterType = iota
		Word
		Line
	)
	var tests = []struct {
		input     string
		countType CounterType
		want      int
	}{
		{"hello", Byte, 5},
		{"The Yellow Monkey", Word, 3},
		{"This is a pen.\nThis is an apple.\nAh, apple pen.\n", Line, 3},
	}
	//It's not counter value check but return value check
	for _, test := range tests {
		var length int
		var err error
		if test.countType == Byte {
			var c ByteCounter
			length, err = c.Write([]byte(test.input))
		} else if test.countType == Word {
			var c WordCounter
			length, err = c.Write([]byte(test.input))
		} else if test.countType == Line {
			var c LineCounter
			length, err = c.Write([]byte(test.input))
		}
		if err != nil {
			t.Errorf("write failed. %s ", err)
		}
		if length != test.want {
			t.Errorf("length != test.want. length:%d test.want:%d", length, test.want)
		}
	}
}

func TestByteCounterValue(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"hello", 5},
		{"hello", 10},
		{"hello", 15},
	}
	var c ByteCounter
	for _, test := range tests {
		c.Write([]byte(test.input))
		if c != ByteCounter(test.want) {
			t.Errorf("count != test.want. count:%d test.want:%d", c, test.want)
		}
	}
}

func TestWordCounterValue(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"The Yellow Monkey", 3},
		{"The Yellow Monkey", 6},
		{"The Yellow Monkey", 9},
	}
	var c WordCounter
	for _, test := range tests {
		c.Write([]byte(test.input))
		if c != WordCounter(test.want) {
			t.Errorf("count != test.want. count:%d test.want:%d", c, test.want)
		}
	}
}

func TestLineCounterValue(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"This is a pen.\nThis is an apple.\nAh, apple pen.\n", 3},
		{"This is a pen.\nThis is an apple.\nAh, apple pen.\n", 6},
		{"This is a pen.\nThis is an apple.\nAh, apple pen.\n", 9},
		{"This is a pen.\nThis is an apple.\nAh, apple pen.\n", 12},
	}
	var c LineCounter
	for _, test := range tests {
		c.Write([]byte(test.input))
		if c != LineCounter(test.want) {
			t.Errorf("count != test.want. count:%d test.want:%d", c, test.want)
		}
	}
}

func TestTooLargeWordCounterValue(t *testing.T) {
	var c WordCounter
	input := []byte("Word1 " + strings.Repeat("X", bufio.MaxScanTokenSize+1) + " Word2")
	count, err := c.Write(input)
	fmt.Printf("%d %v", count, err)
	expectedCount := 1
	if count != expectedCount {
		t.Errorf("count is not expected. count:%d test.want:%d", count, expectedCount)
	}
	expectedErr := fmt.Errorf("bufio.Scanner: token too long")
	if err.Error() != expectedErr.Error() {
		t.Errorf("err is not expectd. \n%s\n%s", err.Error(), expectedErr.Error())
	}
}
