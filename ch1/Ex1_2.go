//Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	sep := " "
	for index, arg := range os.Args[1:] {
		fmt.Println(index,sep,arg)
	}
}
