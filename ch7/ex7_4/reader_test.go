package main

import "testing"

func TestNewReader(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"hello", 5},
		{"<html><title>Under Construction</title></html>", 46},
		{"", 0},
	}
	for _, test := range tests {
		rd := NewReader(test.input)
		if rd.s != test.input && rd.i != int64(test.want) {
			t.Errorf("NewReader result is not expected.  expected: %v actual: %v", test, rd)
		}
	}

}

func TestRead(t *testing.T) {
	type expect struct {
		n   int
		err error
		s   string
	}
	var tests = []struct {
		input string
		want  expect
	}{
		{"hello", expect{5, nil, "hello"}},
		{"<html><title>Under Construction</title></html>", expect{46, nil, "<html><title>Under Construction</title></html>"}},
		{"", expect{0, nil, ""}},
	}
	for _, test := range tests {
		rd := NewReader(test.input)
		var b []byte
		n, err := rd.Read(b)
		if n != test.want.n && err != test.want.err && string(b) != test.want.s {
			t.Errorf("Read result is not expected.  expected: %v, actual: %v %v", test.want, n, err)
		}
	}

}

//func (r *Reader) Read(b []byte) (n int, err error) {
