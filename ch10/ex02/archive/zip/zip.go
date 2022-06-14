package zip

import (
	"archive/zip"
	"ch10/ex02/archive"
	"fmt"
	"log"
)

const magicNum = "PK"
const magicOffset = 0
const magicSize = 2

func decode(name string) (archive.Reader, error) {
	rc, err := zip.OpenReader(name)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	return &zipReader{rc}, nil
}

func init() {
	archive.RegisterFormat("zip", magicNum, magicOffset, magicSize, decode)
}

type zipReader struct {
	readCloser *zip.ReadCloser
}

func (zr *zipReader) List() {
	for _, f := range zr.readCloser.File {
		fmt.Println(f.Name)
	}
}
