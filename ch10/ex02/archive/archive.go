package archive

import (
	"log"
	"os"
)

type format struct {
	name, magic  string
	offset, size int
	decode       func(name string) (Reader, error)
}

type Reader interface {
	List()
}

var formats []format

func RegisterFormat(name, magic string, offset, size int, decode func(name string) (Reader, error)) {
	formats = append(formats, format{name, magic, offset, size, decode})
}

func ReadArchive(path string) (Reader, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	format := getFormat(f)
	rd, err := format.decode(path)
	return rd, nil
}

func getFormat(f *os.File) format {
	var b [512]byte
	_, err := f.Read(b[:])
	if err != nil {
		log.Printf("%s", err)
		return format{}
	}
	for _, format := range formats {
		if string(b[format.offset:format.offset+format.size]) == format.magic {
			return format
		}
	}
	return format{}
}
