package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type PackagesInfo struct {
	ImportPath string
	Deps       []string
}

func main() {
	length := len(os.Args)
	if length < 2 {
		log.Printf("invalid num of args. Expected num of args is more than 0")
		return
	}
	wsPkgs := getPackagesInfo("...")
	argPkgs := getPackagesInfo(os.Args[1:]...)

	for _, arg := range argPkgs {
		fmt.Printf("Followings transitively depend on %s:\n", arg.ImportPath)
		for _, pkg := range wsPkgs {
			for _, dep := range pkg.Deps {
				if arg.ImportPath == dep {
					fmt.Println(pkg.ImportPath)
					break
				}
			}
		}
	}
}

func getPackagesInfo(pkgs ...string) []PackagesInfo {
	args := []string{"list", "-json"}
	args = append(args, pkgs...)
	cmd := exec.Command("go", args...)

	out, err := cmd.Output()
	if err != nil {
		log.Printf("cmd failed:%s", err)
	}
	r := bytes.NewReader(out)
	var pkgInfos []PackagesInfo
	decoder := json.NewDecoder(r)

	for {
		var pkgInfo PackagesInfo
		err := decoder.Decode(&pkgInfo)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("JSON unmarshaling failed:%s", err)
			}
			return pkgInfos
		}
		pkgInfos = append(pkgInfos, pkgInfo)
	}
}
