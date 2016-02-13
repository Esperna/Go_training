package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		//fmt.Println(<-ch)
		f, err := os.OpenFile("fetch_result.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)		
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString(<-ch); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<-string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	ch <- string(b)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", secs, nbytes, url)
}
