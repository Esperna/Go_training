#!/bin/zsh
echo "go run ./fetch.go https://golang.org"
go run ./fetch.go https://golang.org
echo "fetch.go https://golang.org > answer.html"
fetch https://golang.org > answer.html
echo "diff index.html answer.html"
diff index.html answer.html
echo "rm index.html answer.html"
rm index.html answer.html
