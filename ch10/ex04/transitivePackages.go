package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type TransitivePackagesResult struct {
	Deps []string
}

func main() {
	length := len(os.Args)
	if length != 2 {
		log.Printf("invalid num of args. Expected is 1")
		return
	}
	cmd := exec.Command("go", "list", os.Args[1])
	out, err := cmd.Output()
	if err != nil {
		log.Printf("cmd failed:%s", err)
	}
	fmt.Printf("%s", out)

}
