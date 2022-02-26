package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"太陽", "陽太"},
		{"太平洋", "洋平太"},
		{"abcdef", "fedcba"},
		{"αβγ", "γβα"},
		{"こんにちはHello", "olleHはちにんこ"},
	}

	for _, test := range tests {
		b := []byte(test.input)
		b = reverse(b)
		if string(b) != test.want {
			t.Errorf("string(b) != %s. string(b) is %s", test.want, string(b))
		}
	}
}
