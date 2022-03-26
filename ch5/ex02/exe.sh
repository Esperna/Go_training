#!/bin/zsh

rm ./mapHtmlElemValue
go build ./mapHtmlElemValue.go
fetch https://golang.org | ./mapHtmlElemValue

