package main

import (
	"fmt"
	"math/big"
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	//	fmt.Println(ZB) //overflow
	//	fmt.Println(YB) //overflow
	bigNum := big.NewInt(EB)
	bigThousand := big.NewInt(1000)
	fmt.Println(bigNum.Mul(bigNum, bigThousand)) //	fmt.Println(ZB)
	fmt.Println(bigNum.Mul(bigNum, bigThousand)) //	fmt.Println(YB)
}
