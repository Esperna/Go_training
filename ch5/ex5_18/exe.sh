#!/bin/zsh
echo "go run ./fetch.go https://golang.org"
go run ./fetch.go https://golang.org
echo "diff index.html answer.html"
diff index.html answer.html
echo "rm index.html"
rm index.html
