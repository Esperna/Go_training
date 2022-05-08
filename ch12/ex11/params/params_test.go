package params

import (
	"log"
	"testing"
)

func TestPack(t *testing.T) {
	type queryParam struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	data := queryParam{Labels: []string{"golang", "programming"}, MaxResults: 100, Exact: true}
	url, err := Pack(&data)
	if err != nil {
		log.Fatalf("Pack failed. data: %v err: %s", data, err)
	}

	want := "http://localhost:12345/search?l=golang&l=programming&max=100&x=true"
	actual := url.String()
	if actual != want {
		t.Errorf("url is %s, want %s", actual, want)
	}
}
