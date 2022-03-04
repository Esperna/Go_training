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
	rd := NewReader("<html><head><title>under construction</title></head><body></body></html>")
	doc, err := html.Parse(rd)
	if err != nil {
		fmt.Errorf("html parse failed:%s", err)
	}
	dispNode(nil, doc)
}

func dispNode(stack []string, n *html.Node) {
	fmt.Printf("Type: %v, DataAtom: %v, Data: %s, Namespace: %s, Attr: %v\n", n.Type, n.DataAtom, n.Data, n.Namespace, n.Attr)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dispNode(stack, c)
	}
}

func NewReader(s string) *Reader {
	var rd Reader
	rd = Reader{s, 0}
	return &rd
}

func (r *Reader) Read(b []byte) (n int, err error) {
	length := len(b)
	if length == 0 {
		return 0, nil
	} else if length < 0 {
		err = fmt.Errorf("invalid byte length")
		return n, err
	}
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}

	r.s = string(b)
	r.i += int64(len(b))
	return n, err
}
