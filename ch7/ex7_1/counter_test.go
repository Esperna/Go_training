package main

import (
	"fmt"
	"testing"
)

func TestByteCounterWrite(t *testing.T) {
	var c ByteCounter
	length, err := c.Write([]byte("hello"))
	if err != nil {
		t.Errorf("write failed. %s ", err)
	}
	if length != 5 {
		t.Errorf("length != 5. length:%d", length)

	}
	fmt.Printf("TestByteCounterWrite result:%d\n", c) // "5", = len("hello")
}

func TestWordCounterWrite(t *testing.T) {
	var c WordCounter
	length, err := c.Write([]byte("The Yellow Monkey"))
	if err != nil {
		t.Errorf("write failed. %s ", err)
	}
	if length != 3 {
		t.Errorf("length != 3. length:%d", length)

	}
	fmt.Printf("TestWordCounterWrite result:%d\n", c)
}
