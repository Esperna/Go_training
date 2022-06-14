package main

import (
	"ch10/ex02/archive"
	_ "ch10/ex02/archive/tar"
	_ "ch10/ex02/archive/zip"
	"fmt"
	"log"
)

func main() {
	files := []string{"archive1.zip", "archive2.zip", "archive3.tar"}
	for _, file := range files {
		fmt.Println(file)
		rd, err := archive.ReadArchive(file)
		if err != nil {
			log.Printf("%s", err)
		}
		rd.List()
	}
}
