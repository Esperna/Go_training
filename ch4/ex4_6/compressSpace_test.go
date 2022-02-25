package main

import (
	"testing"
)

func TestCompressSpace(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"あ　　い　　う", "あ い う"},
		{"a   b       c", "a b c"},
		{"a        b 　　c", "a b c"},
	}

	for _, test := range tests {
		b := []byte(test.input)
		b = compressUnicodeSpaces(b)
		if string(b) != test.want {
			t.Errorf("string(b) != %s. string(b) is %s", test.want, string(b))
		}
	}
}
