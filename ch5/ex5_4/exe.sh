#!/bin/zsh

go build ./dispOtherlinks.go
fetch https://golang.org | ./dispOtherlinks
