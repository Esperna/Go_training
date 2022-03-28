package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Invalid number of argument. more than 1 is expected")
		os.Exit(1)
	}
	responses := make(chan string, len(os.Args))
	cancelled := func() bool {
		select {
		case <-responses:
			return true
		default:
			return false
		}
	}
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				responses <- fmt.Errorf("failed to get %s: %v", url, err).Error()
			}
			if cancelled() {
				fmt.Printf("cancelled %s after GET\n", url)
				return
			}
			b, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				responses <- fmt.Errorf("failed to read %s: %v", url, err).Error()
			}
			if cancelled() {
				fmt.Printf("cancelled after reading%s\n", url)
				return
			}
			responses <- string(b)
		}(url)
	}
	s := <-responses
	close(responses)
	wg.Wait()
	fmt.Printf("%s", s)
}
