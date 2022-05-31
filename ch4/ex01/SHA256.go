package main

import (
	"bytes"
	"ch4/ex01/popcount"
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
	var i1, i2 uint64
	buf1 := bytes.NewReader(b1[:])
	if err := binary.Read(buf1, binary.LittleEndian, &i1); err != nil {
		fmt.Printf("binary read failure:%s", err)
	}
	buf2 := bytes.NewReader(b2[:])
	if err := binary.Read(buf2, binary.LittleEndian, &i2); err != nil {
		fmt.Printf("binary read failure:%s", err)
	}

	return popcount.PopCount(i1 ^ i2)
}
