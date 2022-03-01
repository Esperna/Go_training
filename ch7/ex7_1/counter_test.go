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

func TestLineCounterWrite(t *testing.T) {
	var c LineCounter
	length, err := c.Write([]byte("This is a pen.\nThis ia an apple.\nAh, apple pen.\n"))
	if err != nil {
		t.Errorf("write failed. %s ", err)
	}
	if length != 3 {
		t.Errorf("length != 3. length:%d", length)

	}
	fmt.Printf("TestLineCounterWrite result:%d\n", c)

}
