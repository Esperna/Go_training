package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type Reader struct {
	i    int64
	size int64
	r    io.Reader
}

func main() {
	r := strings.NewReader("Hello Hello")
	//lr := io.LimitReader(r, 6)
	lr := LimitReader(r, 6)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &Reader{0, n, r}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.i >= r.size {
		return n, io.EOF
	}
	n, _ = r.r.Read(p[:r.size])
	if int64(n) > r.size {
		n = int(r.size)
	}
	r.i += int64(n)
	return n, err
}
