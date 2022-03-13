package main

import "testing"

func TestNewReader(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"hello", 5},
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
	}
	var tests = []struct {
		input []byte
		want  expect
	}{
		{[]byte("hello"), expect{5, nil}},
		{[]byte("<html><title>Under Construction</title></html>"), expect{46, nil}},
		{[]byte(""), expect{0, nil}},
	}
	for _, test := range tests {
		var r Reader
		n, err := r.Read(test.input)
		if n != test.want.n && err != test.want.err {
			t.Errorf("Read result is not expected.  expected: %v, actual: %v %v", test.want, n, err)
		}
	}

}

//func (r *Reader) Read(b []byte) (n int, err error) {
