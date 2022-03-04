package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Reader struct {
	s string
	i int64
}

func main() {
	rd := NewReader("<html><title>under construction</title></html>")
	doc, err := html.Parse(rd)
	if err != nil {
		fmt.Errorf("html parse failed:%s", err)
	}
	fmt.Printf("First node type: %d\n", doc.Type)

}

func NewReader(s string) *Reader {
	var rd Reader
	rd = Reader{s, 0}
	return &rd
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	length := len(b)
	if length == 0 {
		return 0, nil
	} else if length < 0 {
		err = fmt.Errorf("invalid byte length")
		return n, err
	}

	r.s = string(b)
	r.i += int64(len(b))
	return n, err
}
