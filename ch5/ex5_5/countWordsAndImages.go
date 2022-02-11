package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	words, images, _ := CountWordsAndImages(os.Args[1])
	fmt.Printf("%v\t%v\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return words, images, err
}

func countWordsAndImages(n *html.Node) (words, images int) {
	wordfreq := make(map[string]int)
	wordfreq = wordfreqTextNode(nil, n, wordfreq)
	for k, v := range wordfreq {
		fmt.Printf("%s\n", k)
		words += v
	}
	return words, images
}

func wordfreqTextNode(stack []string, n *html.Node, seen map[string]int) map[string]int {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
	}
	if !(n.Data == "style" || n.Data == "script") {
		if n.Type == html.TextNode {
			//fmt.Printf("%s", n.Data)
			seen = wordfreq(bytes.NewBufferString(n.Data), seen)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		seen = wordfreqTextNode(stack, c, seen)
	}
	return seen
}

func wordfreq(r io.Reader, seen map[string]int) map[string]int {
	input := bufio.NewScanner(r)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		seen[line]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
	return seen
}
