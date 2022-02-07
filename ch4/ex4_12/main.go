package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:safe_title`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func main() {
	length := len(os.Args)
	if length == 2 {
		end := 600
		for i := 1; i < end; i++ {
			if i == 404 {
				continue
			}
			resp, err := http.Get("https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json")
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			b, err := ioutil.ReadAll(resp.Body)
			var comic Comic
			if err := json.Unmarshal(b, &comic); err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			if strings.Contains(comic.Transcript, os.Args[1]) {
				url := "https://xkcd.com/" + strconv.Itoa(i)
				fmt.Printf("%s/\t%s\n%s\n", comic.Title, url, comic.Transcript)
			}
		}
	} else {
		fmt.Println("Invalid number of Argument. \"./main \"search word\"\" is expected")
	}

}
