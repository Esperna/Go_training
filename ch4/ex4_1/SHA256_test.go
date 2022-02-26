package main

import (
	"crypto/sha256"
	"testing"
)

func TestCountBitDiff(t *testing.T) {

	var tests = []struct {
		input1, input2 []byte
		want           int
	}{
		{[]byte("x"), []byte("x"), 0},
		{[]byte("X"), []byte("X"), 0},
		{[]byte("HelloWorld"), []byte("HelloWorld"), 0},
	}
	for _, test := range tests {
		sum1 := sha256.Sum256(test.input1)
		sum2 := sha256.Sum256(test.input2)
		if CountBitDiff(sum1, sum2) != test.want {
			t.Errorf("countBitDiff(sum1, sum2) != 0. the value is %d\n %x\n%x\n", CountBitDiff(sum1, sum2), sum1, sum2)
		}
	}
}
