package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type Reader struct {
	s   string
	i   int64
	buf []byte
}

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	rd := NewReader(s)
	doc, err := html.Parse(rd)
	if err != nil {
		log.Fatalf("html parse failed:%s", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	fmt.Println(n.Type)
	if n.Type == html.ElementNode && n.Data == "a" {
		fmt.Println(n.Data)
		for _, a := range n.Attr {
			fmt.Println(a)
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func NewReader(s string) *Reader {
	var rd Reader
	buffer := make([]byte, len(s))
	rd = Reader{s, 0, buffer}
	return &rd
}

func (r *Reader) Read(p []byte) (n int, err error) {
	length := len(p)
	if r.i >= int64(length) && r.i != 0 {
		return n, io.EOF
	}
	n = copy(r.buf, p)
	r.i += int64(n)
	return n, err
}
