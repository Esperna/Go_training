package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

type ArchiveType int

const (
	Unknown ArchiveType = iota
	Zip
	Tar
)

func main() {
	filepaths := []string{"archive1.zip", "archive2.zip", "archive3.tar"}
	for _, path := range filepaths {
		acvType, err := checkArchiveType(path)
		if err != nil {
			log.Fatal(err)
		}
		if acvType == Zip {
			err = readZip(path)
			if err != nil {
				log.Fatal(err)
			}
		} else if acvType == Tar {

		}
	}
}

func checkArchiveType(filepath string) (ArchiveType, error) {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		return Unknown, err
	}
	fmt.Printf("Archive File type:")
	if isZip(f) {
		fmt.Println("zip")
		return Zip, err
	} else if isTar(f) {
		fmt.Println("tar")
		return Tar, err
	} else {
		fmt.Println("unknown")
		return Unknown, err
	}
}

func readZip(filepath string) error {
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

func readTar(filepath string) error {
	return nil
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
	const offset = 257
	const size = 6
	var b [512]byte
	n, err := r.Read(b[offset : offset+size])
	if err != nil {
		log.Printf("%s", err)
		return false
	}
	if n > 0 {
		for i := 0; i < size; i++ {
			fmt.Printf("%s\n", string(b[i]))
		}
	}
	return false
}
