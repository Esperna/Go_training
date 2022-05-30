package main

import (
	"encoding/json"
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
	seen := make(map[string]bool)
	var checkDependencies func(string) TransitivePackagesResult

	checkDependencies = func(arg string) TransitivePackagesResult {
		cmd := exec.Command("go", "list", "-json", arg)
		out, err := cmd.Output()
		if err != nil {
			log.Printf("cmd failed:%s", err)
		}
		var result TransitivePackagesResult
		if err := json.Unmarshal(out, &result); err != nil {
			log.Fatalf("JSON unmarshaling failed:%s", err)
		}
		return result
	}

	result1 := checkDependencies(os.Args[1])
	for _, dep := range result1.Deps {
		seen[dep] = true
	}
	for _, arg := range result1.Deps {
		result2 := checkDependencies(arg)
		for _, dep := range result2.Deps {
			seen[dep] = true
		}
	}
	for k := range seen {
		fmt.Println(k)
	}
}
