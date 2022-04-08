package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestCharCount(t *testing.T) {
	given := "in1"
	f, err := os.Open(given)
	if err != nil {
		t.Errorf("%s", err)
	}
	actual := countChar(f)
	for c, n := range actual.counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for c := range actual.counts {
		if actual.counts[c] != 1 {
			t.Errorf("counts is not expected. actual %v", actual.counts)
		}
	}
	if actual.utflen[0] != 0 || actual.utflen[1] != 4 || actual.utflen[2] != 0 || actual.utflen[3] != 0 {
		t.Errorf("utflen is not expected. actual %v", actual.utflen)
	}
	if actual.invalid != 0 {
		t.Errorf("invalid is not expected. actual %v", actual.invalid)
	}
}

func TestCharCountInvalid(t *testing.T) {
	rd := bytes.NewReader([]byte{0xff})
	actual := countChar(rd)
	for c, n := range actual.counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for c := range actual.counts {
		if actual.counts[c] != 1 {
			t.Errorf("counts is not expected. actual %v", actual.counts)
		}
	}
	if actual.utflen[0] != 0 || actual.utflen[1] != 0 || actual.utflen[2] != 0 || actual.utflen[3] != 0 {
		t.Errorf("utflen is not expected. actual %v", actual.utflen)
	}
	if actual.invalid != 1 {
		t.Errorf("invalid is not expected. actual %v", actual.invalid)
	}
}
