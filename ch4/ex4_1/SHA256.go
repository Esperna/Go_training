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
	count := 0
	for i := 0; i < 4; i++ {
		c1_i := binary.LittleEndian.Uint64(c1[8*i : 8*(i+1)])
		c2_i := binary.LittleEndian.Uint64(c2[8*i : 8*(i+1)])
		c3_i := c1_i ^ c2_i
		count += popcount.PopCount(c3_i)
		fmt.Println(count)
	}
}
