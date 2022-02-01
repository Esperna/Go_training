package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length == 2 {
		c := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("%x\n", c)
	} else if length == 3 {
		if os.Args[2] == "-256" {
			c := sha256.Sum256([]byte(os.Args[1]))
			fmt.Printf("%x\n", c)
		} else if os.Args[2] == "-384" {
			c := sha512.Sum384([]byte(os.Args[1]))
			fmt.Printf("%x\n", c)
		} else if os.Args[2] == "-512" {
			c := sha512.Sum512([]byte(os.Args[1]))
			fmt.Printf("%x\n", c)

		} else {
			fmt.Printf("Invalid Option. -256,-384,or -512 is Expected\n")
		}
	} else {
		fmt.Printf("Invalid Number of Argument. Expected Number of Argument is 2 or 3.\n")
		fmt.Printf("Ex: ./main abc -256\n")
	}
}
