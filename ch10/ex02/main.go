package main

import (
	"archive/tar"
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
			err = readTar(path)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}

func checkArchiveType(filepath string) (ArchiveType, error) {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		return Unknown, err
	}
	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Printf("%s", err)
		return Unknown, err
	}
	if n > 0 {
		const zipOffset = 0
		const zipSize = 2
		//magic field offset and size
		const tarOffset = 257
		const tarSize = 6
		if string(b[zipOffset:zipOffset+zipSize]) == "PK" {
			fmt.Println("Archive File type:zip")
			return Zip, nil
		} else if string(b[tarOffset:tarOffset+tarSize]) == "ustar"+string([]byte{0}) {
			fmt.Println("Archive File type:tar")
			return Tar, nil
		} else {
			fmt.Println("Archive File type:unknown")
		}
	}
	return Unknown, nil
}

func readZip(filepath string) error {
	r, err := zip.OpenReader(filepath)
	defer r.Close()
	for _, f := range r.File {
		fmt.Println(f.Name)
	}
	return err
}

func readTar(filepath string) error {
	f, _ := os.Open(filepath)
	defer f.Close()
	tr := tar.NewReader(f)

	for {
		trHeader, err := tr.Next()
		if err == io.EOF {
			break
		}
		fmt.Println(trHeader.Name)
	}

	return nil
}
