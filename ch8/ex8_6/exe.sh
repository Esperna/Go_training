#!/bin/zsh
go build findlinks.go
./findlinks -depth=3 http://gopl.io/
