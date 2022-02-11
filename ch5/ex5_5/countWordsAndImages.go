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
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return words, images, err
}

func countWordsAndImages(n *html.Node) (words, images int) {
	words, images = wordAndImageFreq(nil, n)
	return words, images
}

func wordAndImageFreq(stack []string, n *html.Node) (words int, images int) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
	}
	if !(n.Data == "style" || n.Data == "script") {
		if n.Type == html.TextNode {
			//fmt.Printf("%s", n.Data)
			words += wordfreq(bytes.NewBufferString(n.Data))
		}
	}
	if n.Data == "img" {
		for _, img := range n.Attr {
			if img.Key == "src" {
				//fmt.Println(img.Val)
				images++
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wordCount, imageCount := wordAndImageFreq(stack, c)
		words += wordCount
		images += imageCount
	}
	return words, images
}

func wordfreq(r io.Reader) (words int) {
	seen := make(map[string]int)
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
	for _, v := range seen {
		words += v
	}
	return words
}
