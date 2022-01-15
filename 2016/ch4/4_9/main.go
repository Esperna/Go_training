package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = bufio.NewScanner(os.Stdin)

func main() {
	input.Split(bufio.ScanWords)
	wordfreq()
}

func wordfreq() {
	counts := make(map[string]int)
	for {
		if input.Scan() {
			var r = input.Text()
			counts[r]++
		} else {
			break
		}
	}

	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
