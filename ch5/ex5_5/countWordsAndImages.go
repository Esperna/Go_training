package main

import (
	"fmt"
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
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	dispTextNode(nil, n)
	return
}

func dispTextNode(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
	}
	if !(n.Data == "style" || n.Data == "script") {
		if n.Type == html.TextNode {
			fmt.Printf("%s", n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dispTextNode(stack, c)
	}
}
