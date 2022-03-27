#!/bin/zsh
go build mirror.go
./mirror -depth=1 https://www.google.com
#./findlinks -depth=1 https://yahoo.co.jp
#./findlinks -depth=1 http://golang.org
#./findlinks -depth=1 http://gopl.io
