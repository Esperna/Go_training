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
		if CountBitDiff(sum1, sum2) != popCountDiff(sum1, sum2) {
			t.Errorf("countBitDiff(sum1, sum2) != popCountDiff(sum1,sum2). countBitDiff is %d popCountDiff is %d", CountBitDiff(sum1, sum2), popCountDiff(sum1, sum2))
		}
	}
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func totalPopCount(bytes []byte) int {
	total := 0
	for _, b := range bytes {
		total += int(pc[b])
	}
	return total
}

func popCountDiff(sum1, sum2 [sha256.Size]byte) int {
	xorBytes := make([]byte, 0, sha256.Size)
	for i := 0; i < sha256.Size; i++ {
		xorBytes = append(xorBytes, sum1[i]^sum2[i])
	}
	return totalPopCount(xorBytes)
}
