package main

import (
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
