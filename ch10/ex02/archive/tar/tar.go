package tar

import (
	"archive/tar"
	"ch10/ex02/archive"
	"fmt"
	"io"
	"log"
	"os"
)

const magicNum = "ustar"
const magicOffset = 257
const magicSize = 5

func decode(name string) (archive.Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	tr := tar.NewReader(f)
	return &tarReader{tr}, nil
}

func init() {
	archive.RegisterFormat("tar", magicNum, magicOffset, magicSize, decode)
}

type tarReader struct {
	reader *tar.Reader
}

func (tr *tarReader) List() {
	for {
		trHeader, err := tr.reader.Next()
		if err == io.EOF {
			break
		}
		fmt.Println(trHeader.Name)
	}
}
