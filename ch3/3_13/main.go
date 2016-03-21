package main

import "fmt"

func main() {
	const (
		KB int64 = 1000
		GB int64 = 1000 * KB
		TB int64 = 1000 * GB
		PB int64 = 1000 * TB
		EB int64 = 1000 * PB
		ZB int64 = 1000 * EB
		YB int64 = 1000 * ZB
	)
	fmt.Println(KB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB) //int64の表示範囲を超えるので、Print文の拡張が必要
	fmt.Println(YB) //同上
}
