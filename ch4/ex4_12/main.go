package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
	end := 600
	for i := 1; i < end; i++ {
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
		fmt.Printf("%d\t%s/%s\t%s\n", comic.Num, comic.Month, comic.Year, comic.Title)
	}
}
