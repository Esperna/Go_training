package main

import (
	"io"
	"strings"
	"testing"
)

func TestReadSingle(t *testing.T) {
	type Output struct {
		n   int
		err error
	}
	type Input struct {
		str string
		val int
	}
	var tests = []struct {
		expected Output
		given    Input
	}{
		{Output{5, nil}, Input{"Hello Hello", 5}},
	}
	for _, test := range tests {
		r := strings.NewReader(test.given.str)
		lr := LimitReader(r, int64(test.given.val))
		p := make([]byte, test.given.val)
		n, err := lr.Read(p)
		if err != test.expected.err {
			t.Errorf("Read error:%s", err)
		}
		if n != test.expected.n {
			t.Errorf("Not expected n. Expected %d Actual %d", test.expected.n, n)
		}
	}
}

func TestReadTwice(t *testing.T) {
	type Output struct {
		n   int
		err error
	}
	type Input struct {
		str string
		val int
	}
	var tests = []struct {
		expected Output
		given    Input
	}{
		{Output{5, nil}, Input{"Hello Hello", 5}},
	}
	for _, test := range tests {
		r := strings.NewReader(test.given.str)
		lr := LimitReader(r, int64(test.given.val))
		p := make([]byte, test.given.val)
		n, err := lr.Read(p)
		n, err = lr.Read(p)
		if n != 0 {
			t.Errorf("Not expected n. Expected 0 Actual %d", n)
		}
		if err != io.EOF {
			t.Errorf("Not expected err. Expected io.EOF Actual %s", err)
		}
	}
}

func TestReadLessThanSizeOfLimitReader(t *testing.T) {
	type Output struct {
		n   int
		err error
	}
	type Input struct {
		str string
		val int
	}
	var tests = []struct {
		expected Output
		given    Input
	}{
		{Output{4, nil}, Input{"Hello Hello", 5}},
	}
	for _, test := range tests {
		r := strings.NewReader(test.given.str)
		lr := LimitReader(r, int64(test.given.val))
		p1 := make([]byte, test.given.val-1)
		n, err := lr.Read(p1)
		if n != test.expected.n {
			t.Errorf("Not expected n. Expected %d Actual %d", test.expected.n, n)
		}
		if err != test.expected.err {
			t.Errorf("Not expected err. Expected %s Actual %s", test.expected.err, err)
		}
	}
}
