package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filepaths := []string{"archive1.zip", "archive2.zip"}
	for _, path := range filepaths {
		err := checkArchiveType(path)
		if err != nil {
			log.Fatal(err)
		}
		err = readArchive(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func checkArchiveType(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	fmt.Printf("Archive File type:")
	if isZip(f) {
		fmt.Println("zip")
	} else if isTar(f) {
		fmt.Println("tar")
	} else {
		fmt.Println("unknown")
	}
	f.Close()
	return nil
}

func readArchive(filepath string) error {
	r, err := zip.OpenReader(filepath)
	defer r.Close()
	for i, f := range r.File {
		fmt.Printf("---Entry%d---\n", i+1)
		if f.FileInfo().IsDir() {
			fmt.Printf("name:%s\n", f.Name)
		} else {
			fmt.Printf("name:%s\n", f.Name)
			fmt.Println("Contents:")
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.CopyN(os.Stdout, rc, 68)
			if err != nil {
				fmt.Println(err)
			}
			rc.Close()
		}
	}
	return err
}

func isZip(r io.Reader) bool {
	var b [2]byte
	n, err := r.Read(b[:])
	if err != nil {
		log.Printf("%s", err)
		return false
	}
	if n > 0 {
		if string(b[0]) == "P" && string(b[1]) == "K" {
			return true
		}
	}
	return false
}

func isTar(r io.Reader) bool {
	return false
}
