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
		reverse(b)
		if string(b) != test.want {
			t.Errorf("string(b) != %s. string(b) is %s", test.want, string(b))
		}
	}
}

func TestRotate(t *testing.T) {
	var tests = []struct {
		input     string
		rotateNum int
		want      string
	}{
		{"abc", 1, "bca"},
		{"abc", 2, "cab"},
		{"あbc", 3, "bcあ"},
	}

	for _, test := range tests {
		b := []byte(test.input)
		rotate(b, test.rotateNum)
		if string(b) != test.want {
			t.Errorf("string(b) != %s. string(b) is %s", test.want, string(b))
		}
	}
}
