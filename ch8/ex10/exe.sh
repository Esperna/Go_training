#!/bin/zsh
go build findlinks.go
#./findlinks -depth=1 https://www.google.com
#./findlinks -depth=1 https://yahoo.co.jp
#./findlinks -depth=1 http://golang.org
./findlinks -depth=1 http://gopl.io
