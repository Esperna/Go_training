#!/bin/zsh

go build ./outline.go
echo "./outline https://golang.org/ quote_slide0 > out.txt"
./outline https://golang.org/ quote_slide0 > out.txt
open out.txt
