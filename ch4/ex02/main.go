package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hash = flag.String("h", "sha256", "hash function type")

func main() {
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		b := sc.Bytes()
		fmt.Printf("%v = %v\n", b, string(b))
		if *hash == "sha256" {
			c := sha256.Sum256(b)
			fmt.Printf("%x\n", c)
		} else if *hash == "sha384" {
			c := sha512.Sum384(b)
			fmt.Printf("%x\n", c)
		} else if *hash == "sha512" {
			c := sha512.Sum512(b)
			fmt.Printf("%x\n", c)
		}
	}
}
