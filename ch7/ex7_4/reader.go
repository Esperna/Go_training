package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type Reader struct {
	s string
	i int64
}

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	rd := NewReader(s)
	//rd := strings.NewReader(s)
	doc, err := html.Parse(rd)
	if err != nil {
		log.Fatalf("html parse failed:%s", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
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
	return &Reader{s, 0}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return n, io.EOF
	}
	n = copy(p, r.s[r.i:])
	r.i += int64(n)
	return n, err
}
