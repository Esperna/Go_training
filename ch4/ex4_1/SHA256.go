package main

import (
	"ch4/ex4_1/popcount"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := CountBitDiff(c1, c2)
	fmt.Println(count)

}

func CountBitDiff(b1, b2 [sha256.Size]byte) int {
	count := 0
	for i := 0; i < 4; i++ {
		b1_i := binary.LittleEndian.Uint64(b1[8*i : 8*(i+1)])
		b2_i := binary.LittleEndian.Uint64(b2[8*i : 8*(i+1)])
		count += popcount.PopCount(b1_i ^ b2_i)
	}
	return count
}
