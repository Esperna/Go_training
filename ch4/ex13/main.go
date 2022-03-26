package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Movie struct {
	Poster string
}

func main() {
	length := len(os.Args)
	if length == 3 {
		searchword := os.Args[1]
		apikey := os.Args[2]
		resp, err := http.Get("http://www.omdbapi.com/?t=" + searchword + "&apikey=" + apikey)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		var movie Movie
		if err := json.Unmarshal(b, &movie); err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", movie.Poster)
		resp, err = http.Get(movie.Poster)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		file, err := os.Create("movie.jpg")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		io.Copy(file, resp.Body)
		resp.Body.Close()
	} else {
		fmt.Println("Invalid Number of Argument. \"./main \"search word\" \"API key\"\" ")
	}
}