package main

import (
	"fmt"
	"testing"
)

func TestCompressSpace(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"あ　　い　　う", "あ い う"},
		{"a   b       c", "a b c"},
		{"a        b 　　c", "a b c"},
		{"a", "a"},
	}
	for _, test := range tests {
		b := []byte(test.input)
		b = compressUnicodeSpaces(b)
		if string(b) != test.want {
			t.Errorf("string(b) != %s. string(b) is %s", test.want, string(b))
		}
	}
}

func TestFunction(t *testing.T) {
	b := []byte{0xff}
	defer func() {
		switch p := recover(); p {
		case fmt.Sprintf("Invalid byte s %q %d", b, 1):
			fmt.Printf("This panic is expected by test\n")
		default:
			panic(p)
		}
	}()
	compressUnicodeSpaces(b)
}
